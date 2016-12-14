package highlight

import (
	"log"
	"sync"

	"github.com/d4l3k/go-highlight/registry"
)

// Workers is the number of workers to use to detect the language.
var Workers = 1

// Detect returns the detected language.
func Detect(code []byte) (string, error) {
	languageChan := make(chan string)
	type result struct {
		lang  string
		score float64
		err   error
	}
	resultChan := make(chan result)

	go func() {
		for _, lang := range registry.Languages() {
			languageChan <- lang
		}
		close(languageChan)
	}()

	var wg sync.WaitGroup

	for i := 0; i < Workers; i++ {
		wg.Add(1)
		go func() {
			for l := range languageChan {
				h, err := makeAndHighlight(l, code)
				result := result{l, 0, err}

				for _, h := range h.highlights {
					if l == "apache" || l == "xml" {
						log.Printf("contains %q, rel %f", h.class, h.contains.Relevance)
					}
					result.score += h.contains.Relevance
				}

				resultChan <- result
			}
			wg.Done()
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

		if res.lang == "apache" || res.lang == "xml" {
			log.Printf("lang %q, score %f", res.lang, res.score)
		}

		if res.score > bestLangVal {
			bestLang = res.lang
			bestLangVal = res.score
		}
	}

	return bestLang, err
}
