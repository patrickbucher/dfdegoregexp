package dfdegoregexp

import (
	"bufio"
	"errors"
	"io"
	"regexp"
)

const functionDeclaration = `` // TODO

func FilterLines(r io.Reader, w io.Writer) {
	p := regexp.MustCompilePOSIX(functionDeclaration)
	s := bufio.NewReader(r)
	var l []byte
	var err error
	for ; !errors.Is(err, io.EOF); l, err = s.ReadBytes('\n') {
		if p.Match(l) {
			w.Write(l)
		}
	}
}
