package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("clojure-repl", `{"contains":[{"className":"meta","begin":"^([\\w.-]+|\\s*#_)=>","starts":{"end":"$","subLanguage":"clojure"}}]}`)
}