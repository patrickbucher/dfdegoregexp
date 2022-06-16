package main

import (
	"fmt"
	"os"
	"os/exec"

	dr "github.com/patrickbucher/dfdegoregexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [term]", os.Args[0])
		os.Exit(1)
	}

	cmd := exec.Command("go", "doc", os.Args[1])
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dr.FilterLines(dr.FunctionDeclaration, cmdOut, os.Stdout)

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
