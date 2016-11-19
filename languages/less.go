package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("less", `{"case_insensitive":true,"illegal":"[=>'/<($\"]","contains":[0,1,2,3,4,5]}`)
}