package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5 word6")

	if count(b) != 6 {
		t.Errorf("Expected %d got %d instead. \n", 6, count(b))
	}
}
