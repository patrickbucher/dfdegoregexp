package main

import (
	"dfde"
	"fmt"
	"os"
	"os/exec"
)

func main() {
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

	dfde.FilterLines(`^func [A-Za-z]+\(`, cmdOut, os.Stdout)

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
