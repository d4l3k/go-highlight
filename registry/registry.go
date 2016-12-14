package registry

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"

	pcre "github.com/gijsbers/go-pcre"
	"github.com/pkg/errors"
)

var languagesMu = struct {
	sync.RWMutex

	defs  map[string]*unparsedLanguage
	cache map[string]Contains
	names []string
}{
	defs:  map[string]*unparsedLanguage{},
	cache: map[string]Contains{},
}

// Register registers a specific language with the highlighter.
func Register(names []string, json string) {
	lang := &unparsedLanguage{
		name:    names[0],
		aliases: names[1:],
		body:    json,
	}

	languagesMu.Lock()
	defer languagesMu.Unlock()

	languagesMu.names = append(languagesMu.names, lang.name)
	for _, name := range names {
		languagesMu.defs[name] = lang
	}
}

// Languages returns an slice of all the language names.
func Languages() []string {
	languagesMu.RLock()
	defer languagesMu.RUnlock()

	return languagesMu.names
}

type unparsedLanguage struct {
	name    string
	aliases []string
	body    string
}

type keywordsJSON struct {
	Keyword string `json:"keyword"`
	Literal string `json:"literal"`
	BuiltIn string `json:"built_in"`
}

// arrayToUpper capitalizes all elements in the array.
func arrayToUpper(parts []string) []string {
	for i, part := range parts {
		parts[i] = strings.ToUpper(part[0:1]) + part[1:]
	}
	return parts
}

type containsJSON struct {
	CaseInsensitive bool     `json:"case_insensitive"`
	Aliases         []string `json:"aliases"`
	Illegal         string   `json:"illegal"`

	ClassName string      `json:"className"`
	Contains  []*Contains `json:"contains"`
	Variants  []*Contains `json:"variants"`
	Starts    *Contains   `json:"starts"`

	Begin          string    `json:"begin"`
	BeginLookahead string    `json:"beginLookahead"`
	End            string    `json:"end"`
	BeginKeywords  string    `json:"beginKeywords"`
	Keywords       *Keywords `json:"keywords"`
	ExcludeEnd     bool      `json:"excludeEnd"`
	Relevance      float64   `json:"relevance"`

	// Ref and IsArray are used for resolving circular references within the
	// definitions.
	Ref     []string
	IsArray bool
}

// Keywords represents a set of keywords that should be matched and highlighted.
type Keywords struct {
	Keyword []string
	Literal []string
	BuiltIn []string
}

func parseWords(words string) []string {
	if len(words) == 0 {
		return nil
	}

	// Avoid empty strings.
	var final []string
	for _, s := range strings.Split(words, " ") {
		if len(s) == 0 {
			continue
		}

		final = append(final, s)
	}

	return final
}

// UnmarshalJSON unmarshals.
func (k *Keywords) UnmarshalJSON(b []byte) error {
	var kw keywordsJSON
	if err := json.Unmarshal(b, &kw); err != nil {
		// Unmarshalling failed. Try unmarshalling into string.
		if err := json.Unmarshal(b, &kw.Keyword); err != nil {
			return errors.Wrap(err, "Keywords UnmarshalJSON")
		}
	}

	k.Keyword = parseWords(kw.Keyword)
	k.Literal = parseWords(kw.Literal)
	k.BuiltIn = parseWords(kw.BuiltIn)

	return nil
}

// Contains represents a subsection that can match different parts of the code.
type Contains struct {
	CaseInsensitive bool
	Aliases         []string
	Illegal         string

	ClassName string
	Contains  []*Contains
	Variants  []*Contains
	Starts    *Contains

	Begin         *pcre.Regexp
	End           *pcre.Regexp
	BeginKeywords []string
	Keywords      *Keywords
	ExcludeEnd    bool
	Relevance     float64

	// Ref and IsArray are used for resolving circular references within the
	// definitions.
	Ref     []string
	IsArray bool
}

