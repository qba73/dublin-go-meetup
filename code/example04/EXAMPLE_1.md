# Testing CLI app with two behaviours

`hello.txtar`

```bash
# With no arguments, fail and print a usage message
! exec godublin
! stdout .
stderr 'usage: Hello Gopher NAME'

# With an argument, greet a Gopher using provided name
exec godublin Shawn
stdout 'Hello Gopher, Shawn'
! stderr .
```

Run test:
```
➜  godublin git:(main) ✗ go test
PASS
ok  	godublin	0.177s
```

Run tests in verbose mode:
```bash
➜  godublin git:(main) ✗ go test -v
=== RUN   TestHello
=== RUN   TestHello/hello
=== PAUSE TestHello/hello
=== CONT  TestHello/hello
  ...

        # With no arguments, fail and print a usage message (0.076s)
        > ! exec godublin
        [stderr]
        usage: Hello Gopher NAME
        [exit status 1]
        > ! stdout .
        > stderr 'usage: Hello Gopher NAME'
        # With an argument, greet a Gopher using provided name (0.005s)
        > exec godublin Shawn
        [stdout]
        Hello Gopher, Shawn
        > stdout 'Hello Gopher, Shawn'
        > ! stderr .
        PASS

--- PASS: TestHello (0.00s)
    --- PASS: TestHello/hello (0.08s)
PASS
ok  	godublin	0.195s
```
