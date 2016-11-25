package registry

import (
	"encoding/json"
	"regexp"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

var languagesMu = struct {
	sync.RWMutex

	defs  map[string]*unparsedLanguage
	cache map[string]Language
	names []string
}{
	defs:  map[string]*unparsedLanguage{},
	cache: map[string]Language{},
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

// Language represents a language definition.
type Language struct {
	CaseInsensitive bool        `json:"case_insensitive"`
	Aliases         []string    `json:"aliases"`
	Keywords        *Keywords   `json:"keywords"`
	Illegal         string      `json:"illegal"`
	Contains        []*Contains `json:"contains"`
}

type keywordsJSON struct {
	Keyword string `json:"keyword"`
	Literal string `json:"literal"`
	BuiltIn string `json:"built_in"`
}

func parseContainsRaw(parent *Contains, cs []json.RawMessage) ([]*Contains, error) {
	final := make([]*Contains, len(cs))
	for i, cm := range cs {
		var c Contains
		if err := json.Unmarshal(cm, &c); err != nil {
			var s string
			if err2 := json.Unmarshal(cm, &s); err2 != nil {
				return nil, errors.Wrap(err, err.Error())
			}
			if s == "self" {
				final[i] = parent
				continue
			} else {
				return nil, err
			}
		}
		final[i] = &c
	}
	return final, nil
}

type containsJSON struct {
	ClassName string            `json:"className"`
	Contains  []json.RawMessage `json:"contains"`
	Variants  []json.RawMessage `json:"variants"`

	Begin          string    `json:"begin"`
	BeginLookahead string    `json:"beginLookahead"`
	End            string    `json:"end"`
	BeginKeywords  string    `json:"beginKeywords"`
	Keywords       *Keywords `json:"keywords"`
	ExcludeEnd     bool      `json:"excludeEnd"`
	Relevance      float64   `json:"relevance"`
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
	return strings.Split(words, " ")
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
	ClassName string
	Contains  []*Contains
	Variants  []*Contains

	Begin          *regexp.Regexp
	BeginLookahead *regexp.Regexp
	End            *regexp.Regexp
	BeginKeywords  []string
	Keywords       *Keywords
	ExcludeEnd     bool
	Relevance      float64
}

// UnmarshalJSON unmarshals.
func (c *Contains) UnmarshalJSON(b []byte) error {
	var con containsJSON
	err := json.Unmarshal(b, &con)
	if err != nil {
		return errors.Wrapf(err, "Contains UnmarshalJSON(%s)", b)
	}

	c.ClassName = con.ClassName

	c.Contains, err = parseContainsRaw(c, con.Contains)
	if err != nil {
		return err
	}
	c.Variants, err = parseContainsRaw(c, con.Variants)
	if err != nil {
		return err
	}

	if len(con.Begin) > 0 {
		c.Begin, err = regexp.Compile(con.Begin)
		if err != nil {
			return err
		}
	}

	if len(con.BeginLookahead) > 0 {
		c.BeginLookahead, err = regexp.Compile(con.BeginLookahead)
		if err != nil {
			return err
		}
	}

	if len(con.End) > 0 {
		// Regex needs to be in multi line mode and match starting at the
		// beginning of the string.
		c.End, err = regexp.Compile(`(?m:` + con.End + `)`)
		if err != nil {
			return err
		}
	}

	c.BeginKeywords = parseWords(con.BeginKeywords)
	c.Keywords = con.Keywords
	c.ExcludeEnd = con.ExcludeEnd
	c.Relevance = con.Relevance

	return nil
}

func parseLang(def string) (Language, error) {
	var lang Language
	if err := json.Unmarshal([]byte(def), &lang); err != nil {
		return Language{}, err
	}
	return lang, nil
}

// ErrLanguageNotFound is returned when a requested language is not present in
// the registry.
var ErrLanguageNotFound = errors.New("can't find language in registry")

// Lookup finds and returns the parsed Language that has been saved in the
// registry.
func Lookup(name string) (Language, error) {
	languagesMu.RLock()
	lang, ok := languagesMu.cache[name]
	if ok {
		languagesMu.RUnlock()
		return lang, nil
	}
	langDef, ok := languagesMu.defs[name]
	languagesMu.RUnlock()
	if !ok {
		return Language{}, ErrLanguageNotFound
	}

	lang, err := parseLang(langDef.body)
	if err != nil {
		return Language{}, errors.Wrapf(err, "failed to parse %s", name)
	}

	languagesMu.Lock()
	defer languagesMu.Unlock()

	languagesMu.cache[name] = lang
	return lang, nil
}
