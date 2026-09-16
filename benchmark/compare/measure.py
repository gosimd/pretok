import json,subprocess,pathlib,random,sys
root=pathlib.Path(sys.argv[1]).resolve()
repo=root/'current'
out=root/'results';out.mkdir(exist_ok=False)
servers={};gates={};rng=random.Random(15092026)
try:
 for name in ['old-go','new-go','old-simd','new-simd']:
  cwd=root/'baseline' if name.startswith('old') else repo
  p=subprocess.Popen([str(root/name),str(pathlib.Path(sys.argv[2]).resolve())],cwd=cwd,stdin=subprocess.PIPE,stdout=subprocess.PIPE,text=True)
  servers[name]=p;line=p.stdout.readline();gate=json.loads(line);assert gate['ready'];gates[name]=gate
  print('Verified',name,gate['requests'],flush=True)
 (out/'gates.json').write_text(json.dumps(gates,indent=2)+'\n')
 with (out/'samples.jsonl').open('w') as f:
  for lang in ['en-wikitext2','ru-wikipedia','ja-wikipedia','ar-wikipedia','zh-wikipedia']:
   for size in [256,4096,65536,1048576]:
    engines=['pretok-reused']
    if size==4096: engines+=['pretok-fresh','pretok-cursor']
    for engine in engines:
     for round in range(5):
      names=list(servers);rng.shuffle(names)
      for name in names:
       p=servers[name];p.stdin.write(json.dumps(dict(corpus=lang,size=size,engine=engine,round=round))+'\n');p.stdin.flush()
       row=json.loads(p.stdout.readline());row['build']=name;f.write(json.dumps(row)+'\n');f.flush()
     print('Timed',lang,size,engine,flush=True)
 (out/'complete.json').write_text(json.dumps(dict(complete=True,samples=600,rounds=5,min_time_ms=150,workers=1),indent=2)+'\n')
finally:
 for p in servers.values():
  p.stdin.close()
  try:p.wait(timeout=5)
  except subprocess.TimeoutExpired:p.kill();p.wait()
