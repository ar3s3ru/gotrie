package gotrie_test

import (
	"sort"
	"testing"

	"github.com/ar3s3ru/gotrie"

	"github.com/stretchr/testify/assert"
)

func TestNewCommonPrefixTree(t *testing.T) {
	tr := gotrie.NewCommonPrefixTree("hello", "world")
	assert.NotNil(t, tr, "nil created tree")
	v, ok := tr.Get("hello")
	assert.True(t, ok, "not found with Get()")
	assert.Equal(t, "world", v, "meta data mismatch")
}

func TestCommonPrefixTree_Insert(t *testing.T) {
	testcases := []struct {
		name string
		keys []string
		repr string
	}{
		{
			name: "No key",
			repr: `<uninit>`,
		},
		{
			name: "One key",
			keys: []string{"key"},
			repr: `"key" -> data: <nil>`,
		},
		{
			name: "Same keys",
			keys: []string{"hello", "hello"},
			repr: `"hello" -> data: <nil>`,
		},
		{
			name: "Simple test with 3 similar keys",
			keys: []string{"love", "live", "friends"},
			repr: `"": ["l": ["ove" -> data: <nil>,"ive" -> data: <nil>],"friends" -> data: <nil>]`,
		},
		{
			name: "2 likely similar keys and 1 totally different",
			keys: []string{"love", "and", "live together"},
			repr: `"": ["l": ["ove" -> data: <nil>,"ive together" -> data: <nil>],"and" -> data: <nil>]`,
		},
		{
			name: "2 likey similar keys and 2 totally different",
			keys: []string{"love", "and", "live", "together"},
			repr: `"": ["l": ["ove" -> data: <nil>,"ive" -> data: <nil>],"and" -> data: <nil>,"together" -> data: <nil>]`,
		},
		{
			name: "a couple of 2 likely similar keys",
			keys: []string{"love", "and", "live", "alone"},
			repr: `"": ["l": ["ove" -> data: <nil>,"ive" -> data: <nil>],"a": ["nd" -> data: <nil>,"lone" -> data: <nil>]]`,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			var tt *gotrie.CommonPrefixTree
			for _, s := range test.keys {
				nt, _ := tt.Insert(s, nil)
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
	}{
		{
			name: "No key",
		},
		{
			name: "One key",
			keys: []string{"key"},
		},
		{
			name: "Same keys",
			keys: []string{"hello"},
		},
		{
			name: "Simple test with 3 similar keys",
			keys: []string{"love", "live", "friends"},
		},
		{
			name: "2 likely similar keys and 1 totally different",
			keys: []string{"love", "and", "live together"},
		},
		{
			name: "2 likey similar keys and 2 totally different",
			keys: []string{"love", "and", "live", "together"},
		},
		{
			name: "a couple of 2 likely similar keys",
			keys: []string{"love", "and", "live", "alone"},
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
			name:    "Query at an empty tree",
			results: []result{{key: "test", ok: false}},
		},
		{
			name:    "Simple test with 3 similar keys",
			keys:    []string{"amore", "amare", "amici"},
			results: []result{{key: "amore", ok: true}, {key: "amori", ok: false}, {key: "a", ok: false}},
		},
		{
			name:    "2 likely similar keys and 1 totally different",
			keys:    []string{"amore", "insieme", "a te"},
			results: []result{{key: "a te", ok: true}, {key: "amori", ok: false}, {key: "ins", ok: false}},
		},
		{
			name:    "2 likely similar keys and 2 totally different",
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
