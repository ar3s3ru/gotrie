package gotrie_test

import (
	"sort"
	"testing"

	"github.com/ar3s3ru/gotrie"
	"github.com/stretchr/testify/assert"
)

func TestSimpleTrie_Insert(t *testing.T) {
	testcases := []struct {
		name string
		keys []string
		repr string
	}{
		{
			name: "Simple test with 3 similar keys",
			keys: []string{"amore", "amare", "amici"},
			repr: `"am": ["ore" -> data: <nil>,"are" -> data: <nil>,"ici" -> data: <nil>]`,
		},
		{
			name: "2 loosely similar keys and 1 totally different",
			keys: []string{"amore", "insieme", "a te"},
			repr: `"": ["a": ["more" -> data: <nil>," te" -> data: <nil>],"insieme" -> data: <nil>]`,
		},
		{
			name: "2 loosely similar keys and 2 totally different",
			keys: []string{"amore", "insieme", "a", "te"},
			repr: `"": ["a": ["more" -> data: <nil>],"insieme" -> data: <nil>,"te" -> data: <nil>]`,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			var tt gotrie.SimpleTrie
			for _, s := range test.keys {
				nt, ok := tt.Insert(s, nil)
				if !assert.True(t, ok, "not ok") {
					break
				}
				tt = nt.(gotrie.SimpleTrie)
			}
			assert.Equal(t, test.repr, tt.String(), "result mismatch")
		})
	}
}

func TestSimpleTrie_Words(t *testing.T) {
	testcases := []struct {
		name string
		keys []string
		repr string
	}{
		{
			name: "Simple test with 3 similar keys",
			keys: []string{"amore", "amare", "amici"},
		},
		{
			name: "2 loosely similar keys and 1 totally different",
			keys: []string{"amore", "insieme", "a te"},
		},
		{
			name: "2 loosely similar keys and 2 totally different",
			keys: []string{"amore", "insieme", "a", "te"},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			var tt gotrie.SimpleTrie
			for _, s := range test.keys {
				nt, ok := tt.Insert(s, nil)
				if !assert.True(t, ok, "not ok") {
					break
				}
				tt = nt.(gotrie.SimpleTrie)
			}
			words := tt.Words()
			sort.Strings(test.keys)
			sort.Strings(words)
			assert.Equal(t, test.keys, words, "words mismatch")
		})
	}
}
