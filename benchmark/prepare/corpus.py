#!/usr/bin/env python3
"""Explicit online preparation; generated corpus and reference files run offline."""
import argparse
import gzip
import hashlib
import io
import json
from pathlib import Path
import time
import requests
import pyarrow.parquet as pq
import tokenizers
from tokenizers import Regex, pre_tokenizers

ROOT = Path(__file__).resolve().parents[2]
SOURCES = [
    ('en-wikitext2','Salesforce/wikitext','b08601e04326c79dfdd32d625aee71d232d685c3','wikitext-2-raw-v1/train-00000-of-00001.parquet','CC-BY-SA-4.0'),
    ('ru-wikipedia','wikimedia/wikipedia','b04c8d1ceb2f5cd4588862100d08de323dccfbaa','20231101.ru/train-00000-of-00021.parquet','CC-BY-SA-3.0; GFDL where applicable'),
    ('ja-wikipedia','wikimedia/wikipedia','b04c8d1ceb2f5cd4588862100d08de323dccfbaa','20231101.ja/train-00000-of-00015.parquet','CC-BY-SA-3.0; GFDL where applicable'),
    ('ar-wikipedia','wikimedia/wikipedia','b04c8d1ceb2f5cd4588862100d08de323dccfbaa','20231101.ar/train-00000-of-00007.parquet','CC-BY-SA-3.0; GFDL where applicable'),
    ('zh-wikipedia','wikimedia/wikipedia','b04c8d1ceb2f5cd4588862100d08de323dccfbaa','20231101.zh/train-00000-of-00006.parquet','CC-BY-SA-3.0; GFDL where applicable'),
]

def sha(data): return hashlib.sha256(data).hexdigest()
def prefix(data,n): return data[:n].decode('utf-8',errors='ignore').encode('utf-8')
def write_gzip(path,data):
    with path.open('wb') as raw, gzip.GzipFile(filename='',fileobj=raw,mode='wb',mtime=0) as f: f.write(data)
    return sha(path.read_bytes())

class RemoteFile(io.RawIOBase):
    """Bounded verified HTTP ranges; no full Wikipedia shard download."""
    def __init__(self,url):
        self.url=url; self.pos=0; self.session=requests.Session(); self.transferred=0
        data,total=self.fetch(0,0); self.length=total
    def fetch(self,start,end):
        error=None
        for attempt in range(4):
            try:
                # Distinct URLs prevent intermediate caches confusing byte ranges.
                r=self.session.get(self.url+'?range=%d-%d'%(start,end),headers={'Range':'bytes=%d-%d'%(start,end)},timeout=90,stream=True)
                with r:
                    if r.status_code!=206: raise RuntimeError('Range request returned HTTP %s'%r.status_code)
                    cr=r.headers.get('Content-Range','')
                    expected='bytes %d-%d/'%(start,end)
                    if not cr.startswith(expected): raise RuntimeError('Wrong Content-Range: '+cr)
                    total=int(cr.split('/')[1])
                    data=r.raw.read(end-start+2)
                    if len(data)!=end-start+1: raise RuntimeError('Truncated/oversized range')
                    self.transferred+=len(data)
                    return data,total
            except (requests.RequestException,RuntimeError) as e:
                error=e
                if attempt<3: time.sleep(2**attempt)
        raise error
    def readable(self): return True
    def seekable(self): return True
    def tell(self): return self.pos
    def seek(self,offset,whence=0):
        p=offset if whence==0 else self.pos+offset if whence==1 else self.length+offset
        if p<0: raise ValueError('negative seek')
        self.pos=p; return p
    def read(self,n=-1):
        n=self.length-self.pos if n<0 else min(n,self.length-self.pos)
        if n<=0:return b''
        data,total=self.fetch(self.pos,self.pos+n-1)
        if total!=self.length:raise RuntimeError('Remote size changed')
        self.pos+=n; return data
    def close(self):
        self.session.close(); super().close()

def prepare_pool(source,budget):
    ident,repo,rev,shard,license=source
    url='https://huggingface.co/datasets/'+repo+'/resolve/'+rev+'/'+shard
    out=bytearray(); records=[]
    with RemoteFile(url) as remote:
        parquet=pq.ParquetFile(remote)
        columns=[c for c in ['text','id','url','title'] if c in parquet.schema.names]
        index=0
        for batch in parquet.iter_batches(batch_size=512,columns=columns):
            for row in batch.to_pylist():
                row_index=index;index+=1
                original=row['text'].encode()
                if not original:continue
                text=prefix(original,32768)
                separator=b'\n\n' if out else b''
                available=budget-len(out)
                if available<=len(separator):break
                text=prefix(text,available-len(separator))
                if not text:break
                out.extend(separator);start=len(out);out.extend(text)
                records.append({'row':row_index,'id':row.get('id'),'url':row.get('url'),'title':row.get('title'),
                                'source_sha256':sha(original),'source_bytes':[0,len(text)],'output_bytes':[start,len(out)],'separator_bytes':len(separator)})
                if budget-len(out)<4:break
            if budget-len(out)<4:break
        if len(out)<budget-3:raise RuntimeError('First pinned shard did not fill budget: '+ident)
        transferred=remote.transferred
    print(ident,len(out),'bytes;',len(records),'records;',transferred,'downloaded',flush=True)
    return bytes(out),{'repo':repo,'revision':rev,'shard':shard,'split':'train','license':license,'url':url,'records':records}

