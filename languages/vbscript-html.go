package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"vbscript-html", }, `{"subLanguage":["xml"],"contains":[{"begin":"<%","end":"%>","subLanguage":["vbscript"]}]}`)
}