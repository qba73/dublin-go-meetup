# Example 2

## Errors

```bash
exec echo 'Hello Gophers!'
stdout 'Hi Gophers\n'
```

Run test:
```bash
➜  godublin git:(main) ✗ go test -run TestScript/hello-error
--- FAIL: TestScript (0.00s)
    --- FAIL: TestScript/hello-error (0.00s)
        testscript.go:429: > exec echo 'Hello Gophers!'
            [stdout]
            Hello Gophers!
            > stdout 'Hi Gophers\n'
            FAIL: testdata/script/hello-error.txtar:2: no match for `Hi Gophers\n` found in stdout

FAIL
exit status 1
FAIL	godub	0.109s
```
