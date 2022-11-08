# Comparing output with multiple golden files

`hello.txtar`

```
# With an argument, greet a Gopher and compare with file content.
exec godublin Kate
cmp stdout golden.txt

# With an argument, greet a Gopher and compare with invalid content.
exec godublin Michael
! cmp stdout golden-invalid.txt

-- golden.txt --
Hello Gopher, Kate!
-- golden-invalid.txt --
Hello Gopher, Kate!
```


Run test:
```bash
$ go test -run TestHello/hello
PASS
ok  	godublin	0.115s
```

Run test in verbose mode:
```bash
➜  godublin git:(main) ✗ go test -run TestHello/hello -v
=== RUN   TestHello
=== RUN   TestHello/hello
=== PAUSE TestHello/hello
=== CONT  TestHello/hello
  ...
        # With an argument, greet a Gopher and compare with file content. (0.005s)
        > exec godublin Kate
        [stdout]
        Hello Gopher, Kate!
        > cmp stdout golden.txt
        # With an argument, greet a Gopher and compare with invalid content. (0.005s)
        > exec godublin Michael
        [stdout]
        Hello Gopher, Michael!
        > ! cmp stdout golden-invalid.txt
        PASS

--- PASS: TestHello (0.00s)
    --- PASS: TestHello/hello (0.01s)
PASS
ok  	godublin	0.115s
```