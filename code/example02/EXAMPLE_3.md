# Negating assertion with "!"

`cat2.txtar`
```
! exec cat not_existing_file.txt
stderr 'cat: not_existing_file.txt: No such file or directory'
! stdout .
```

Run test:
```bash
➜  godublin git:(main) ✗ go test -run TestScript/cat2
PASS
ok  	godub	0.255s
```

Run test in verbose mode:
```
➜  godublin git:(main) ✗ go test -run TestScript/cat2 -v
=== RUN   TestScript
=== RUN   TestScript/cat2
=== PAUSE TestScript/cat2
=== CONT  TestScript/cat2
    testscript.go:429: WORK=$WORK
 ...

        > ! exec cat not_existing_file.txt
        [stderr]
        cat: not_existing_file.txt: No such file or directory
        [exit status 1]
        > stderr 'cat: not_existing_file.txt: No such file or directory'
        > ! stdout .
        PASS

--- PASS: TestScript (0.00s)
    --- PASS: TestScript/cat2 (0.00s)
PASS
ok  	godub	0.111s
```

echo success succeeds
it prints the expected message to stdout, and
it prints nothing to stderr”
