package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r, e := add(1, 2)
	if e != nil {
		t.Fatal("1 + 2 でエラーが発生")
	}
	if r != 3 {
		t.Fatalf("1 + 2 = %d となりました", r)
	}
}
