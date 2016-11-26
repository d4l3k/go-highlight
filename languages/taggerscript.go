package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"taggerscript", }, `{"contains":[{"className":"comment","begin":"\\$noop\\(","end":"\\)","contains":[{"begin":"\\(","end":"\\)","contains":[{"Ref":["contains","0","contains","0"]},{"begin":"\\\\."}]}],"relevance":10},{"className":"keyword","begin":"\\$(?!noop)[a-zA-Z][_a-zA-Z0-9]*","end":"\\(","excludeEnd":true},{"className":"variable","begin":"%[_a-zA-Z0-9:]*","end":"%"},{"className":"symbol","begin":"\\\\."}]}`)
}