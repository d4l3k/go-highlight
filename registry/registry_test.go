package registry

import (
	"reflect"
	"testing"
)

func TestRegistry(t *testing.T) {
	name := "test"
	json := `{"aliases": ["test"]}`
	Register([]string{name}, json)

	if jsonOut := languagesMu.defs[name].body; jsonOut != json {
		t.Fatalf("Register(%q, %q); languages[%q] = %q; not %q", name, json, name, jsonOut, json)
	}

	lang, err := Lookup(name)
	if err != nil {
		t.Fatal(err)
	}
	expectedLang := Contains{
		Aliases:   []string{"test"},
		Relevance: 1,
	}
	if !reflect.DeepEqual(lang, expectedLang) {
		t.Fatalf("Lookup(%q) = %+v; not %+v", name, lang, expectedLang)
	}
}
