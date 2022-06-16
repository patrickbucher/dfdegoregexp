package dfdegoregexp

import (
	"bytes"
	"strings"
	"testing"
)

const (
	goDocOSExcerpt = `
const PathSeparator = '/' ...
const ModeDir = fs.ModeDir ...
const DevNull = "/dev/null"
var ErrInvalid = fs.ErrInvalid ...
var Stdin = NewFile(uintptr(syscall.Stdin), "/dev/stdin") ...
var Args []string
var ErrProcessDone = errors.New("os: process already finished")
func Chdir(dir string) error
func Chmod(name string, mode FileMode) error
func Chown(name string, uid, gid int) error
`
	goDocOSFuncs = `
func Chdir(dir string) error
func Chmod(name string, mode FileMode) error
func Chown(name string, uid, gid int) error
`
)

func TestFilterFuncLines(t *testing.T) {
	source := bytes.NewBufferString(goDocOSExcerpt)
	sink := bytes.NewBufferString("")
	FilterLines(source, sink)
	got := strings.TrimSpace(sink.String())
	expected := strings.TrimSpace(goDocOSFuncs)
	if got != expected {
		t.Errorf("filter lines by pattern '%s'\ngot:\n%s\nexpected:\n%s\n",
			functionDeclaration, got, expected)
	}
}
