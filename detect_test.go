package highlight

import "testing"

func TestDetect(t *testing.T) {
	cases := []struct {
		code, lang string
	}{
		{
			`
package main

import "fmt"

func main() {
	a := make(chan int)
	a <- 100
	fmt.Println(<- a)
}`,
			"go",
		},
		{"body > #duck.foo { max-width: 100px }", "css"},
		{`<head><!-- foo --></head><body class="amp"><script src="./foo.js"></script></body>`, "xml"},
	}

	for i, c := range cases {
		code := []byte(c.code)
		lang, err := Detect(code)
		if err != nil {
			t.Fatal(err)
		}
		if lang != c.lang {
			highlighted, err := highlightTest(lang, code)
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("%d. Detect(%q) = %q; not %q;\n%s", i, c.code, lang, c.lang, highlighted)
		}
	}
}
