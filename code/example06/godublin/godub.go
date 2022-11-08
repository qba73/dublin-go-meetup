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
	message := fmt.Sprintf("Hello Gopher, %s\n", os.Args[1])
	content := []byte(message)
	fmt.Fprintln(os.Stdout, message)

	err := os.WriteFile("greet.out", content, 0644)
	if err != nil {
		return 1
	}
	return 0
}
