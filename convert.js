const hljs = require('highlight.js');
const fs = require('fs');
const path = require('path');


const dir = './node_modules/highlight.js/lib/languages/';

function goArray(strings) {
  return "[]string{"+strings.map(a => JSON.stringify(a)).join(", ")+"}";
}

function cleanRegex(obj, parents, path) {
  if (!parents) {
    parents = [];
  }
  if (!path) {
    path = [];
  }
  parents = parents.concat([obj]);
  for (let prop in obj) {
    const val = obj[prop];
    if (val instanceof RegExp) {
      obj[prop] = val.source;
    } else if (val === Object(val)) {
      var idx = parents.indexOf(val);
      if (idx === -1) {
        obj[prop] = cleanRegex(val, parents, path.concat(prop.toString()));
      } else {
        if (val instanceof Array) {
          obj[prop] = [{Ref: path.slice(0, idx), IsArray: true}];
        } else {
          obj[prop] = {Ref: path.slice(0, idx)};
        }
      }
    }

    // Hack to fix edge case where there should be no end.
    if (prop == "end" && val === false) {
      obj[prop] = ".^";
    }

    if (prop == "subLanguage") {
      if (typeof val === "string") {
        obj[prop] = [val];
      } else if (val.length === 0) {
        obj[prop] = ["all"];
      }
    }

    // Need to call JSON.parse to correctly parse unicode escape sequences in
    // the regexps.
    let regexp = obj[prop];
    if (typeof regexp === "string") {
      // lisp strangeness.
      regexp = regexp.replace("[^]*", ".*");
      while (true) {
        const matches = regexp.match(/\\u[0-9A-Fa-f]{4}/)
          if (!matches) {
            break;
          }
        const match = matches[0];
        regexp = regexp.replace(match, JSON.parse('"'+match+'"'));
      }
      obj[prop] = regexp;
    }

    if (obj instanceof Array && obj[prop] === "self") {
      obj[prop] = {Ref: path.slice(0,path.length-1)};
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
