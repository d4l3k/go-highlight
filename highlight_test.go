package highlight

import (
	"log"
	"testing"
)

func TestHighlight(t *testing.T) {
	resp, err := Highlight("go", "package main")
	if err != nil {
		t.Error(err)
	}
	log.Println(resp, err)
}
