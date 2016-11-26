package highlight

import (
	"bytes"
	"container/heap"
	"io"
	"sort"
	"strings"
	"unicode"

	"github.com/d4l3k/go-highlight/registry"
	pcre "github.com/gijsbers/go-pcre"

	// Import for language registration side-effect.
	_ "github.com/d4l3k/go-highlight/languages"
)

// The text highlight classes.
const (
	Keyword = "keyword"
	Literal = "literal"
	BuiltIn = "built_in"
)

func makeAndHighlight(lang string, code []byte) (highlighter, error) {
	h, err := makeHighlighter(lang, code)
	if err != nil {
		return highlighter{}, err
	}
	if _, err := h.highlight(h.lang.Contains, 0, nil); err != nil {
		return highlighter{}, err
	}
	return h, nil
}

type highlight struct {
	start, end int
	class      string
	contains   *registry.Contains
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
		Keyword: kw.Keyword,
		Literal: kw.Literal,
		BuiltIn: kw.BuiltIn,
	}
}

type highlighter struct {
	langName   string
	code       []byte
	lang       registry.Contains
	highlights []highlight
	basics     map[string][]string

	// indexCache is a cache for FindIndex results.
	indexCache map[*pcre.Regexp][]int
}

func makeHighlighter(lang string, code []byte) (highlighter, error) {
	langDef, err := registry.Lookup(lang)
	if err != nil {
		return highlighter{}, err
	}

	return highlighter{
		langName:   lang,
		code:       []byte(code),
		lang:       langDef,
		basics:     parseKeywords(langDef.Keywords),
		indexCache: map[*pcre.Regexp][]int{},
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
		h.addHighlight(typ, *start, *start+len(word), nil)
		*start += len(word)
		return true, nil
	}
	return false, nil
}

// findIndex uses the index cache to avoid doing numerous regex lookups.
func (h *highlighter) findIndex(r *pcre.Regexp, view []byte, start int) []int {
	idx, ok := h.indexCache[r]
	if ok {
		if idx == nil {
			return nil
		} else if idx[0] >= start {
			return []int{idx[0] - start, idx[1] - start}
		}
	}
	idx = r.FindIndex(view, 0)
	if idx == nil {
		h.indexCache[r] = nil
	} else {
		h.indexCache[r] = []int{idx[0] + start, idx[1] + start}
	}
	return idx
}

func (h *highlighter) highlight(mode []*registry.Contains, start int, end *pcre.Regexp) (int, error) {
	root := start == 0

outer:
	for start < len(h.code) {
		view := h.code[start:]

		isWordBoundary := start == 0
		if start > 0 {
			isWordBoundary = !isWord(h.code[start-1]) && isWord(h.code[start])
		}

		// Check for the end of the previous section.
		if end != nil {
			index := h.findIndex(end, view, start)
			if index != nil && index[0] == 0 {
				return start + index[1], nil
			}
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
					beginIndex = h.findIndex(v.Begin, view, start)
				} else if isWordBoundary && len(v.BeginKeywords) > 0 {
					word, matched, err := h.wordsMatch(view, v.BeginKeywords)
					if err != nil {
						return 0, err
					}
					if matched {
						h.addHighlight("keyword", start, start+len(word), nil)
						beginIndex = []int{0, len(word)}
					}
				} else {
					continue
				}

				if beginIndex == nil || beginIndex[0] != 0 {
					continue
				}

				// Simple Begin only matches
				if v.End == nil {
					if len(c.ClassName) > 0 {
						h.addHighlight(c.ClassName, start, start+beginIndex[1], c)
					}
					start += beginIndex[1]
					continue
				}

				// Highlight subsections.
				newStart, err := h.highlight(c.Contains, start+beginIndex[1], v.End)
				if err != nil {
					return 0, err
				}

				if len(c.ClassName) > 0 {
					h.addHighlight(c.ClassName, start, newStart, c)
				}
				start = newStart
				continue outer
			}

		}
		start++
	}
	return start, nil
}

func (h *highlighter) addHighlight(class string, start, end int, c *registry.Contains) {
	h.highlights = append(h.highlights, highlight{
		start:    start,
		end:      end,
		class:    class,
		contains: c,
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

func (h *highlighter) render(w io.Writer, f func(w io.Writer, class string, body []byte)) {
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
				class := ""
				if max.Len() > 1 {
					class = oldMax.class
				}
				f(w, class, h.code[i:p.i])
				i = p.i
			}
		} else {
			oldMax := max.Peek()
			if oldMax.class == p.class {
				f(w, p.class, h.code[i:p.i])
				i = p.i
				heap.Pop(max)
			}
		}
	}
	f(w, "", h.code[i:])
}
