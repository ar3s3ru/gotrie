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
		return nil, false
	}
	return t.root.get(key)
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

func (n *commonPrefixNode) get(key string) (interface{}, bool) {
	_, k1, k2 := lcs(n.key, key)
	switch {
	case k1 == "" && k2 == "": // Found if n is a word
		return n.data, n.word
	case k1 == "" && k2 != "": // Search through sons
		for _, s := range n.sons {
			if prefix, _, _ := lcs(s.key, k2); prefix != "" {
				return s.get(k2)
			}
		}
		fallthrough
	default:
		return nil, false
	}
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
