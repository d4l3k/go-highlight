package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"fix", }, `{"contains":[{"begin":"[^␁\u0001]+","end":"[␁\u0001]","excludeEnd":true,"returnBegin":true,"returnEnd":false,"contains":[{"begin":"([^␁\u0001=]+)","end":"=([^␁\u0001=]+)","returnEnd":true,"returnBegin":false,"className":"attr"},{"begin":"=","end":"([␁\u0001])","excludeEnd":true,"excludeBegin":true,"className":"string"}]}],"case_insensitive":true}`)
}