package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// validate関数で，与えられたコマンドライン引数が受け付けるべきものか判断
	if err := validate(args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	// clac関数で計算
	if result, err := calc(args[1], args[2], args[3]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}

// calc関数はcommand次第で
func calc(command, lhstr, rhstr string) (result int, err error) {
	//komento ireteyaru ze!
	// 文字列で与えられたlhsとrhsを数値に変換し，エラーチェック
	lhs, err1 := convert(lhstr)
	rhs, err2 := convert(rhstr)
	if err1 != nil {
		return 0, errors.New("LHSでエラー！\n" + err1.Error())
	}
	if err2 != nil {
		return 0, errors.New("RHSでエラー！\n" + err1.Error())
	}
	// commandの内容に従って分岐
	switch command {
	case "add":
		result, err = add(lhs, rhs)
	}
	return
}
