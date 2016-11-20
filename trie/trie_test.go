package trie

import "testing"

func TestTrieEmpty(t *testing.T) {
	var trie Trie
	if out := trie.MatchPrefix([]byte("asdfasdf")); out != "" {
		t.Errorf("empty trie should return an empty label; not %q", out)
	}

	trie.Add([]byte("asdf"), "test")
	if out := trie.MatchPrefix(nil); out != "" {
		t.Errorf("trie.MatchPrefix(nil) should return an empty label; not %q", out)
	}
}

func TestTrie(t *testing.T) {
	var trie Trie
	trie.Add([]byte("asdf"), "test")
	trie.Add([]byte("fdsa"), "test2")

	if out := trie.MatchPrefix([]byte("asdfasdf")); out != "test" {
		t.Errorf("MatchPrefix(\"asdfasdf\") = %q; not %q", out, "test")
	}

	if out := trie.MatchPrefix([]byte("fdsa")); out != "test2" {
		t.Errorf("MatchPrefix(\"asdfasdf\") = %q; not %q", out, "test2")
	}

}
