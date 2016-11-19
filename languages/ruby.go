package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("ruby", `{"aliases":[0,1,2,3,4],"keywords":{"keyword":"and then defined module in return redo if BEGIN retry end for self when next until do begin unless END rescue else break undef not super class case require yield alias while ensure elsif or include attr_reader attr_writer attr_accessor","literal":"true false nil"},"illegal":"/\\/\\*/","contains":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18]}`)
}