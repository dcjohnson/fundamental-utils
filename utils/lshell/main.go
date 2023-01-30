package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("> ")
		if scanner.Scan() {
			execute(
				strings.TrimSpace(
					scanner.Text(),
				),
			)
		}
	}
}

const (
	name = "lshell: A lisp inspired shell"
	CMDhelp = "help"
)

var (
	commands = map[string]string{
		CMDhelp: "The help menu",
	}
)

func execute(line string)  {
	switch line {
	case CMDhelp:
		help()
	}
}

func help() {
	fmt.Println(name)
	for k, v := range commands {
		fmt.Printf("%s %s", k, v)
	}
}
