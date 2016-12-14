package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"tap", }, `{"case_insensitive":true,"contains":[{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"meta","variants":[{"begin":"^TAP version (\\d+)$"},{"begin":"^1\\.\\.(\\d+)$"}]},{"begin":"(s+)?---$","end":"\\.\\.\\.$","subLanguage":["yaml"],"relevance":0},{"className":"number","begin":" (\\d+) "},{"className":"symbol","variants":[{"begin":"^ok"},{"begin":"^not ok"}]}]}`)
}