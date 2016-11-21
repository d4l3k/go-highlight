package registry

import (
	"reflect"
	"testing"
)

func TestRegistry(t *testing.T) {
	name := "test"
	json := `{"aliases": ["test"]}`
	Register(name, json)

	if jsonOut := languages[name]; jsonOut != json {
		t.Fatalf("Register(%q, %q); languages[%q] = %q; not %q", name, json, name, jsonOut, json)
	}

	lang, err := Lookup(name)
	if err != nil {
		t.Fatal(err)
	}
	expectedLang := Language{
		Aliases: []string{"test"},
	}
	if !reflect.DeepEqual(lang, expectedLang) {
		t.Fatalf("Lookup(%q) = %+v; not %+v", name, lang, expectedLang)
	}
}
