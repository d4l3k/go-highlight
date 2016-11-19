package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("diff", `{"aliases":["patch"],"contains":[{"className":"meta","relevance":10,"variants":[{"begin":"^@@ +\\-\\d+,\\d+ +\\+\\d+,\\d+ +@@$"},{"begin":"^\\*\\*\\* +\\d+,\\d+ +\\*\\*\\*\\*$"},{"begin":"^\\-\\-\\- +\\d+,\\d+ +\\-\\-\\-\\-$"}]},{"className":"comment","variants":[{"begin":"Index: ","end":"$"},{"begin":"={3,}","end":"$"},{"begin":"^\\-{3}","end":"$"},{"begin":"^\\*{3} ","end":"$"},{"begin":"^\\+{3}","end":"$"},{"begin":"\\*{5}","end":"\\*{5}$"}]},{"className":"addition","begin":"^\\+","end":"$"},{"className":"deletion","begin":"^\\-","end":"$"},{"className":"addition","begin":"^\\!","end":"$"}]}`)
}