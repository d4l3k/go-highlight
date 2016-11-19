package highlight

import (
	"log"
	"strings"

	"github.com/d4l3k/go-highlight/registry"
	"github.com/davecgh/go-spew/spew"

	// Import for language registration side-effect.
	_ "github.com/d4l3k/go-highlight/languages"
)

// Highlight highlights a piece of code.
func Highlight(lang, code string) (string, error) {
	langDef, err := registry.Lookup(lang)
	if err != nil {
		return "", err
	}
	log.Printf("lang: %q", lang)
	spew.Dump(langDef)

	h := highlighter{code: code, lang: langDef}
	h.highlight(langDef.Contains, 0, len(code))
	spew.Dump(h.highlights)
	return "", nil
}

type highlight struct {
	start, end int
	class      string
}

type highlighter struct {
	code       string
	lang       registry.Language
	highlights []highlight
}

func (h *highlighter) highlight(mode []registry.Contains, start, end int) {
	basic := map[string][]string{}
	if start == 0 {
		basic["keyword"] = strings.Split(h.lang.Keywords.Keyword, " ")
		basic["literal"] = strings.Split(h.lang.Keywords.Literal, " ")
		basic["built_in"] = strings.Split(h.lang.Keywords.BuiltIn, " ")
	}
outer:
	for start < end {
		view := h.code[start:]
		for typ, words := range basic {
			for _, word := range words {
				if strings.HasPrefix(view, word) {
					h.highlights = append(h.highlights, highlight{start: start, end: start + len(word), class: typ})
					start += len(word)
					continue outer
				}
			}
		}
		start++
	}
}
