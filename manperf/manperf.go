package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	lines := commandOutput(os.Args[1], os.Args[1:]...)
	//fmt.Print(strings.Join(extractSectionsBad(lines), "\n"))
	fmt.Print(strings.Join(extractSectionsBetter(lines), "\n"))
	//fmt.Print(strings.Join(extractSectionsBetterPOSIX(lines), "\n"))
}

func extractSectionsBad(manpageLines []string) []string {
	sections := make([]string, 0)
	for _, line := range manpageLines {
		if ok, _ := regexp.MatchString("([A-Z]{2,})", line); ok {
			r := regexp.MustCompile("([A-Z]{2,})")
			sections = append(sections, r.FindString(line))
		}
	}
	return sections
}

func extractSectionsBetter(manpageLines []string) []string {
	sections := make([]string, 0)
	pattern := regexp.MustCompile(`^([A-Z][A-Z ]+)$`)
	for _, line := range manpageLines {
		section := pattern.FindString(strings.TrimRight(line, "\n"))
		if section != "" {
			sections = append(sections, section)
		}
	}
	return sections
}

func extractSectionsBetterPOSIX(manpageLines []string) []string {
	sections := make([]string, 0)
	pattern := regexp.MustCompilePOSIX(`^([A-Z][A-Z ]+)$`)
	for _, line := range manpageLines {
		section := pattern.FindString(strings.TrimRight(line, "\n"))
		if section != "" {
			sections = append(sections, section)
		}
	}
	return sections
}

func commandOutput(prog string, args ...string) []string {
	cmd := exec.Command(prog, args...)
	out, _ := cmd.StdoutPipe()
	cmd.Start()
	r := bufio.NewReader(out)
	lines := make([]string, 0)
	var err error
	var line string
	for ; err != io.EOF; line, err = r.ReadString('\n') {
		lines = append(lines, line)
	}
	return lines
}
