package highlight

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestHighlight(t *testing.T) {
	cases := []struct {
		in, out string
	}{
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
	for i, c := range cases {
		resp, err := Highlight("go", c.in)
		if err != nil {
			t.Error(err)
		}
		if resp != c.out {
			t.Errorf("%d. Highlight(\"go\", %q)\n  = %q;\nnot %q", i, c.in, resp, c.out)
		}
	}
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
