package highlight

import "testing"

func TestDetect(t *testing.T) {
	cases := []struct {
		code, lang string
	}{
		{
			`package main
		func main() {
		}`,
			"go",
		},
	}

	for i, c := range cases {
		lang, err := Detect([]byte(c.code))
		if err != nil {
			t.Fatal(err)
		}
		if lang != c.lang {
			t.Errorf("%d. Detect(%q) = %q; not %q", i, c.code, lang, c.lang)
		}
	}
}
