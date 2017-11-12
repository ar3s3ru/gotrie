package gotrie_test

import (
	"bufio"
	"io"
	"os"
	"testing"
)

var words []string

func TestMain(m *testing.M) {
	f, err := os.Open("english-words/words.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		words = append(words, s)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
	m.Run()
}
