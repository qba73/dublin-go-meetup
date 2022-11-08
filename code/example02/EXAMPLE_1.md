# Negating assertions with "!"

`hello.txtar`:
```
exec echo 'Hello Gophers'
! stdout 'Hello Gophers\n'
```

Run test:
```bash
➜  godublin git:(main) ✗ go test
--- FAIL: TestScript (0.00s)
    --- FAIL: TestScript/hello (0.00s)
        testscript.go:429: > exec echo 'Hello Gophers'
            [stdout]
            Hello Gophers
            > ! stdout 'Hello Gophers\n'
            FAIL: testdata/script/hello.txtar:2: unexpected match for `Hello Gophers\n` found in stdout: Hello Gophers


FAIL
exit status 1
FAIL	godub	0.109s
```


