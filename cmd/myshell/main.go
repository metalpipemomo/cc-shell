package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

	cmd = strings.TrimSpace(cmd)

	if err != nil {
		fmt.Fprintln(os.Stderr, "no command provided")
	}

	fmt.Fprintf(os.Stdout, "%s: command not found", cmd)
}
