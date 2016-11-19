package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("json", `{"contains":[0,1,2,3],"keywords":{"literal":"true false null"},"illegal":"\\S"}`)
}