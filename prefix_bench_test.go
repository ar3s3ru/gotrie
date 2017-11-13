package gotrie_test

import (
	"testing"

	"github.com/ar3s3ru/gotrie"

	"github.com/stretchr/testify/assert"
)

func BenchmarkCommonPrefixTree_Insert(b *testing.B) {
	var tt *gotrie.CommonPrefixTree
	for n := 0; n < b.N; n++ {
		i := n % len(words)
		nt, ok := tt.Insert(words[i], nil)
		assert.True(b, ok, "insertion failed: %s", words[i])
		tt = nt.(*gotrie.CommonPrefixTree)
	}
}

func BenchmarkCommonPrefixTree_InsertAll(b *testing.B) {
	b.Log("Words length is ", len(words))
	b.Run("", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var tt *gotrie.CommonPrefixTree
			for _, w := range words {
				nt, ok := tt.Insert(w, nil)
				assert.True(b, ok, "insertion failed: %s", w)
				tt = nt.(*gotrie.CommonPrefixTree)
			}
		}
	})
}

func BenchmarkCommonPrefixTree_Get(b *testing.B) {
	var tt *gotrie.CommonPrefixTree
	for _, w := range words {
		nt, _ := tt.Insert(w, nil)
		tt = nt.(*gotrie.CommonPrefixTree)
	}
	b.Run("Query terms", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			i := n % len(words)
			_, ok := tt.Get(words[i])
			assert.True(b, ok, "query %s returned false", words[i])
		}
	})
}

// Geohash indexing

func BenchmarkCommonPrefixTree_Geohash(b *testing.B) {
	geohashes := []string{
		"gbsuv7zq", "gbsuv7zw", "gbsuv7zy",
		"gbsuv7zm", "gbsuv7zt", "gbsuv7zv",
		"gbsuv7zk", "gbsuv7zs", "gbsuv7zu",
		"gbszkkp3", "gbszkkp9", "gbszkkpc",
		"gbszkkp2", "gbszkkp8", "gbszkkpb",
		"gbszk7zr", "gbszk7zx", "gbszk7zz",
	}
	var t gotrie.Tree
	b.Run("Insert geohashes", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var nt gotrie.Tree = &gotrie.CommonPrefixTree{}
			for _, w := range geohashes {
				var ok bool
				nt, ok = nt.Insert(w, nil)
				assert.True(b, ok, "insertion failed: %s", w)
			}
			t = nt
		}
	})
	b.Logf("Generated tree is: %s\n", t)
	b.Run("Query geohashes", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, w := range geohashes {
				_, ok := t.Get(w)
				assert.True(b, ok, "get failed: %s", w)
			}
		}
	})
}