def cuts(data,size):
    out=[];pos=0
    while pos<len(data):
        end=min(pos+size,len(data))
        while end<len(data) and data[end]&0xc0==0x80:end-=1
        if end<=pos:raise ValueError('request size too small')
        out.append((pos,end));pos=end
    return out

def expected(pre,data):
    text=data.decode(); offsets=[0]
    for c in text:offsets.append(offsets[-1]+len(c.encode()))
    spans=[]
    for piece,(a,b) in pre.pre_tokenize_str(text):
        assert piece==text[a:b]
        spans.append([offsets[a],offsets[b]])
    return spans

def main():
    ap=argparse.ArgumentParser();ap.add_argument('--output',type=Path,default=ROOT/'benchmark/corpus');ap.add_argument('--offline',action='store_true');args=ap.parse_args()
    assert tokenizers.__version__=='0.22.2'
    dest=args.output;dest.mkdir(parents=True,exist_ok=True)
    patterns={}
    scanners={}
    for ident,directory in [('P01','testdata/patterns/P01'),('P05','testdata/qwen')]:
        fixture_dir=ROOT/directory
        fixture=json.loads((fixture_dir/'manifest.json').read_text())
        patterns[ident]={key:fixture[key] for key in ['pattern','pattern_sha256']}
        patterns[ident]['fixture_sha256']=sha((fixture_dir/'cases.jsonl.gz').read_bytes())
        scanners[ident]=pre_tokenizers.Split(Regex(fixture['pattern']),'isolated',invert=False)
    manifest={'schema':2,'patterns':patterns,
              'oracle_version':tokenizers.__version__,'oracle_native_sha256':sha((Path(tokenizers.__file__).parent/'tokenizers.abi3.so').read_bytes()),
              'generator_sha256':sha(Path(__file__).read_bytes()),
              'selection':'Pinned first shard, stored row order; nonempty records; 32 KiB per record; LF LF joins; 1 MiB UTF-8 prefix.', 'corpora':[]}
    pools=[]
    for source in SOURCES:
        ident=source[0]; cache=dest/(ident+'.source.json')
        if args.offline:
            info=json.loads(cache.read_text());data=gzip.decompress((dest/(ident+'.txt.gz')).read_bytes())
            assert info['pool_sha256']==sha(data)
            assert info['revision']==source[2]
        else:
            data,info=prepare_pool(source,1<<20);info['pool_sha256']=sha(data)
            cache.write_text(json.dumps(info,ensure_ascii=False,indent=2)+'\n')
        pools.append((ident,'natural',data,info,[ (size,cuts(data,size)) for size in [256,4096,65536,1048576] ]))
    data=bytearray();ranges=[]
    for token in ['a','Ж','中','😀','\u0301','1','²',"'re ",'!',' ','\u2003',' \r\n\t']:
        for n in [1,2,3,15,16,17,31,32,33,63,64,65,127,128,129,4096]:
            for suffix in ['', 'X']:
                start=len(data);data.extend((token*n+suffix).encode());ranges.append((start,len(data)))
    pools.append(('synthetic','synthetic',bytes(data),{'license':'CC0-1.0','generator':'Explicit runs in corpus.py'},[(0,ranges)]))
    for ident,kind,data,info,groups in pools:
        name=ident+'.txt.gz'; compressed=write_gzip(dest/name,data)
        text=data.decode()
        whole={ident:expected(pre,data) for ident,pre in scanners.items()}
        row={'id':ident,'kind':kind,'file':name,'sha256':sha(data),'compressed_sha256':compressed,'bytes':len(data),'scalars':len(text),
             'ascii_fraction':sum(b<128 for b in data)/len(data),
             'pattern_stats':{ident:{'whole_spans':len(spans),'max_span_bytes':max((b-a for a,b in spans),default=0),'mean_span_bytes':len(data)/max(1,len(spans))} for ident,spans in whole.items()},
             'source_file':ident+'.source.json' if kind=='natural' else None,'batches':[]}
        if kind=='natural':row['source_sha256']=sha((dest/row['source_file']).read_bytes())
        for cls,regex in [('whitespace',r'\s'),('number',r'\p{N}')]:
            row[cls+'_scalars']=len(pre_tokenizers.Split(Regex(regex),'removed',invert=True).pre_tokenize_str(text))
        for size,ranges in groups:
            requests=[{'start':a,'end':b,'spans':{ident:expected(pre,data[a:b]) for ident,pre in scanners.items()}} for a,b in ranges]
            payload=json.dumps(requests,separators=(',',':')).encode()
            file=ident+'-'+str(size)+'.json.gz'
            digest=write_gzip(dest/file,payload)
            row['batches'].append({'size':size,'file':file,'sha256':digest,'requests':len(requests),'spans':{ident:sum(len(r['spans'][ident]) for r in requests) for ident in scanners}})
        manifest['corpora'].append(row)
    (dest/'manifest.json').write_text(json.dumps(manifest,ensure_ascii=False,indent=2)+'\n')
    print('Prepared',len(pools),'pools and',sum(len(p['batches']) for p in manifest['corpora']),'batches',flush=True)

if __name__=='__main__':main()
