package gotrie

import (
	"fmt"
	"strings"
)

type CommonPrefixTree struct {
	root *commonPrefixNode
	size uint
}

func (t *CommonPrefixTree) isUninit() bool   { return t.root == nil }
func (t *CommonPrefixTree) isLeaf() bool     { return !t.isUninit() && t.root.isLeaf() }
func (t *CommonPrefixTree) isWildcard() bool { return !t.isUninit() && t.root.isWildcard() }

// Insert tries to insert the key-value pair inside the tree.
// If ok, returns a new Tree with the key-value pair inside of it.
// If not ok, it means that the key is already in the tree.
func (t *CommonPrefixTree) Insert(key string, data interface{}) (Tree, bool) {
	// N.B. Insert is always called on the root of the tree
	if t == nil {
		t = &CommonPrefixTree{}
	}
	if t.isUninit() {
		t.root, t.size = &commonPrefixNode{key: key, data: data, word: true}, 1
		return t, true
	}
	n, ok := t.root.insert(key, data)
	if ok {
		t.size++
		t.root = n
	}
	return t, ok
}

func (t *CommonPrefixTree) Update(key string, data interface{}) (Tree, bool) { panic("not implemented") }
func (t *CommonPrefixTree) Delete(key string) (interface{}, Tree, bool)      { panic("not implemented") }

// Keys return the list of words contained inside the tree.
func (t *CommonPrefixTree) Keys() (keys []string) {
	if !t.isUninit() {
		keys = t.root.keys("")
	}
	return
}

// Get returns the meta data of associated key, if any.
// The bool flag indicates if the key has been found.
func (t *CommonPrefixTree) Get(key string) (interface{}, bool) {
	if t.isUninit() {
		goto fail
	}
	if !t.isWildcard() {
		i, ok, word := t.root.get(key)
		return i, ok && word
	}
	for _, s := range t.root.sons {
		if i, ok, word := s.get(key); ok {
			return i, ok && word
		}
	}
fail:
	return nil, false
}

func (t *CommonPrefixTree) String() string {
	if t.isUninit() {
		return "<uninit>"
	}
	return t.root.String()
}

type commonPrefixNode struct {
	key  string
	data interface{}
	word bool
	sons []*commonPrefixNode
}

func (n *commonPrefixNode) isLeaf() bool     { return n.sons == nil }
func (n *commonPrefixNode) isWildcard() bool { return n.key == "" && !n.isLeaf() }

func (n *commonPrefixNode) insert(key string, data interface{}) (*commonPrefixNode, bool) {
	prefix, k1, k2 := lcs(n.key, key)
	switch {
	case k1 == "" && k2 == "": // Same object
		n.word = true
		return n, false
	case k1 == "" && k2 != "": // Insert between sons
		for i, s := range n.sons {
			if prefix, _, _ := lcs(s.key, k2); prefix != "" {
				nn, ok := s.insert(k2, data)
				if ok {
					n.sons[i] = nn
				}
				return n, ok
			}
		}
		n.sons = append(n.sons, &commonPrefixNode{key: k2, data: data, word: true})
		return n, true
	default: // Create a new wildcard
		n.key = k1
		wc := &commonPrefixNode{key: prefix, sons: []*commonPrefixNode{n}, word: k2 == ""}
		if k2 != "" {
			wc.sons = append(wc.sons, &commonPrefixNode{key: k2, data: data, word: true})
		}
		return wc, true
	}
}

func (n *commonPrefixNode) Update(key string, data interface{}) (Tree, bool) { panic("not implemented") }

func (n *commonPrefixNode) Delete(key string) (interface{}, Tree, bool) { panic("not implemented") }

func (n *commonPrefixNode) get(key string) (interface{}, bool, bool) {
	prefix, k1, k2 := lcs(n.key, key)
	if prefix == "" {
		goto fail
	}
	if k2 == "" && k1 == "" {
		return n.data, true, n.word
	}
	if k1 != "" {
		goto fail
	}
	for _, s := range n.sons {
		if i, ok, word := s.get(k2); ok {
			return i, ok, word
		}
	}
fail:
	return nil, false, false
}

func (n *commonPrefixNode) keys(prefix string) (res []string) {
	prefix = fmt.Sprintf("%s%s", prefix, n.key)
	if n.word {
		res = append(res, prefix)
	}
	for _, s := range n.sons {
		if sres := s.keys(prefix); sres != nil {
			res = append(res, sres...)
		}
	}
	return
}

// String returns a string'd representation of the tree.
func (n *commonPrefixNode) String() string {
	if n.isLeaf() {
		return fmt.Sprintf("\"%s\" -> data: %v", n.key, n.data)
	}
	var ss []string
	for _, s := range n.sons {
		ss = append(ss, s.String())
	}
	return fmt.Sprintf("\"%s\": [%s]", n.key, strings.Join(ss, ","))
}
