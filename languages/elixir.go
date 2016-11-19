package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("elixir", `{"lexemes":"[a-zA-Z_][a-zA-Z0-9_]*(\\!|\\?)?","keywords":"and false then defined module in return redo retry end for true self when next until do begin unless nil break not case cond alias while ensure or include use alias fn quote","contains":[0,1,2,3,4,5,6,7,8,9]}`)
}