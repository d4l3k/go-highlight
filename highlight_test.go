package highlight

import (
	"fmt"
	"strings"
	"testing"
)

type highlightCase struct {
	in, out string
}

func testCases(t *testing.T, lang string, cases []highlightCase) {
	for i, c := range cases {
		resp, err := highlightTest(lang, []byte(c.in))
		if err != nil {
			t.Error(err)
		}
		if string(resp) != c.out {
			t.Errorf("%d. Highlight(%q, %q)\n  = %q;\nnot %q", i, lang, c.in, resp, c.out)
		}
	}
}

func TestHighlight(t *testing.T) {
	cases := []highlightCase{
		{
			`package main`,
			`<keyword>package</keyword> main`,
		},
		{
			`// hello
			`,
			`<comment>// hello</comment>
			`,
		},
		{
			`/* hello */
			`,
			`<comment>/* hello */</comment>
			`,
		},
		{
			`var testfunc`,
			`<keyword>var</keyword> testfunc`,
		},
		{
			`=append()`,
			`=<built_in>append</built_in>()`,
		},
		{
			`
			func main() {
				log.Println("duck")
			}`,
			`
			<keyword>func</keyword> main() {
				log.Println(<string>"duck"</string>)
			}`,
		},
		{
			`/* NOTE: test */`,
			`<comment>/* </comment><doctag>NOTE:</doctag><comment> test */</comment>`,
		},
	}
	testCases(t, "go", cases)
}

func TestHighlightCaseInsensitive(t *testing.T) {
	insensitiveCases := []highlightCase{
		{
			`If ($True -eq $True) { "duck" }`,
			`<keyword>If</keyword> (<variable>$True</variable> -eq <variable>$True</variable>) { <string>"duck"</string> }`,
		},
		{
			`iF ($True -eq $True) { "duck" }`,
			`<keyword>iF</keyword> (<variable>$True</variable> -eq <variable>$True</variable>) { <string>"duck"</string> }`,
		},
	}
	testCases(t, "powershell", insensitiveCases)

	sensitiveCases := []highlightCase{
		{
			`package main`,
			`<keyword>package</keyword> main`,
		},
		{
			`PACKAGE main`,
			`PACKAGE main`,
		},
	}
	testCases(t, "go", sensitiveCases)
}

func TestHighlightMultiLine(t *testing.T) {
	cases := []highlightCase{
		{
			"`\n`",
			"<string>`\n`</string>",
		},
	}
	testCases(t, "go", cases)
}

func TestBeginKeywords(t *testing.T) {
	cases := []highlightCase{
		{
			`select * from duck;`,
			`<keyword>select</keyword> * <keyword>from</keyword> duck;`,
		},
	}
	testCases(t, "sql", cases)
}

func TestAliases(t *testing.T) {
	cases := []highlightCase{
		{
			`package main`,
			`<keyword>package</keyword> main`,
		},
	}

	for _, lang := range []string{"go", "golang"} {
		testCases(t, lang, cases)
	}
}

func BenchmarkHighlight(b *testing.B) {
	n := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, n := range n {
		b.Run(fmt.Sprintf("BenchmarkHighlight%dBytes", n), func(b *testing.B) {
			code := strings.Repeat(" ", n)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				if _, err := HTML("go", []byte(code)); err != nil {
					b.Error(err)
				}
			}
		})
	}
}
