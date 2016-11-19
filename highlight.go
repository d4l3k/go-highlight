package highlight

import (
	"bytes"
	"container/heap"
	"errors"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/d4l3k/go-highlight/registry"
	"github.com/davecgh/go-spew/spew"

	// Import for language registration side-effect.
	_ "github.com/d4l3k/go-highlight/languages"
)

// Highlight highlights a piece of code.
func Highlight(lang, code string) (string, error) {
	h, err := makeHighlighter(lang, code)
	if err != nil {
		return "", nil
	}
	if _, err := h.highlight(h.lang.Contains, 0, nil); err != nil {
		return "", err
	}
	return h.renderTest()
}

type highlight struct {
	start, end int
	class      string
	content    string
}

type highlighter struct {
	code       string
	lang       registry.Language
	highlights []highlight
}

func makeHighlighter(lang, code string) (highlighter, error) {
	langDef, err := registry.Lookup(lang)
	if err != nil {
		return highlighter{}, err
	}
	spew.Dump(langDef)

	return highlighter{code: code, lang: langDef}, nil
}

func (h *highlighter) highlight(mode []registry.Contains, start int, end *regexp.Regexp) (int, error) {
	basic := map[string][]string{}
	if start == 0 {
		basic["keyword"] = strings.Split(h.lang.Keywords.Keyword, " ")
		basic["literal"] = strings.Split(h.lang.Keywords.Literal, " ")
		basic["built_in"] = strings.Split(h.lang.Keywords.BuiltIn, " ")
	}
outer:
	for start < len(h.code) {
		view := h.code[start:]

		// Check for the end of the previous section.
		if end != nil && end.MatchString(view) {
			return start, nil
		}

		// Highlight basic keywords, literals and built_ins.
		for typ, words := range basic {
			for _, word := range words {
				matched, err := regexp.MatchString("^\\b"+word+"\\b", view)
				if err != nil {
					return 0, err
				}
				if matched {
					h.addHighlight(typ, start, start+len(word))
					start += len(word)
					continue outer
				}
			}
		}

		for _, c := range mode {
			// Skip non class contains. Typically used for boosting relevance.
			if len(c.ClassName) == 0 {
				continue
			}

			for _, v := range append([]registry.Contains{c}, c.Variants...) {
				if len(v.Begin) == 0 {
					continue
				}

				beginRegex, err := regexp.Compile("^" + v.Begin)
				if err != nil {
					return 0, err
				}
				beginIndex := beginRegex.FindStringIndex(view)
				if beginIndex == nil {
					continue
				}

				// Simple Begin only matches
				if len(v.End) == 0 {
					h.addHighlight(c.ClassName, start, start+beginIndex[1])
					start += beginIndex[1]
					continue
				}

				// Regex needs to be in multi line mode and match starting at the
				// beginning of the string.
				endRegex, err := regexp.Compile(`^(?m:` + v.End + `)`)
				if err != nil {
					return 0, err
				}

				// Highlight subsections.
				newStart, err := h.highlight(c.Contains, start+beginIndex[1], endRegex)
				if err != nil {
					return 0, err
				}

				// Avoid matching start of section.
				endView := h.code[newStart:]
				index := endRegex.FindStringIndex(endView)
				if index == nil {
					return 0, errors.New("can't find ending")
				}
				newStart += index[1]
				h.addHighlight(c.ClassName, start, newStart)
				start = newStart
				continue outer
			}

		}
		start++
	}
	return start - 1, nil
}

func (h *highlighter) addHighlight(class string, start, end int) {
	h.highlights = append(h.highlights, highlight{
		start:   start,
		end:     end,
		class:   class,
		content: h.code[start:end],
	})
}

// toPOI returns a slice of poi elements sorted according to i and then start
func (h *highlighter) toPOI() []poi {
	pois := make([]poi, len(h.highlights)*2)
	for i, h := range h.highlights {
		pois[i*2] = poi{
			i:     h.start,
			start: true,
			class: h.class,
		}
		pois[i*2+1] = poi{
			i:     h.end,
			start: false,
			class: h.class,
		}
	}
	sort.Sort(sort.Reverse(poiHeap(pois)))
	return pois
}

func (h *highlighter) render(w io.Writer, f func(class string, start bool) string) {
	pois := h.toPOI()
	max := &poiHeap{}
	i := 0
	for _, p := range pois {
		if p.start {
			oldMax := p
			if max.Len() > 0 {
				oldMax = max.Peek()
			}
			heap.Push(max, p)
			if max.Peek() == p {
				w.Write([]byte(h.code[i:p.i]))
				i = p.i
				if max.Len() > 1 {
					w.Write([]byte(f(oldMax.class, false)))
				}
				w.Write([]byte(f(p.class, true)))
			}
		} else {
			oldMax := max.Peek()
			if oldMax.class == p.class {
				w.Write([]byte(h.code[i:p.i]))
				i = p.i
				w.Write([]byte(f(p.class, false)))
				heap.Pop(max)
				if max.Len() > 0 {
					w.Write([]byte(f(max.Peek().class, true)))
				}
			}
		}
	}
	w.Write([]byte(h.code[i:]))
}

func (h *highlighter) renderTest() (string, error) {
	spew.Dump(h.highlights)
	var buf bytes.Buffer
	h.render(&buf, func(class string, start bool) string {
		if start {
			return "<" + class + ">"
		}
		return "</" + class + ">"
	})
	return buf.String(), nil
}

type poi struct {
	i     int
	start bool
	class string
}

type poiHeap []poi

func (h poiHeap) Len() int { return len(h) }
func (h poiHeap) Less(i, j int) bool {
	if h[i].i == h[j].i {
		return !h[i].start && h[j].start
	}
	return -h[i].i < -h[j].i
}
func (h poiHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *poiHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(poi))
}

func (h *poiHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *poiHeap) Peek() poi {
	return (*h)[0]
}
