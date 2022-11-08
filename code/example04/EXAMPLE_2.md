# Checking test coverage

## Full coverage
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

Run test with coverage:
```
➜  godublin git:(main) ✗ go test -coverprofile=cover.out
PASS
coverage: 0.0% of statements
total coverage: 100.0% of statements
ok  	godublin	0.118s
```


## Reduced coverage

`hello.txtar`

```bash
# With no arguments, fail and print a usage message
! exec godublin
! stdout .
stderr 'usage: Hello Gopher NAME'

# With an argument, greet a Gopher using provided name
#exec godublin Shawn
#stdout 'Hello Gopher, Shawn'
#! stderr .
```

Run tests with coverage:

```bash
➜  godublin git:(main) ✗ go test -coverprofile=cover.out
PASS
coverage: 0.0% of statements
total coverage: 60.0% of statements
ok  	godublin	0.111s
```
