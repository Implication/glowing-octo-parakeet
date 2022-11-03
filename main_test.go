package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5 word6")

	if exp := count(b, false, false); exp != 6 {
		t.Errorf("Expected %d got %d instead. \n", 6, exp)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5 word6")

	if exp := count(b, true, false); exp != 1 {
		t.Errorf("Expected %d got %d instead. \n", 6, exp)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5 word6")

	if exp := count(b, false, true); exp != 30 {
		t.Errorf("Expected %d got %d instead. \n", 30, exp)
	}
}
