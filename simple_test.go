package gotrie_test

import (
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
	}

	for i, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			var tt gotrie.SimpleTrie
			for _, s := range test.keys {
				nt, ok := tt.Insert(s, nil)
				if !assert.True(t, ok, "test %d: not ok", i) {
					break
				}
				tt = nt.(gotrie.SimpleTrie)
			}
			assert.Equal(t, test.repr, tt.String(), "test %d: result mismatch", i)
		})
	}
}
