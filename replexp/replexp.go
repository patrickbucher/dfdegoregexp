package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [regexp]\n", os.Args[0])
		os.Exit(1)
	}

	var err error
	pattern, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "compile regexp `%s`: %v\n", os.Args[1], err)
		os.Exit(2)
	}

	input := bufio.NewReader(os.Stdin)
	var line []byte
	for ; err != io.EOF; line, err = input.ReadBytes('\n') {
		if len(line) == 0 {
			continue
		}
		if pattern.Match(line) {
			fmt.Print(string(line))
		}
	}
}
