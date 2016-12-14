package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"dust", "dst"}, `{"aliases":["dst"],"case_insensitive":true,"subLanguage":["xml"],"contains":[{"className":"template-tag","begin":"\\{[#\\/]","end":"\\}","illegal":";","contains":[{"className":"name","begin":"[a-zA-Z\\.-]+","starts":{"endsWithParent":true,"relevance":0,"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]}}]},{"className":"template-variable","begin":"\\{","end":"\\}","illegal":";","keywords":"if eq ne lt lte gt gte select default math sep"}]}`)
}