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
		"echo": regexp.MustCompile("echo (.+)"),
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
				case "echo":
					if len(match) != 2 {
						fmt.Fprintln(os.Stderr, "Error occured during echo...")
					} else {
						fmt.Fprintln(os.Stdout, match[1])
					}
				}
			}
		}

		if !matched {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}
