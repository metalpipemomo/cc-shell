package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	cmds := map[string]*regexp.Regexp{
		"exit": regexp.MustCompile("exit ([0-9])+"),
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command: ", err)
			return
		}

		cmd = strings.TrimSpace(cmd)

		if cmd == "" {
			continue
		}

		matched := false
		for name, regex := range cmds {
			if match := regex.FindStringSubmatch(cmd); match != nil {
				matched = true
				switch name {
				case "exit":
					exitCode, parseErr := strconv.Atoi(match[1])
					if parseErr != nil {
						fmt.Fprintln(os.Stderr, "Error occured during exit: ", err)
					} else {
						os.Exit(exitCode)
					}
				}
			}
		}

		if !matched {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}
