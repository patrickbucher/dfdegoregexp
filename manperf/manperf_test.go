package main

import "testing"

// FIXME: this is from Arch Linux, use Debian version for forum post
var expectedSections = []string{"NAME", "SYNOPSIS", "DESCRIPTION", "EXAMPLES",
	"OVERVIEW", "DEFAULTS", "OPTIONS", "EXIT STATUS", "ENVIRONMENT", "FILES",
	"SEE ALSO", "HISTORY", "BUGS"}

func TestExtractSectionsBad(t *testing.T) {
	lines := commandOutput("man", "man")
	actual := extractSectionsBad(lines)
	for i, e := range expectedSections {
		if actual[i] != e {
			t.Errorf("expected actual[%d]==%s, was %s", i, e, actual[i])
		}
	}
}

func TestExtractSectionsBetter(t *testing.T) {
	lines := commandOutput("man", "man")
	actual := extractSectionsBetter(lines)
	for i, e := range expectedSections {
		if actual[i] != e {
			t.Errorf("expected actual[%d]==%s, was %s", i, e, actual[i])
		}
	}
}

func TestExtractSectionsBetterPOSIX(t *testing.T) {
	lines := commandOutput("man", "man")
	actual := extractSectionsBetterPOSIX(lines)
	for i, e := range expectedSections {
		if actual[i] != e {
			t.Errorf("expected actual[%d]==%s, was %s", i, e, actual[i])
		}
	}
}

// run only benchmark: go test -bench . -run ^$

func BenchmarkExtractSectionsBad(b *testing.B) {
	b.StopTimer()
	lines := commandOutput("man", "man")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		extractSectionsBad(lines)
	}
}

func BenchmarkExtractSectionsBetter(b *testing.B) {
	b.StopTimer()
	lines := commandOutput("man", "man")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		extractSectionsBetter(lines)
	}
}

func BenchmarkExtractSectionsBetterPOSIX(b *testing.B) {
	b.StopTimer()
	lines := commandOutput("man", "man")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		extractSectionsBetterPOSIX(lines)
	}
}
