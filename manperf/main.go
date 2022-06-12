package main

import (
	"fmt"
	"os"
	"strings"

	dr "github.com/patrickbucher/dfdegoregexp"
)

func main() {
	lines := dr.CommandOutput(os.Args[1], os.Args[1:]...)
	fmt.Print(strings.Join(dr.ExtractSectionsBad(lines), "\n"))

	// TODO: implement
	// fmt.Print(strings.Join(dr.ExtractSectionsBetter(lines), "\n"))
	// fmt.Print(strings.Join(dr.ExtractSectionsBetterPOSIX(lines), "\n"))
}
