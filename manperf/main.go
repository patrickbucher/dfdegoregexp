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

	// TODO: implement
	// fmt.Print(strings.Join(dr.ExtractSectionsBetter(lines), "\n"))
	// fmt.Print(strings.Join(dr.ExtractSectionsBetterPOSIX(lines), "\n"))
}
