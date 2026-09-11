#!/usr/bin/env python3
"""Explicit online preparation; generated corpus inputs run offline."""
import argparse
import gzip
import hashlib
import io
import itertools
import re
import random
import json
from pathlib import Path
import time
import requests
import pyarrow.parquet as pq

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

def main():
    ap=argparse.ArgumentParser();ap.add_argument('--output',type=Path,default=ROOT/'benchmark/corpus');ap.add_argument('--offline',action='store_true');args=ap.parse_args()
    dest=args.output;dest.mkdir(parents=True,exist_ok=True)
    patterns={}
    for ident,directory in [('P01','testdata/patterns/P01'),('P02','testdata/patterns/P02'),('P03','testdata/patterns/P03'),('P04','testdata/patterns/P04'),('P05','testdata/qwen'),('P09','testdata/patterns/P09'),('P13','testdata/patterns/P13'),('P16','testdata/patterns/P16'),('P07','testdata/patterns/P07'),('P24','testdata/patterns/P24'),('P20','testdata/patterns/P20'),('P19','testdata/patterns/P19'),('P06','testdata/patterns/P06'),('P11','testdata/patterns/P11'),('P21','testdata/patterns/P21'),('P17','testdata/patterns/P17'),('P18','testdata/patterns/P18'),('P08','testdata/patterns/P08'),('P12','testdata/patterns/P12'),('P15','testdata/patterns/P15'),('P23','testdata/patterns/P23')]:
        fixture_dir=ROOT/directory
        fixture=json.loads((fixture_dir/'manifest.json').read_text())
        patterns[ident]={key:fixture[key] for key in ['pattern','pattern_sha256']}
        # New patterns and aliases have no stored outputs: retain the existing P05 native-build guard.
        guard_dir=ROOT/'testdata/qwen' if ident not in ('P01','P05') else fixture_dir
        guard=json.loads((guard_dir/'manifest.json').read_text())
        patterns[ident]['fixture_pattern']=guard['pattern']
        patterns[ident]['fixture_sha256']=sha((guard_dir/'cases.jsonl.gz').read_bytes())
        patterns[ident]['full_coverage']=ident in ('P01','P04','P05','P09','P13','P24','P20','P19','P06')
    manifest={'schema':3,'patterns':patterns,
              'generator_sha256':sha(Path(__file__).read_bytes()),
              'case_tables_sha256':sha((ROOT/'internal/scan/case_tables.go').read_bytes()),
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
        for n in [1,2,3,4,5,6,7,8,15,16,17,31,32,33,63,64,65,127,128,129,4096]:
            for suffix in ['', 'X']:
                start=len(data);data.extend((token*n+suffix).encode());ranges.append((start,len(data)))
    # Common gap/ASCII-number diagnostics, shared by every pattern.
    # Exhaustive short subjects expose partial matches and UTF-8 transitions.
    for n in range(6):
        for chars in itertools.product(['0','9','a','１','²','😀'],repeat=n):
            start=len(data);data.extend(''.join(chars).encode());ranges.append((start,len(data)))
    # Ordered case-word branches: every short class transition, property-range
    # boundary, and reproducible random mix. Inputs are shared by all patterns.
    case_inputs=set()
    alphabet=['A','a','ǅ','中','ʰ','\u0301',' ','\n','1','/']
    for n in range(1,5):
        case_inputs.update(''.join(chars) for chars in itertools.product(alphabet,repeat=n))
    table=ROOT/'internal/scan/case_tables.go'
    for lo,hi in re.findall(r'\{0x([0-9A-Fa-f]+), 0x([0-9A-Fa-f]+), [123]\}',table.read_text()):
        for c in [int(lo,16)-1,int(lo,16),int(hi,16),int(hi,16)+1]:
            if 0<=c<0x110000 and not 0xD800<=c<=0xDFFF:
                ch=chr(c)
                case_inputs.update([ch,'A'+ch+'ABa '+ch+ch,'\u0301'+ch+'\u0301'])
    rng=random.Random(4)
    for _ in range(2000):
        case_inputs.add(''.join(rng.choices(alphabet+['\r',"'",'😀','²','\u2003'],k=rng.randrange(1,129))))
    for n in [1,2,3,15,16,17,31,32,33,63,64,65,4096]:
        for ch in ['A','a','ǅ','中','ʰ','\u0301']:
            case_inputs.update([ch*n+'ABC','\u0301'+ch*n+'ABC','ABC'+ch*n+'DEF',ch*n+'a'])
    for text in sorted(case_inputs):
        start=len(data);data.extend(text.encode());ranges.append((start,len(data)))
    # Line-end assertions and greedy whitespace, shared across all patterns.
    whitespace_inputs=set()
    for n in range(1,5):
        for chars in itertools.product([' ','\t','\r','\n','\u2003','\u2028','X'],repeat=n):
            text=''.join(chars)
            whitespace_inputs.update([text,'a'+text+'Z','a'+text+'\nZ'])
    for n in [1,2,3,15,16,17,31,32,33,63,64,65,127,128,129]:
        for ch in [' ','\t','\r','\n','\u2003','\u2028']:
            for suffix in ['X','\nX','\rX','\r\nX','\u2028X','']:
                whitespace_inputs.add('a'+ch*n+suffix)
    for text in sorted(whitespace_inputs):
        start=len(data);data.extend(text.encode());ranges.append((start,len(data)))
    number_inputs=set()
    for n in [1,2,3,4,5,6,7,8,9,15,16,17,31,32,33,63,64,65,4096]:
        for ch in ['0','１','²','Ⅳ','١','𐒠','1²Ⅳ𐒠']:
            number_inputs.update([ch*n,'A'+ch*n+'B',' '+ch*n+'\r\nX',ch*n+'\u0301'+ch*n])
    table=ROOT/'internal/scan/unicode_tables.go'
    number_section=table.read_text().split('var numberRanges')[1].split('}\n')[0]
    for lo,hi in re.findall(r'\{0x([0-9A-Fa-f]+), 0x([0-9A-Fa-f]+)\}',number_section):
        for c in [int(lo,16)-1,int(lo,16),int(hi,16),int(hi,16)+1]:
            if 0<=c<0x110000 and not 0xD800<=c<=0xDFFF:
                number_inputs.add('1'+chr(c)*4+'²Ⅳ𐒠A')
    for text in sorted(number_inputs):
        start=len(data);data.extend(text.encode());ranges.append((start,len(data)))
    # Optional suffixes follow the selected case-word endpoint, not the longest
    # later alternative; exercise prefix backtracking and failed/partial suffixes.
    suffix_inputs=set()
    suffixes=["'s","'t","'re","'ve","'m","'ll","'d","'RE","'LL","'ſ","'r","'v","'l","'x","’s","'s's"]
    words=['we','WE','AA中BB','\u0301ABC','\u0301','ǅaǅA','ʰ','中','a\u0301','1','!','\r\n',' ']
    for word,suffix,tail in itertools.product(words,suffixes,['','X','!','\r\n','1234567']):
        suffix_inputs.add(word+suffix+tail)
    for n in [1,2,3,15,16,17,31,32,33,63,64,65,4096]:
        for word in ['A','a','中','\u0301']:
            for suffix in suffixes:
                suffix_inputs.add(word*n+suffix+'ABC')
    for text in sorted(suffix_inputs):
        start=len(data);data.extend(text.encode());ranges.append((start,len(data)))
    # Exact literal-set boundaries and optional-prefix overlap, shared inputs.
    set_inputs=set()
    for lo,hi in re.findall(r'\{0x([0-9A-Fa-f]+), 0x([0-9A-Fa-f]+)\}',(ROOT/'internal/scan/exact_ranges.go').read_text()):
        for c in [int(lo,16)-1,int(lo,16),int(hi,16),int(hi,16)+1]:
            if 0<=c<0x110000 and not 0xD800<=c<=0xDFFF:
                ch=chr(c)
                for prefix in ['', ' ', '\t', '\u3000', '\u2003\u2003']:
                    set_inputs.update([prefix+ch,prefix+ch*3+'A中',ch+' '+ch])
    for n in range(1,5):
        set_inputs.update(''.join(chars) for chars in itertools.product(['\u3000',' ', 'A','1','中','(',')','|','[',']','\r','\n'],repeat=n))
    set_inputs.update(chr(c)+'X'+chr(c) for c in range(128))
    set_inputs.update(['[abc] (abc) a|b .,!?:;_', '.,!?…。，、।۔،', '\r\n'])
    for text in sorted(set_inputs):
        start=len(data);data.extend(text.encode());ranges.append((start,len(data)))
    # Ordered group transitions, overlaps, and exact group boundary neighbors.
    grouped_inputs=set()
    for lo,hi in re.findall(r'\{0x([0-9A-Fa-f]+), 0x([0-9A-Fa-f]+)\}',(ROOT/'internal/scan/grouped_ranges.go').read_text()):
        for c in [int(lo,16)-1,int(lo,16),int(hi,16),int(hi,16)+1]:
            if 0<=c<0x110000 and not 0xD800<=c<=0xDFFF:
                ch=chr(c)
                grouped_inputs.update([ch, '中'+ch+'ก', '가'+ch+'、', '、'+ch+'가', ch*2+'〮〯'])
    for n in range(1,5):
        grouped_inputs.update(''.join(chars) for chars in itertools.product(['中','あ','가','ก','ກ','ក','က','〮','〯','、','ㄅ','A'],repeat=n))
    for n in [1,2,3,15,16,17,31,32,33,63,64,65,4096]:
        for body in ['中あ','ก','가〮〯','、〮〯','ㄅ']:
            grouped_inputs.add('A'+body*n+'、가X')
    for text in sorted(grouped_inputs):
        start=len(data);data.extend(text.encode());ranges.append((start,len(data)))
    pools.append(('synthetic','synthetic',bytes(data),{'license':'CC0-1.0','generator':'Explicit runs in corpus.py'},[(0,ranges)]))
    for ident,kind,data,info,groups in pools:
        name=ident+'.txt.gz'; compressed=write_gzip(dest/name,data)
        text=data.decode()
        row={'id':ident,'kind':kind,'file':name,'sha256':sha(data),'compressed_sha256':compressed,'bytes':len(data),'scalars':len(text),
             'ascii_fraction':sum(b<128 for b in data)/len(data),
             'source_file':ident+'.source.json' if kind=='natural' else None,'batches':[]}
        if kind=='natural':row['source_sha256']=sha((dest/row['source_file']).read_bytes())
        for size,ranges in groups:
            requests=[{'start':a,'end':b} for a,b in ranges]
            payload=json.dumps(requests,separators=(',',':')).encode()
            file=ident+'-'+str(size)+'.json'
            (dest/file).write_bytes(payload)
            digest=sha(payload)
            row['batches'].append({'size':size,'file':file,'sha256':digest,'requests':len(requests)})
        manifest['corpora'].append(row)
    (dest/'manifest.json').write_text(json.dumps(manifest,ensure_ascii=False,indent=2)+'\n')
    print('Prepared',len(pools),'pools and',sum(len(p['batches']) for p in manifest['corpora']),'batches',flush=True)

if __name__=='__main__':main()
