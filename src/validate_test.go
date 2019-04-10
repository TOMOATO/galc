package main

import (
	"testing"
)

func TestValidateLength(t *testing.T) {
	if e := validate([]string{"a", "b", "c", "d"}); e != nil {
		t.Fatalf("長さ4つの配列でエラーが出ました%v", e)
	}
	if e := validate([]string{"a", "b"}); e == nil {
		t.Fatalf("長さ2つの配列でエラーが出ませんでした%v", e)
	}
}
