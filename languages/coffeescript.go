package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"coffeescript", "coffee", "cson", "iced"}, `{"aliases":["coffee","cson","iced"],"keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"illegal":"\\/\\*","contains":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"className":"string","variants":[{"begin":"'''","end":"'''","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"'","end":"'","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"\"\"\"","end":"\"\"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","2"]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]},{"begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","2"]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]}]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"},{"className":"comment","begin":"###","end":"###","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"function","begin":"^\\s*[A-Za-z$_][0-9A-Za-z$_]*\\s*=\\s*(\\(.*\\))?\\s*\\B[-=]>","end":"[-=]>","returnBegin":true,"contains":[{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0},{"className":"params","begin":"\\([^\\(]","returnBegin":true,"contains":[{"begin":"\\(","end":"\\)","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"Ref":["contains","8","contains","1","contains","0"]},{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"className":"string","variants":[{"begin":"'''","end":"'''","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"'","end":"'","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"\"\"\"","end":"\"\"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","2"]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]},{"begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","2"]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]}]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]}]},{"begin":"[:\\(,=]\\s*","relevance":0,"contains":[{"className":"function","begin":"(\\(.*\\))?\\s*\\B[-=]>","end":"[-=]>","returnBegin":true,"contains":[{"className":"params","begin":"\\([^\\(]","returnBegin":true,"contains":[{"begin":"\\(","end":"\\)","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"Ref":["contains","8","contains","1","contains","0"]},{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"className":"string","variants":[{"begin":"'''","end":"'''","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"'","end":"'","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"\"\"\"","end":"\"\"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","2"]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]},{"begin":"\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"#\\{","end":"}","keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"contains":[{"className":"number","begin":"\\b(0b[01]+)","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0,"starts":{"end":"(\\s*/)?","relevance":0}},{"Ref":["contains","2"]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]}]},{"className":"regexp","variants":[{"begin":"///","end":"///","contains":[{"Ref":["contains","2","variants","2","contains","1"]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"//[gim]*","relevance":0},{"begin":"\\/(?![ *])(\\\\\\/|.)*?\\/[gim]*(?=\\W|$)"}]},{"begin":"@[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"`+"`"+`","end":"`+"`"+`","excludeBegin":true,"excludeEnd":true,"subLanguage":"javascript"}]}]}]}]},{"className":"class","beginKeywords":"class","end":"$","illegal":"[:=\"\\[\\]]","contains":[{"beginKeywords":"extends","endsWithParent":true,"illegal":"[:=\"\\[\\]]","contains":[{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0}]},{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0}]},{"begin":"[A-Za-z$_][0-9A-Za-z$_]*:","end":":","returnBegin":true,"returnEnd":true,"relevance":0}]}`)
}