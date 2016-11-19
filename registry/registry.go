package registry

import (
	"encoding/json"
	"errors"
)

var languages = map[string]string{}
var lookupCache = map[string]Language{}

// Register registers a specific language with the highlighter.
func Register(name, json string) {
	languages[name] = json
}

// Language ...
type Language struct {
	CaseInsensitive bool       `json:"case_insensitive"`
	Aliases         []string   `json:"aliases"`
	Keywords        Keywords   `json:"keywords"`
	Illegal         string     `json:"illegal"`
	Contains        []Contains `json:"contains"`
}

// Keywords ...
type Keywords struct {
	Keyword string `json:"keyword"`
	Literal string `json:"literal"`
	BuiltIn string `json:"built_in"`
}

// Contains ...
type Contains struct {
	ClassName string     `json:"className"`
	Contains  []Contains `json:"contains"`
	Variants  []Contains `json:"variants"`

	Begin         string  `json:"begin"`
	End           string  `json:"end"`
	BeginKeywords string  `json:"beginKeywords"`
	ExcludeEnd    bool    `json:"excludeEnd"`
	Relevance     float64 `json:"relevance"`
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
