package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("ceylon", `{"keywords":{"keyword":"assembly module package import alias class interface object given value assign void function new of extends satisfies abstracts in out return break continue throw assert dynamic if else switch case for while try catch finally then let this outer super is exists nonempty shared abstract formal default actual variable late native deprecatedfinal sealed annotation suppressWarnings small","meta":"doc by license see throws tagged"},"illegal":"\\$[^01]|#[^0-9a-fA-F]","contains":[0,1,2,3,4,5,6]}`)
}