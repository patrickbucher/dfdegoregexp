package main

import (
	"fmt"
	"os"
	"strings"

	dr "github.com/patrickbucher/dfdegoregexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [command] [arguments]", os.Args[0])
		os.Exit(1)
	}

	lines := dr.CommandOutput(os.Args[1], os.Args[1:]...)
	fmt.Print(strings.Join(dr.ExtractSectionsBad(lines), "\n"))

	// TODO: call function, once implemented
	// fmt.Print(strings.Join(dr.ExtractSectionsBetter(lines), "\n"))
}
