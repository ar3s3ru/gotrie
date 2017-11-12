package gotrie

type SimpleTrie struct {
	root *simpleNode
	size uint
}

type simpleNode struct {
	val  rune
	word bool
	occs uint
	meta interface{}
	sons map[rune]*simpleNode
}

func (n *simpleNode) Insert(key string, meta interface{}) (Tree, bool) {
	k, key := rune(key[0]), key[1:]
	if n.sons == nil && k != n.val { // Is leaf
		if len(key) > 0 {

		}
		_ = &simpleNode{
			val: '*', occs: n.occs + 1, word: len(key) == 1,
			sons: map[rune]*simpleNode{n.val: n},
		}
	}
	panic("")
}

func (n *simpleNode) Remove(key string) bool                   { panic("") }
func (n *simpleNode) Update(key string, meta interface{}) bool { panic("") }
func (n *simpleNode) Keys() []string                           { panic("") }
