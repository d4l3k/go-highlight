package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("crystal", `{"aliases":[0],"lexemes":"[a-zA-Z_]\\w*[!?=]?","keywords":{"keyword":"abstract alias as asm begin break case class def do else elsif end ensure enum extend for fun if ifdef include instance_sizeof is_a? lib macro module next of out pointerof private protected rescue responds_to? return require self sizeof struct super then type typeof union unless until when while with yield __DIR__ __FILE__ __LINE__","literal":"false nil true"},"contains":[0,1,2,3,4,5,6,7,8,9,10,11,12]}`)
}