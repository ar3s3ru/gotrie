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
		nt, _ := tt.Insert(words[i], nil)
		tt = nt.(*gotrie.CommonPrefixTree)
	}
}

func BenchmarkCommonPrefixTree_InsertAll(b *testing.B) {
	b.Log("Words length is ", len(words))
	b.Run("", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var tt *gotrie.CommonPrefixTree
			for _, w := range words {
				nt, _ := tt.Insert(w, nil)
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
