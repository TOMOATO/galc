package main

import (
	"testing"
)

func TestConvertIntLiteral(t *testing.T) {
	if _, e := convert("10"); e != nil {
		t.Fatalf("文字列 '10' の変換でエラーです\n%v", e)
	}
	if _, e := convert("-10"); e != nil {
		t.Fatalf("文字列 '-10' の変換でエラーです\n%v", e)
	}
	if _, e := convert("00000000999999"); e != nil {
		t.Fatalf("文字列 '00000000999999' の変換でエラーです\n%v", e)
	}
}

func TestConvertStrLiteral(t *testing.T) {
	if _, e := convert("hogehoge"); e == nil {
		t.Fatalf("文字列 'hogehoge' の変換でエラーになりませんでした")
	}
}
