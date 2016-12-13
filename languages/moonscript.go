package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"moonscript", "moon"}, `{"aliases":["moon"],"keywords":{"keyword":"if then not for in while do return else elseif break continue switch and or unless when class extends super local import export from using","literal":"true false nil","built_in":"_G _VERSION assert collectgarbage dofile error getfenv getmetatable ipairs load loadfile loadstring module next pairs pcall print rawequal rawget rawset require select setfenv setmetatable tonumber tostring type unpack xpcall coroutine debug io math os package string table"},"illegal":"\\/\\*","contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"className":"string","variants":[{"begin":"'","end":"'","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"if then not for in while do return else elseif break continue switch and or unless when class extends super local import export from using","literal":"true false nil","built_in":"_G _VERSION assert collectgarbage dofile error getfenv getmetatable ipairs load loadfile loadstring module next pairs pcall print rawequal rawget rawset require select setfenv setmetatable tonumber tostring type unpack xpcall coroutine debug io math os package string table"},"contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","1"]},{"className":"built_in","begin":"@__[a-zA-Z]\\w*"},{"begin":"@[a-zA-Z]\\w*"},{"begin":"[a-zA-Z]\\w*\\\\[a-zA-Z]\\w*"}]}]}]},{"className":"built_in","begin":"@__[a-zA-Z]\\w*"},{"begin":"@[a-zA-Z]\\w*"},{"begin":"[a-zA-Z]\\w*\\\\[a-zA-Z]\\w*"},{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"function","begin":"^\\s*[A-Za-z$_][0-9A-Za-z$_]*\\s*=\\s*(\\(.*\\))?\\s*\\B[-=]>","end":"[-=]>","returnBegin":true,"contains":[{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0},{"className":"params","begin":"\\([^\\(]","returnBegin":true,"contains":[{"begin":"\\(","end":"\\)","keywords":{"keyword":"if then not for in while do return else elseif break continue switch and or unless when class extends super local import export from using","literal":"true false nil","built_in":"_G _VERSION assert collectgarbage dofile error getfenv getmetatable ipairs load loadfile loadstring module next pairs pcall print rawequal rawget rawset require select setfenv setmetatable tonumber tostring type unpack xpcall coroutine debug io math os package string table"},"contains":[{"Ref":["contains","6","contains","1","contains","0"]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"className":"string","variants":[{"begin":"'","end":"'","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"if then not for in while do return else elseif break continue switch and or unless when class extends super local import export from using","literal":"true false nil","built_in":"_G _VERSION assert collectgarbage dofile error getfenv getmetatable ipairs load loadfile loadstring module next pairs pcall print rawequal rawget rawset require select setfenv setmetatable tonumber tostring type unpack xpcall coroutine debug io math os package string table"},"contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","1"]},{"className":"built_in","begin":"@__[a-zA-Z]\\w*"},{"begin":"@[a-zA-Z]\\w*"},{"begin":"[a-zA-Z]\\w*\\\\[a-zA-Z]\\w*"}]}]}]},{"className":"built_in","begin":"@__[a-zA-Z]\\w*"},{"begin":"@[a-zA-Z]\\w*"},{"begin":"[a-zA-Z]\\w*\\\\[a-zA-Z]\\w*"}]}]}]},{"begin":"[\\(,:=]\\s*","relevance":0,"contains":[{"className":"function","begin":"(\\(.*\\))?\\s*\\B[-=]>","end":"[-=]>","returnBegin":true,"contains":[{"className":"params","begin":"\\([^\\(]","returnBegin":true,"contains":[{"begin":"\\(","end":"\\)","keywords":{"keyword":"if then not for in while do return else elseif break continue switch and or unless when class extends super local import export from using","literal":"true false nil","built_in":"_G _VERSION assert collectgarbage dofile error getfenv getmetatable ipairs load loadfile loadstring module next pairs pcall print rawequal rawget rawset require select setfenv setmetatable tonumber tostring type unpack xpcall coroutine debug io math os package string table"},"contains":[{"Ref":["contains","6","contains","1","contains","0"]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"className":"string","variants":[{"begin":"'","end":"'","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"if then not for in while do return else elseif break continue switch and or unless when class extends super local import export from using","literal":"true false nil","built_in":"_G _VERSION assert collectgarbage dofile error getfenv getmetatable ipairs load loadfile loadstring module next pairs pcall print rawequal rawget rawset require select setfenv setmetatable tonumber tostring type unpack xpcall coroutine debug io math os package string table"},"contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","1"]},{"className":"built_in","begin":"@__[a-zA-Z]\\w*"},{"begin":"@[a-zA-Z]\\w*"},{"begin":"[a-zA-Z]\\w*\\\\[a-zA-Z]\\w*"}]}]}]},{"className":"built_in","begin":"@__[a-zA-Z]\\w*"},{"begin":"@[a-zA-Z]\\w*"},{"begin":"[a-zA-Z]\\w*\\\\[a-zA-Z]\\w*"}]}]}]}]},{"className":"class","beginKeywords":"class","end":"$","illegal":"[:=\"\\[\\]]","contains":[{"beginKeywords":"extends","endsWithParent":true,"illegal":"[:=\"\\[\\]]","contains":[{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0}]},{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0}]},{"className":"name","begin":"[A-Za-z$_][0-9A-Za-z$_]*:","end":":","returnBegin":true,"returnEnd":true,"relevance":0}]}`)
}