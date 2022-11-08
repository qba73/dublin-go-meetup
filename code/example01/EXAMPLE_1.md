# Example 1

## Hello Gophers

`hello.txtar`:
```
exec echo 'Hello Gophers!'
stdout 'Hello Gophers\n'
```

Run test:
```bash
➜  godublin git:(main) ✗ go test
PASS
ok  	godub	0.149s
```

Run test in verbose mode:
```bash
➜  godublin git:(main) ✗ go test  -v
=== RUN   TestScript
=== RUN   TestScript/hello
=== PAUSE TestScript/hello
=== CONT  TestScript/hello
  ...

        > exec echo 'Hello Gophers!'
        [stdout]
        Hello Gophers!
        > stdout 'Hello Gophers'
        PASS

--- PASS: TestScript (0.00s)
    --- PASS: TestScript/hello (0.00s)
PASS
ok  	godub	0.112s
```


- The `exec` statement runs the program and performs `assertion`.
- The executed program must succeed. It's exit status must be zero.
- If any other assertion fails the exit status is not zero and it behaves as `t.Fatal` in a Go test
- `stdout` asserts that the output will match some regular expression. In this case the string 'Hello Gophers\n'
