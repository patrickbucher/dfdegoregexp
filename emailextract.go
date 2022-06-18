package dfdegoregexp

import (
	"fmt"
	"regexp"
)

var (
	// TODO: This regexp must be written. Figure out group names according to switch/case below.
	r = ``
	p = regexp.MustCompile(r)
)

type emailInfo struct {
	first, last, company string
	year                 int
}

func (e emailInfo) String() string {
	if e.first == "" || e.company == "" {
		return ""
	}
	var y int
	if e.year != 0 {
		if e.year >= 100 {
			y = e.year
		} else {
			y = 1900 + e.year
		}
	}
	if e.last != "" && y != 0 {
		return fmt.Sprintf("%s %s, *%d, %s", e.first, e.last, y, e.company)
	}
	if e.last != "" && y == 0 {
		return fmt.Sprintf("%s %s, %s", e.first, e.last, e.company)
	}
	if e.last == "" && y == 0 {
		return fmt.Sprintf("%s, %s", e.first, e.company)
	}
	return ""
}

func Extract(email string) string {
	matches := p.FindStringSubmatch(email)
	if len(matches) == 0 {
		return ""
	}
	var ei emailInfo
	for i, name := range p.SubexpNames() {
		switch name {
		case "first":
			// TODO
		case "last":
			// TODO
		case "year":
			// TODO
		case "comp":
			// TODO
		}
	}
	return ei.String()
}
