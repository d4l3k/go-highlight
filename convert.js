const hljs = require('highlight.js');
const fs = require('fs');
const path = require('path');

const dir = './node_modules/highlight.js/lib/languages/';

function cleanRegex(obj, parents) {
  if (!parents) {
    parents = [];
  }
  parents = parents.concat([obj]);
  for (let prop in obj) {
    const val = obj[prop];
    if (val instanceof RegExp) {
      obj[prop] = val.source;
    } else if (val === Object(val) && parents.indexOf(val) === -1) {
      obj[prop] = cleanRegex(val, parents);
    }
  }
  return obj;
}

fs.readdir(dir, (err, files) => {
  files.forEach(file => {
    const p = dir + file;
    const lang = path.basename(file, ".js");
    fs.readFile(p, (err, data) => {
      if (err) throw err;
      const def = cleanRegex(eval(data.toString())(hljs));
      console.log("Language:", lang);
      try {
        const json = JSON.stringify(def).replace(/`/g, "`+\"`\"+`");
        fs.writeFile("languages/"+lang+".go",
`package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("${lang}", \`${json}\`)
}`, (err) => {
              if (err) throw err;
            });
      } catch (e) {
        console.log(" - failed: ", e);
      };
    });

  });
});
