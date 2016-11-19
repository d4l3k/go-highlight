package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("erlang", `{"aliases":[0],"keywords":{"keyword":"after and andalso|10 band begin bnot bor bsl bzr bxor case catch cond div end fun if let not of orelse|10 query receive rem try when xor","literal":"false true"},"illegal":"(</|\\*=|\\+=|-=|/\\*|\\*/|\\(\\*|\\*\\))","contains":[0,1,2,3,4,5,6,7,8,9]}`)
}