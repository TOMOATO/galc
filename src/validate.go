package main

import (
	"errors"
)

func validate(args []string) error {
	// 長さが4未満 (command, lhs, rhsが全て与えられていない) ならエラー
	if len(args) < 4 {
		return errors.New("引数の長さが足りません (4つ以上)")
	}
	return nil
}
