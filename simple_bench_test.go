package gotrie_test

import (
	"testing"

	"github.com/ar3s3ru/gotrie"
)

func BenchmarkSimpleTrie_Insert(b *testing.B) {
	var tt gotrie.SimpleTrie
	for n := 0; n < b.N; n++ {
		i := n % len(words)
		nt, _ := tt.Insert(words[i], nil)
		tt = nt.(gotrie.SimpleTrie)
	}
}

func BenchmarkSimpleTrie_InsertAll(b *testing.B) {
	b.Log("Words length is ", len(words))
	for n := 0; n < b.N; n++ {
		var tt gotrie.SimpleTrie
		for _, w := range words {
			nt, _ := tt.Insert(w, nil)
			tt = nt.(gotrie.SimpleTrie)
		}
	}
}
