package gotrie_test

import (
	"testing"

	"github.com/derekparker/trie"
	"github.com/stretchr/testify/assert"
)

func BenchmarkDerekParkerTrie_Insert(b *testing.B) {
	tt := trie.New()
	for n := 0; n < b.N; n++ {
		i := n % len(words)
		node := tt.Add(words[i], nil)
		assert.NotNil(b, node, "insertion failed [%s]: node: %v", words[i], node)
	}
}

func BenchmarkDerekParkerTrie_InsertAll(b *testing.B) {
	b.Log("Words length is ", len(words))
	b.Run("", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tt := trie.New()
			for _, w := range words {
				node := tt.Add(w, nil)
				assert.NotNil(b, node, "insertion failed [%s]: node: %v", w, node)
			}
		}
	})
}

func BenchmarkDerekParkerTrie_Get(b *testing.B) {
	tt := trie.New()
	for _, w := range words {
		node := tt.Add(w, nil)
		assert.NotNil(b, node, "insertion failed [%s]: node: %v", w, node)
	}
	b.Run("Query terms", func(b *testing.B) {
		for _, w := range words {
			_, ok := tt.Find(w)
			assert.True(b, ok, "not found: %s", w)
		}
	})
}
