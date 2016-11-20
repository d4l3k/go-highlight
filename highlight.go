package highlight

import (
	"bytes"
	"container/heap"
	"errors"
	"io"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/d4l3k/go-highlight/registry"

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

func parseWords(words string) []string {
	if len(words) == 0 {
		return nil
	}
	return strings.Split(words, " ")
}

func parseKeywords(kw *registry.Keywords) map[string][]string {
	if kw == nil {
		return map[string][]string{}
	}
	return map[string][]string{
		"keyword":  kw.Keyword,
		"literal":  kw.Literal,
		"built_in": kw.BuiltIn,
	}
}

type highlighter struct {
	code       []byte
	lang       registry.Language
	highlights []highlight
	basics     map[string][]string
}

func makeHighlighter(lang, code string) (highlighter, error) {
	langDef, err := registry.Lookup(lang)
	if err != nil {
		return highlighter{}, err
	}
	//spew.Dump(langDef)

	return highlighter{
		code:   []byte(code),
		lang:   langDef,
		basics: parseKeywords(langDef.Keywords),
	}, nil
}

func (h *highlighter) wordsMatch(view []byte, words []string) (string, bool, error) {
	for _, word := range words {
		if len(word) > len(view) {
			continue
		}

		matched := len(word) == len(view) || !isWord(view[len(word)])
		if h.lang.CaseInsensitive {
			matched = matched && bytes.EqualFold([]byte(word), view[:len(word)])
		} else {
			matched = matched && bytes.Equal([]byte(word), view[:len(word)])
		}

		if matched {
			return word, true, nil
		}
	}
	return "", false, nil
}

func isWord(a byte) bool {
	b := rune(a)
	return b == '_' || unicode.IsLetter(b) || unicode.IsNumber(b)
}

func (h *highlighter) matchKeywords(start *int, view []byte, typ string, words []string) (bool, error) {
	word, matched, err := h.wordsMatch(view, words)
	if err != nil {
		return false, err
	}
	if matched {
		h.addHighlight(typ, *start, *start+len(word))
		*start += len(word)
		return true, nil
	}
	return false, nil
}

func (h *highlighter) highlight(mode []*registry.Contains, start int, end *regexp.Regexp) (int, error) {
	root := start == 0

outer:
	for start < len(h.code) {
		view := h.code[start:]

		isWordBoundary := start == 0
		if start > 0 {
			isWordBoundary = !isWord(h.code[start-1]) && isWord(h.code[start])
		}

		// Check for the end of the previous section.
		if end != nil && end.Match(view) {
			return start, nil
		}

		for _, c := range mode {
			// Highlight basic keywords, literals and built_ins.
			if isWordBoundary {
				keywords := []map[string][]string{}
				if c.Keywords != nil {
					keywords = append(keywords, parseKeywords(c.Keywords))
				}
				if root {
					keywords = append(keywords, h.basics)
				}
				for _, kw := range keywords {
					for typ, words := range kw {
						cont, err := h.matchKeywords(&start, view, typ, words)
						if err != nil {
							return 0, err
						}
						if cont {
							continue outer
						}
					}
				}
			}

			for _, v := range append([]*registry.Contains{c}, c.Variants...) {
				var beginIndex []int
				if v.Begin != nil && len(c.ClassName) > 0 {
					beginIndex = v.Begin.FindIndex(view)
				} else if isWordBoundary && len(v.BeginKeywords) > 0 {
					word, matched, err := h.wordsMatch(view, v.BeginKeywords)
					if err != nil {
						return 0, err
					}
					if matched {
						h.addHighlight("keyword", start, start+len(word))
						beginIndex = []int{0, len(word)}
					}
				} else {
					continue
				}

				if beginIndex == nil {
					continue
				}

				// Simple Begin only matches
				if v.End == nil {
					if len(c.ClassName) > 0 {
						h.addHighlight(c.ClassName, start, start+beginIndex[1])
					}
					start += beginIndex[1]
					continue
				}

				// Highlight subsections.
				newStart, err := h.highlight(c.Contains, start+beginIndex[1], v.End)
				if err != nil {
					return 0, err
				}

				// Avoid matching start of section.
				endView := h.code[newStart:]
				index := v.End.FindIndex(endView)
				if index == nil {
					return 0, errors.New("can't find ending")
				}
				newStart += index[1]
				if len(c.ClassName) > 0 {
					h.addHighlight(c.ClassName, start, newStart)
				}
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
		content: string(h.code[start:end]),
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
	//spew.Dump(h.highlights)
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
