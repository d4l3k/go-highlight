package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("nix", `{"aliases":[0],"keywords":{"keyword":"rec with let in inherit assert if else then","literal":"true false or and null","built_in":"import abort baseNameOf dirOf isNull builtins map removeAttrs throw toString derivation"},"contains":[0,1,2,3,4]}`)
}