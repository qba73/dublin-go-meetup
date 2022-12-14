# dublin-go-meetup

[Slides](bin/TestingInGo.pdf) and code samples for [Dublin Go Meetup](https://www.meetup.com/goireland/events/289521947/) talk "Testing in Go - Lessons Learnt."

[![Presentation](bin/TestinginGo.png)](bin/TestingInGo.pdf)



# Examples

All examples presented in this repository are based on an excellent introduction to the [testscript](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript) utility written by [John Arundel](https://bitfieldconsulting.com) [@bitfield](https://github.com/bitfield) in one of the chapters in his book ["The Power of Go Tests"](https://bitfieldconsulting.com/books/tests-print).

## Example 1 - testscript introduction

- [Assertions](code/example01/EXAMPLE_1.md)
- [Errors](code/example01/EXAMPLE_2.md)

## Example 2 - testscript assertions

- [Negative Assertions stdout](code/example02/EXAMPLE_1.md)
- [Negatinv Assertions stderr](code/example02/EXAMPLE_2.md)
- [Negative Assertions not existing files](code/example02/EXAMPLE_3.md)

## Example 3 - testscript CLI binary

- [CLI binary](code/example03/EXAMPLE_1.md)

## Example 4 - testscript CLI handling args

- [CLI binary handling args](code/example04/EXAMPLE_1.md)
- [CLI binary coverage](code/example04/EXAMPLE_2.md)

## Example 5 - testscript 'golden files'

- [CLI binary golden file](code/example05/EXAMPLE_1.md)
- [CLI binary golden files](code/example05/EXAMPLE_2.md)

## Example 6 - testscript 'regex' and file creation

- [CLI binary with regex](code/example06/EXAMPLE_1.md)


# How to run examples

To run examples, navigate to `godublin` packages in each directory and run the `go test` utility.
