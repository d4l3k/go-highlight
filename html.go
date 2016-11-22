package highlight

import (
	"bytes"
	"fmt"
	"html"
	"io"
)

// HTML highlights a piece of code in HTML.
func HTML(lang, code string) (string, error) {
	h, err := makeAndHighlight(lang, code)
	if err != nil {
		return "", err
	}
	return h.renderHTML()
}

func (h *highlighter) renderHTML() (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(`<div class="highlight"><pre><code>`))
	h.render(&buf, func(w io.Writer, class string, body []byte) {
		escaped := html.EscapeString(string(body))
		if len(class) > 0 {
			fmt.Fprintf(w, "<span class=\"%s\">%s</span>", class, escaped)
		} else {
			w.Write([]byte(escaped))
		}
	})
	buf.Write([]byte(`</code></pre></div>`))
	return buf.String(), nil
}
