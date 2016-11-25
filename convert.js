const hljs = require('highlight.js');
const fs = require('fs');
const path = require('path');

const dir = './node_modules/highlight.js/lib/languages/';

function goArray(strings) {
  return "[]string{"+strings.map(a => JSON.stringify(a)).join(", ")+"}";
}

function cleanRegex(obj, parents) {
  if (!parents) {
    parents = [];
  }
  parents = parents.concat([obj]);
  for (let prop in obj) {
    const val = obj[prop];
    if (val instanceof RegExp) {
      // Need to call JSON.parse to correctly parse unicode escape sequences in
      // the regexps.
      let regexp = val.source;
      while (true) {
        const matches = regexp.match(/\\u[0-9A-Fa-f]{4}/)
        if (!matches) {
          break;
        }
        const match = matches[0];
        regexp = regexp.replace(match, JSON.parse('"'+match+'"'));
      }
      obj[prop] = regexp;
    } else if (val === Object(val) && parents.indexOf(val) === -1) {
      obj[prop] = cleanRegex(val, parents);
    }

    // Hack to fix edge case where there should be no end.
    if (prop == "end" && val === false) {
      obj[prop] = ".^";
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
      const aliases = [lang].concat(def.aliases);
      console.log("Language:", lang);
      try {
        const json = JSON.stringify(def).replace(/`/g, "`+\"`\"+`");
        fs.writeFile("languages/"+lang+".go",
`package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register(${goArray(aliases)}, \`${json}\`)
}`, (err) => {
              if (err) throw err;
            });
      } catch (e) {
        console.log(" - failed: ", e);
      };
    });

  });
});
