package gotrie

import (
	"fmt"
	"strings"
)

type SimpleTrie struct {
	key  string
	data interface{}
	sons []SimpleTrie
	word bool
}

func (n SimpleTrie) Insert(key string, data interface{}) (Tree, bool) {
	if n.isUninit() {
		return SimpleTrie{key: key, data: data, word: true}, true
	}
	if n.isWildcard() {
		for i, s := range n.sons {
			p, k1, k2 := lcs(s.key, key)
			if p != "" { // At least one common prefix on direct sons
				ns, ok := s.blindInsert(p, k1, k2, data) // Blind insert on common son
				n.sons[i] = ns
				return n, ok
			}
		}
		return SimpleTrie{sons: append(n.sons, SimpleTrie{key: key, data: data, word: true})}, true
	}
	prefix, k1, k2 := lcs(n.key, key)
	return n.blindInsert(prefix, k1, k2, data)
}

func (n SimpleTrie) Update(key string, data interface{}) (Tree, bool) { panic("not implemented") }

func (n SimpleTrie) Delete(key string) (interface{}, Tree, bool) { panic("not implemented") }

func (n SimpleTrie) Query(key string) (interface{}, bool) { panic("not implemented") }

func (n SimpleTrie) isUninit() bool   { return n.key == "" && n.data == nil && n.sons == nil }
func (n SimpleTrie) isLeaf() bool     { return n.key != "" && n.sons == nil }
func (n SimpleTrie) isWildcard() bool { return n.key == "" && n.sons != nil }

func (n SimpleTrie) blindInsert(prefix, k1, k2 string, data interface{}) (SimpleTrie, bool) {
	if k1 == "" {
		if k2 == "" { // Same object
			return n, false
		}
		for i, s := range n.sons {
			if prefix, k1, k2 := lcs(s.key, k2); prefix != "" {
				ns, ok := s.blindInsert(prefix, k1, k2, data)
				if ok {
					n.sons[i] = ns
				}
				return n, ok
			}
		}
		return SimpleTrie{
			key: n.key, data: n.data, word: n.word,
			sons: append(n.sons, SimpleTrie{key: k2, data: data, word: true}),
		}, true
	}
	n1 := SimpleTrie{key: k1, data: n.data, word: n.word, sons: n.sons}
	if k2 == "" {
		return SimpleTrie{key: prefix, sons: []SimpleTrie{n1}, word: true}, true
	}
	n2 := SimpleTrie{key: k2, data: data, word: true}
	return SimpleTrie{key: prefix, sons: []SimpleTrie{n1, n2}}, true
}

func (n SimpleTrie) String() string {
	if n.isLeaf() {
		return fmt.Sprintf("\"%s\" -> data: %v", n.key, n.data)
	}
	var ss []string
	for _, s := range n.sons {
		ss = append(ss, s.String())
	}
	return fmt.Sprintf("\"%s\": [%s]", n.key, strings.Join(ss, ","))
}

func (n SimpleTrie) Words() []string { return n.words("") }
func (n SimpleTrie) words(prefix string) (res []string) {
	prefix = fmt.Sprintf("%s%s", prefix, n.key)
	if n.word {
		res = append(res, prefix)
	}
	for _, s := range n.sons {
		if sres := s.words(prefix); sres != nil {
			res = append(res, sres...)
		}
	}
	return
}
