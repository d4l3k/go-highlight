package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("moonscript", `{"aliases":[0],"keywords":{"keyword":"if then not for in while do return else elseif break continue switch and or unless when class extends super local import export from using","literal":"true false nil","built_in":"_G _VERSION assert collectgarbage dofile error getfenv getmetatable ipairs load loadfile loadstring module next pairs pcall print rawequal rawget rawset require select setfenv setmetatable tonumber tostring type unpack xpcall coroutine debug io math os package string table"},"illegal":"/\\/\\*/","contains":[0,1,2,3,4,5,6,7,8,9]}`)
}