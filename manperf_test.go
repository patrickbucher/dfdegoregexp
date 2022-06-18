package dfdegoregexp

import (
	"testing"
)

var expectedSections = []string{"NAME", "SYNOPSIS", "DESCRIPTION", "EXAMPLES",
	"OVERVIEW", "DEFAULTS", "OPTIONS", "EXIT STATUS", "ENVIRONMENT", "FILES",
	"SEE ALSO", "HISTORY"}

func TestExtractSectionsBad(t *testing.T) {
	lines := CommandOutput("man", "man")
	actual := ExtractSectionsBad(lines)
	for i, e := range expectedSections {
		if actual[i] != e {
			t.Errorf("expected actual[%d]==%s, was %s", i, e, actual[i])
		}
	}
}

func TestExtractSectionsBetter(t *testing.T) {
	lines := CommandOutput("man", "man")
	actual := ExtractSectionsBetter(lines)
	for i, e := range expectedSections {
		if actual[i] != e {
			t.Errorf("expected actual[%d]==%s, was %s", i, e, actual[i])
		}
	}
}

func BenchmarkExtractSectionsBad(b *testing.B) {
	b.StopTimer()
	lines := CommandOutput("man", "man")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ExtractSectionsBad(lines)
	}
}

func BenchmarkExtractSectionsBetter(b *testing.B) {
	b.StopTimer()
	lines := CommandOutput("man", "man")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ExtractSectionsBetter(lines)
	}
}
