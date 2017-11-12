package gotrie

import (
	"fmt"
	"strings"
)

type SimpleTrie struct {
	key  string
	data interface{}
	sons []SimpleTrie
}

func (n SimpleTrie) Insert(key string, data interface{}) (Tree, bool) { return n.insert(key, data) }

func (n SimpleTrie) Update(key string, data interface{}) (Tree, bool) { panic("not implemented") }

func (n SimpleTrie) Delete(key string) (interface{}, Tree, bool) { panic("not implemented") }

func (n SimpleTrie) Query(key string) (interface{}, bool) { panic("not implemented") }

func (n SimpleTrie) isUninit() bool { return n.key == "" && n.data == nil && n.sons == nil }
func (n SimpleTrie) isLeaf() bool   { return n.key != "" && n.sons == nil }

func (n SimpleTrie) insert(key string, data interface{}) (SimpleTrie, bool) {
	if n.isUninit() {
		return SimpleTrie{key: key, data: data}, true
	}
	prefix, k1, k2 := lcs(n.key, key)
	if prefix == "" { // No common prefix
		return SimpleTrie{sons: []SimpleTrie{n, {key: key, data: data}}}, true
	}
	if k1 == "" {
		if k2 == "" { // Same object
			return n, false
		}
		for i := 0; i < len(n.sons); i++ {
			if prefix, _, _ := lcs(n.sons[i].key, k2); prefix != "" {
				ns, ok := n.sons[i].insert(k2, data)
				if ok {
					n.sons[i] = ns
				}
				return n, ok
			}
		}
		return SimpleTrie{
			key: n.key, data: n.data,
			sons: append(n.sons, SimpleTrie{key: k2, data: data}),
		}, true
	}
	n1 := SimpleTrie{key: k1, data: n.data, sons: n.sons}
	n2 := SimpleTrie{key: k2, data: data}
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
