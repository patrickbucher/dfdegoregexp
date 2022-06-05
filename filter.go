package dfdegoregexp

import (
	"bufio"
	"io"
	"regexp"
)

func FilterLines(pattern string, r io.Reader, w io.Writer) {
	p := regexp.MustCompilePOSIX(pattern)
	s := bufio.NewReader(r)
	var l []byte
	var err error
	for ; err != io.EOF; l, err = s.ReadBytes('\n') {
		if p.Match(l) {
			w.Write(l)
		}
	}
}
