# Negating assertion with "!"

`cat.txtar`
```
! exec cat not_existing_file.txt
stderr 'cat: not_existing_file.txt: No such file or directory'
```

Run test:
```bash
➜  godublin git:(main) ✗ go test -run TestScript/cat
PASS
ok  	godub	0.154s
```
