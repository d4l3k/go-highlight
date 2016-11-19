package highlight

import (
	"container/heap"
	"reflect"
	"testing"
)

type highlightCase struct {
	in, out string
}

func testCases(t *testing.T, lang string, cases []highlightCase) {
	for i, c := range cases {
		resp, err := Highlight(lang, c.in)
		if err != nil {
			t.Error(err)
		}
		if resp != c.out {
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

func TestCaseInsensitive(t *testing.T) {
	cases := []highlightCase{
		{
			`If ($True -eq $True) { "duck" }`,
			`<keyword>If</keyword> (<variable>$True</variable> -eq <variable>$True</variable>) { <string>"duck"</string> }`,
		},
		{
			`iF ($True -eq $True) { "duck" }`,
			`<keyword>iF</keyword> (<variable>$True</variable> -eq <variable>$True</variable>) { <string>"duck"</string> }`,
		},
	}
	testCases(t, "powershell", cases)
}

func TestPOIHeap(t *testing.T) {
	pois := []poi{
		{i: 100},
		{i: 15},
		{i: 15, start: true},
		{i: 10},
		{i: 5},
	}
	hp := &poiHeap{}
	for _, p := range pois {
		heap.Push(hp, p)
	}
	for i, p := range pois {
		h := hp.Peek()
		if !reflect.DeepEqual(p, h) {
			t.Errorf("%d. Peek() expected %+v = %+v", i, p, h)
		}

		h = heap.Pop(hp).(poi)
		if !reflect.DeepEqual(p, h) {
			t.Errorf("%d. Pop() expected %+v = %+v", i, p, h)
		}
	}
}
