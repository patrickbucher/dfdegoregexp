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
	pattern := os.Args[1]

	input := bufio.NewReader(os.Stdin)
	var line []byte
	var err error
	for ; err != io.EOF; line, err = input.ReadBytes('\n') {
		if len(line) == 0 {
			continue
		}
		if ok, _ := regexp.Match(pattern, line); ok {
			fmt.Print(string(line))
		}
	}
}
