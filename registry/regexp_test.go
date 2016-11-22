package registry

import (
	"regexp"
	"testing"
)

func TestIsRegexp(t *testing.T) {
	cases := []struct {
		a, clean string
		regex    bool
	}{
		{`asdf12341!!!A`, `asdf12341!!!A`, false},
		{`//`, `//`, false},
		{`()`, "", true},
		{`\(\)`, "()", false},
		{`$`, "", true},
		{`\b`, "", true},
		{`:(:)?[a-zA-Z0-9\_\-\+\(\)"'.]+`, "", true},
	}

	for i, c := range cases {
		clean, out := cleanRegexp(c.a)

		if out != c.regex || clean != c.clean {
			t.Errorf("%d. isRegex(%q) = %q, %v; not %q, %v", i, c.a, clean, out, c.clean, c.regex)
		}

		if !out {
			r, err := regexp.Compile("^" + c.a + "$")
			if err != nil {
				t.Fatal(err)
			}
			if !r.MatchString(clean) {
				t.Errorf("%d. non-regexp should match themselves; %+v does not", i, c.a)
			}
		}
	}
}

func TestStringFinder(t *testing.T) {
	cases := []struct {
		find, body string
		found      bool
	}{
		{`duck`, `asdfa duck asdf`, true},
		{`duck2`, `asdfa duck asdf`, false},
		{``, `kasdfa`, true},
		{`duck`, ``, false},
	}

	for i, c := range cases {
		idx := StringFinder(c.find).FindIndex([]byte(c.body))
		if (idx != nil) != c.found {
			t.Errorf("%d. StringFinder(%q).FindIndex(%q) found status wrong = %v", i, c.find, c.body, !c.found)
		}

		if c.found {
			out := c.body[idx[0]:idx[1]]
			if out != c.find {
				t.Errorf("%d. StringFinder(%q).FindIndex(%q) = %q; not %q", i, c.find, c.body, out, c.find)
			}
		}
	}
}
