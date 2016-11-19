package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("vbscript-html", `{"subLanguage":"xml","contains":[{"begin":"<%","end":"%>","subLanguage":"vbscript"}]}`)
}