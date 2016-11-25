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
		count int
		err   error
	}
	resultChan := make(chan result)

	go func() {
		for i, lang := range registry.Languages() {
			languageChan <- lang
			log.Printf("processed %d: %s", i, lang)
		}
		log.Printf("done")
		close(languageChan)
	}()

	var wg sync.WaitGroup

	for i := 0; i < Workers; i++ {
		wg.Add(1)
		go func() {
			for l := range languageChan {
				h, err := makeAndHighlight(l, code)
				resultChan <- result{l, len(h.highlights), err}
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
	var bestLangVal int
	for res := range resultChan {
		if res.err != nil {
			err = res.err
			continue
		}

		if res.count > bestLangVal {
			bestLang = res.lang
			bestLangVal = res.count
		}
	}

	return bestLang, err
}
