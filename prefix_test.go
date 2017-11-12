package gotrie_test

import (
	"sort"
	"testing"

	"github.com/ar3s3ru/gotrie"
	"github.com/stretchr/testify/assert"
)

func TestCommonPrefixTree_Insert(t *testing.T) {
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
		{
			name: "Same keys",
			keys: []string{"hello"},
			repr: `"hello" -> data: <nil>`,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			var tt *gotrie.CommonPrefixTree
			for _, s := range test.keys {
				nt, ok := tt.Insert(s, nil)
				if !assert.True(t, ok, "not ok") {
					break
				}
				tt = nt.(*gotrie.CommonPrefixTree)
			}
			assert.Equal(t, test.repr, tt.String(), "result mismatch")
		})
	}
}

func TestCommonPrefixTree_Words(t *testing.T) {
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
			var tt *gotrie.CommonPrefixTree
			for _, s := range test.keys {
				nt, ok := tt.Insert(s, nil)
				if !assert.True(t, ok, "not ok") {
					break
				}
				tt = nt.(*gotrie.CommonPrefixTree)
			}
			words := tt.Keys()
			sort.Strings(test.keys)
			sort.Strings(words)
			assert.Equal(t, test.keys, words, "words mismatch")
		})
	}
}

func TestCommonPrefixTree_Get(t *testing.T) {
	type result struct {
		key string
		ok  bool
	}
	testcases := []struct {
		name    string
		keys    []string
		results []result
	}{
		{
			name:    "Simple test with 3 similar keys",
			keys:    []string{"amore", "amare", "amici"},
			results: []result{{key: "amore", ok: true}, {key: "amori", ok: false}, {key: "a", ok: false}},
		},
		{
			name:    "2 loosely similar keys and 1 totally different",
			keys:    []string{"amore", "insieme", "a te"},
			results: []result{{key: "a te", ok: true}, {key: "amori", ok: false}, {key: "ins", ok: false}},
		},
		{
			name:    "2 loosely similar keys and 2 totally different",
			keys:    []string{"amore", "insieme", "a", "te"},
			results: []result{{key: "insieme", ok: true}, {key: "amori", ok: false}},
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			var tt *gotrie.CommonPrefixTree
			for _, s := range test.keys {
				nt, ok := tt.Insert(s, nil)
				if !assert.True(t, ok, "not ok") {
					break
				}
				tt = nt.(*gotrie.CommonPrefixTree)
			}
			for _, r := range test.results {
				_, ok := tt.Get(r.key)
				assert.Equal(t, r.ok, ok, "value check mismatch")
			}
		})
	}
}
