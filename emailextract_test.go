package dfdegoregexp

import "testing"

type testCase struct {
	email, desc string
}

var tests = []testCase{
	{"joey@foobar.com", "Joey, FOOBAR"},
	{"harry.callahan@sfpd.gov", "Harry Callahan, SFPD"},
	{"homer.simpson69@aol.com", "Homer Simpson, *1969, AOL"},
	{"stan.marsh2012@southpark.com", "Stan Marsh, *2012, SOUTHPARK"},
	{"julius.caesar@rom.it", "Julius Caesar, ROM"},
}

func TestEmailExtract(t *testing.T) {
	for _, test := range tests {
		expected := test.desc
		actual := Extract(test.email)
		if actual != expected {
			t.Errorf(`Extract("%s"): expected "%s", got "%s"`, test.email, expected, actual)
		}
	}
}
