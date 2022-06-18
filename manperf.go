package dfdegoregexp

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"regexp"
)

func ExtractSectionsBad(manpageLines []string) []string {
	sections := make([]string, 0)
	for _, line := range manpageLines {
		if ok, _ := regexp.MatchString("([A-Z]{2,})", line); ok {
			r := regexp.MustCompile("([A-Z]{2,})")
			sections = append(sections, r.FindString(line))
		}
	}
	return sections
}

func ExtractSectionsBetter(manpageLines []string) []string {
	// TODO: implement function
	return make([]string, 0)
}

func CommandOutput(prog string, args ...string) []string {
	cmd := exec.Command(prog, args...)
	out, _ := cmd.StdoutPipe()
	cmd.Start()
	r := bufio.NewReader(out)
	lines := make([]string, 0)
	var err error
	var line string
	for ; !errors.Is(err, io.EOF); line, err = r.ReadString('\n') {
		lines = append(lines, line)
	}
	return lines
}
