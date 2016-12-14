package highlight

import (
	"sync"

	"github.com/d4l3k/go-highlight/registry"
	pcre "github.com/gijsbers/go-pcre"
)

// Workers is the number of workers to use to detect the language.
var Workers = 1

// Detect returns the detected language.
func Detect(code []byte) (string, error) {
	languages := registry.Languages()
	return detect(code, languages, nil)
}

// detect returns the detected language in languages. If no language is detected
// it returns an empty string.
func detect(code []byte, languages []string, end *pcre.Regexp) (string, error) {
	languageChan := make(chan string)
	type result struct {
		lang  string
		score float64
		err   error
	}
	resultChan := make(chan result)

	go func() {
		for _, lang := range languages {
			languageChan <- lang
		}
		close(languageChan)
	}()

	var wg sync.WaitGroup

	for i := 0; i < Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for l := range languageChan {
				h, err := makeHighlighter(l, code)
				if err != nil {
					resultChan <- result{l, 0, err}
					continue
				}
				if _, err := h.highlight(h.lang.Contains, 0, end); err != nil {
					resultChan <- result{l, 0, err}
					continue
				}

				result := result{l, 0, err}

				for _, h := range h.highlights {
					result.score += h.contains.Relevance
				}

				resultChan <- result
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var err error
	var bestLang string
	var bestLangVal float64
	for res := range resultChan {
		if res.err != nil {
			err = res.err
			continue
		}

		if res.score > bestLangVal {
			bestLang = res.lang
			bestLangVal = res.score
		}
	}

	return bestLang, err
}
