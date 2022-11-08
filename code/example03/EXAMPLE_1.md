# Testing CLI tools

`hello.txtar`:
```
exec godublin
stdout 'Hello Dublin Gophers\n'
```

Run test:
```
➜  godublin git:(main) ✗ go test
PASS
ok  	godublin	0.114s
```

Run test in verbose mode:
```bash
➜  godublin git:(main) ✗ go test -v
=== RUN   TestHello
=== RUN   TestHello/hello
=== PAUSE TestHello/hello
=== CONT  TestHello/hello
   ...

        > exec godublin
        [stdout]
        Hello Dublin Gophers
        > stdout 'Hello Dublin Gophers\n'
        PASS

--- PASS: TestHello (0.00s)
    --- PASS: TestHello/hello (0.01s)
PASS
ok  	godublin	0.116s
```

## Failing test

Let's change output of the func `RunCLI()` and run tests again.

Run test:
```bash
➜  godublin git:(main) ✗ go test
--- FAIL: TestHello (0.00s)
    --- FAIL: TestHello/hello (0.01s)
        testscript.go:429: > exec godublin
            [stdout]
            Hello Dublin Gophers!
            > stdout 'Hello Dublin Gophers\n'
            FAIL: testdata/script/hello.txtar:2: no match for `Hello Dublin Gophers\n` found in stdout

FAIL
exit status 1
FAIL	godublin	0.116s
```

Let's change return value in the `RunCLI()` func from 0 (indicating success) to 1 (indicating failure).

Run test:
```
➜  godublin git:(main) ✗ go test
--- FAIL: TestHello (0.00s)
    --- FAIL: TestHello/hello (0.01s)
        testscript.go:429: > exec godublin
            [stdout]
            Hello Dublin Gophers
            [exit status 1]
            FAIL: testdata/script/hello.txtar:1: unexpected command failure

FAIL
exit status 1
FAIL	godublin	0.282s
```