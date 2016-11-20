package registry

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

var languages = map[string]string{}
var lookupCache = map[string]Language{}

// Register registers a specific language with the highlighter.
func Register(name, json string) {
	languages[name] = json
}

// Language ...
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

type containsJSON struct {
	ClassName string      `json:"className"`
	Contains  []*Contains `json:"contains"`
	Variants  []*Contains `json:"variants"`

	Begin         string    `json:"begin"`
	End           string    `json:"end"`
	BeginKeywords string    `json:"beginKeywords"`
	Keywords      *Keywords `json:"keywords"`
	ExcludeEnd    bool      `json:"excludeEnd"`
	Relevance     float64   `json:"relevance"`
}

// Keywords ...
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
		return err
	}

	k.Keyword = parseWords(kw.Keyword)
	k.Literal = parseWords(kw.Literal)
	k.BuiltIn = parseWords(kw.BuiltIn)

	return nil
}

// Contains ...
type Contains struct {
	ClassName string
	Contains  []*Contains
	Variants  []*Contains

	Begin         *regexp.Regexp
	End           *regexp.Regexp
	BeginKeywords []string
	Keywords      *Keywords
	ExcludeEnd    bool
	Relevance     float64
}

// UnmarshalJSON unmarshals.
func (c *Contains) UnmarshalJSON(b []byte) error {
	var con containsJSON
	err := json.Unmarshal(b, &con)
	if err != nil {
		return err
	}

	c.ClassName = con.ClassName
	c.Contains = con.Contains
	c.Variants = con.Variants

	if len(con.Begin) > 0 {
		c.Begin, err = regexp.Compile("^" + con.Begin)
		if err != nil {
			return err
		}
	}

	if len(con.End) > 0 {
		// Regex needs to be in multi line mode and match starting at the
		// beginning of the string.
		c.End, err = regexp.Compile(`^(?m:` + con.End + `)`)
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

// Lookup ...
func Lookup(name string) (Language, error) {
	if lang, ok := lookupCache[name]; ok {
		return lang, nil
	}
	lang, ok := languages[name]
	if !ok {
		return Language{}, errors.New("can't find language")
	}
	return parseLang(lang)
}
