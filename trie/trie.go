package trie

// Trie represents a set of labels.
type Trie struct {
	label    string
	children map[byte]*Trie
}

// Add inserts a word with a label into the trie.
func (t *Trie) Add(a []byte, label string) {
	if len(a) == 0 {
		t.label = label
		return
	}

	if t.children == nil {
		t.children = map[byte]*Trie{}
	}

	l := a[0]
	c, ok := t.children[l]
	if !ok {
		c = &Trie{}
		t.children[l] = c
	}
	c.Add(a[1:], label)
}

// MatchPrefix scans through the view until it finds a matching label.
func (t *Trie) MatchPrefix(view []byte) string {
	for t != nil {
		if len(t.label) != 0 {
			return t.label
		}

		if len(view) == 0 {
			break
		}

		t = t.children[view[0]]
		view = view[1:]
	}

	return ""
}
