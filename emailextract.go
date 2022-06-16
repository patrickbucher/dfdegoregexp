package dfdegoregexp

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// TODO: This regexp must be written. Figure out group names according to switch/case below.
	r = `^(?P<first>[a-z]+)\.?(?P<last>[a-z]+)?(?P<year>[0-9]{2,4})?@(?P<comp>[a-z]+)\.[a-z]+`
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
			ei.first = strings.Title(matches[i])
		case "last":
			ei.last = strings.Title(matches[i])
		case "year":
			ei.year, _ = strconv.Atoi(matches[i])
		case "comp":
			ei.company = strings.ToUpper(matches[i])
		}
	}
	return ei.String()
}
