package highlight

import (
	"bytes"
	"fmt"
	"io"
)

func highlightTest(lang, code string) (string, error) {
	h, err := makeAndHighlight(lang, code)
	if err != nil {
		return "", err
	}
	return h.renderTest()
}

func (h *highlighter) renderTest() (string, error) {
	var buf bytes.Buffer
	h.render(&buf, func(w io.Writer, class string, body []byte) {
		if len(class) > 0 {
			fmt.Fprintf(w, "<%s>%s</%s>", class, body, class)
		} else {
			w.Write(body)
		}
	})
	return buf.String(), nil
}
