package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func searchPaths(cmd string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		joinedPath := filepath.Join(path, cmd)
		if _, err := os.Stat(joinedPath); err == nil {
			return joinedPath, true
		}
	}

	return "", false
}

func main() {

	cmds := map[string]*regexp.Regexp{
		"exit": regexp.MustCompile("exit ([0-9])+"),
		"echo": regexp.MustCompile("echo (.+)"),
		"type": regexp.MustCompile("type (.+)"),
		"pwd":  regexp.MustCompile("^pwd"),
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
						fmt.Fprintln(os.Stderr, "Error: Invalid argument length")
					} else {
						fmt.Fprintln(os.Stdout, match[1])
					}
				case "type":
					if len(match) != 2 {
						fmt.Fprintln(os.Stderr, "Error: Invalid argument length")
					} else {
						_, ok := cmds[match[1]]
						if !ok {
							joinedPath, foundInPath := searchPaths(match[1])
							if !foundInPath {
								fmt.Fprintf(os.Stderr, "%s: not found\n", match[1])
							} else {
								fmt.Fprintf(os.Stdout, "%s is %s\n", match[1], joinedPath)
							}
						} else {
							fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", match[1])
						}
					}
				case "pwd":
					exPath, err := os.Executable()
					if err != nil {
						fmt.Fprintln(os.Stderr, "Error: Unable to get working directory")
					}
					fmt.Fprintln(os.Stdout, filepath.Dir(exPath))
				}
			}
		}

		args := strings.Split(cmd, " ")
		extCmd, restArgs := args[0], args[1:]
		joinedPath, found := searchPaths(extCmd)

		if !matched && found {
			command := exec.Command(joinedPath, restArgs...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			matched = err == nil
		}

		if !matched {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd)
		}
	}
}
