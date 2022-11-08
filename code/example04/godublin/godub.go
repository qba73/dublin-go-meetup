package godublin

import (
	"fmt"
	"os"
)

func RunCLI() int {
	if len(os.Args[1:]) < 1 {
		fmt.Fprintln(os.Stderr, "usage: Hello Gopher NAME")
		return 1
	}
	fmt.Fprintf(os.Stdout, "Hello Gopher, %s\n", os.Args[1])
	return 0
}