func compileRegex(regex string, flags int) (*pcre.Regexp, error) {
	if len(regex) == 0 {
		return nil, nil
	}

	r, err := pcre.Compile(regex, flags)
	if err != nil {
		return nil, err
	}
	if err := r.Study(0); err != nil {
		log.Printf("WARN: failed to JIT regex %q: %s", regex, err)
	}
	return &r, nil
}

// UnmarshalJSON unmarshals.
func (c *Contains) UnmarshalJSON(b []byte) error {
	con := containsJSON{
		Relevance: 1,
	}

	err := json.Unmarshal(b, &con)
	if err != nil {
		return errors.Wrapf(err, "Contains UnmarshalJSON(%s)", b)
	}

	c.CaseInsensitive = con.CaseInsensitive
	c.Aliases = con.Aliases
	c.Illegal = con.Illegal
	c.ClassName = con.ClassName
	c.Contains = con.Contains
	c.Variants = con.Variants
	c.Starts = con.Starts

	c.Begin, err = compileRegex(con.Begin, 0)
	if err != nil {
		return err
	}

	// Regex needs to be in multi line mode and match starting at the
	// beginning of the string.
	c.End, err = compileRegex(con.End, pcre.MULTILINE)
	if err != nil {
		return err
	}

	c.BeginKeywords = parseWords(con.BeginKeywords)
	c.Keywords = con.Keywords
	c.ExcludeEnd = con.ExcludeEnd
	c.Relevance = con.Relevance

	c.Ref = arrayToUpper(con.Ref)
	c.IsArray = con.IsArray

	return nil
}

func parseLang(def string) (Contains, error) {
	lang := &Contains{}
	if err := json.Unmarshal([]byte(def), &lang); err != nil {
		return Contains{}, err
	}

	visited := map[*Contains]struct{}{}
	if err := lang.resolveReferences(visited, lang); err != nil {
		return Contains{}, err
	}

	return *lang, nil
}

func (c *Contains) resolveReferences(visited map[*Contains]struct{}, n *Contains) error {
	// Detect circular references.
	if _, ok := visited[n]; ok {
		return nil
	}
	visited[n] = struct{}{}

	for _, nc := range []*[]*Contains{&n.Contains, &n.Variants} {
		if err := c.resolveReferencesArr(visited, nc); err != nil {
			return err
		}
	}
	return nil
}

func (c *Contains) resolveReferencesArr(visited map[*Contains]struct{}, n *[]*Contains) error {
	for i, inner := range *n {
		if inner.Ref == nil {
			if err := c.resolveReferences(visited, inner); err != nil {
				return err
			}
			continue
		}

		v, err := c.resolveReferencePath(inner.Ref)
		if err != nil {
			return err
		}
		if inner.IsArray {
			*n = v.([]*Contains)
		} else {
			(*n)[i] = v.(*Contains)
		}
	}

	return nil
}

func (c *Contains) resolveReferencePath(ref []string) (interface{}, error) {
	v := reflect.ValueOf(c)
	for len(ref) > 0 {
		r := ref[0]
		kind := v.Type().Kind()
		if kind == reflect.Slice || kind == reflect.Array {
			n, err := strconv.Atoi(r)
			if err != nil {
				return nil, err
			}
			v = v.Index(n)
		} else {
			v = v.Elem().FieldByName(r)
		}
		ref = ref[1:]
	}

	return v.Interface(), nil
}

// ErrLanguageNotFound is returned when a requested language is not present in
// the registry.
var ErrLanguageNotFound = errors.New("can't find language in registry")

// Lookup finds and returns the parsed Language that has been saved in the
// registry.
func Lookup(name string) (Contains, error) {
	languagesMu.RLock()
	lang, ok := languagesMu.cache[name]
	if ok {
		languagesMu.RUnlock()
		return lang, nil
	}
	langDef, ok := languagesMu.defs[name]
	languagesMu.RUnlock()
	if !ok {
		return Contains{}, ErrLanguageNotFound
	}

	lang, err := parseLang(langDef.body)
	if err != nil {
		return Contains{}, errors.Wrapf(err, "failed to parse %s", name)
	}

	languagesMu.Lock()
	defer languagesMu.Unlock()

	languagesMu.cache[name] = lang
	return lang, nil
}
