package internal

type Trie struct {
	valid    bool
	children map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		valid:    false,
		children: make(map[rune]*Trie),
	}
}

func (t *Trie) Insert(s string) {
	for _, c := range s {
		if _, ok := t.children[c]; !ok {
			t.children[c] = NewTrie()
		}
		t = t.children[c]
	}
	t.valid = true
}

func (t *Trie) Match(prefix string) []string {
	for _, c := range prefix {
		if _, ok := t.children[c]; !ok {
			return []string{}
		}
		t = t.children[c]
	}
	return t.Enumerate()
}

func (t *Trie) Enumerate() []string {
	return t.dfs("")
}

func (t *Trie) dfs(prefix string) []string {
	if t == nil {
		return []string{}
	}

	res := []string{}

	if t.valid {
		res = append(res, prefix)
	}

	for c, child := range t.children {
		res = append(res, child.dfs(prefix+string(c))...)
	}

	return res
}

func (t *Trie) LongestMatch(s string) (string, bool) {
	for _, c := range s {
		if _, ok := t.children[c]; !ok {
			return "", false
		}
		t = t.children[c]
	}
	if len(t.children) != 1 {
		return "", false
	}
	prefix := ""
	for !t.valid && len(t.children) == 1 {
		c := Keys(t.children)[0]
		prefix += string(c)
		t = t.children[c]
	}
	return prefix, true
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
