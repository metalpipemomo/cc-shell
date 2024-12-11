package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

		cmd = strings.TrimSpace(cmd)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command: ", err)
			return
		}

		if cmd == "" {
			fmt.Fprintln(os.Stdout, "")
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}

	}
}
