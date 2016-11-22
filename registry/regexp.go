package registry

import (
	"bytes"
	"fmt"
	"regexp"
)

// Finder searches a byte slice and returns the first index.
type Finder interface {
	// FindIndex operates in the same way as (*regexp.Regexp).FindIndex.
	FindIndex(body []byte) []int
}

// cleanRegexp returns whether the provided string uses Regex features or can be
// matched using a simple string match.
func cleanRegexp(a string) (string, bool) {
	metaChars := []byte("^[].${*(+)|?{}")

	var final []byte

	escaped := false
	for _, l := range a {
		if !escaped && l == '\\' {
			escaped = true
			continue
		}

		contains := bytes.ContainsRune(metaChars, l)
		if contains && !escaped || !contains && escaped {
			return "", true
		}
		final = append(final, byte(l))
		escaped = false
	}
	return string(final), false
}

// StringFinder wraps a byte array and allows regex like FindIndex calls.
type StringFinder []byte

// FindIndex implements Finder.
func (s StringFinder) FindIndex(body []byte) []int {
	fmt.Println("StringFinder")
	i := bytes.Index(body, s)
	if i == -1 {
		return nil
	}
	return []int{i, i + len(s)}
}

func finderCompile(pattern string, mode string) (Finder, error) {
	clean, isRegex := cleanRegexp(pattern)
	if isRegex {
		if len(mode) > 0 {
			pattern = fmt.Sprintf("(?%s:%s)", mode, pattern)
		}
		return regexp.Compile(pattern)
	}

	finder := StringFinder(clean)

	return &finder, nil
}
