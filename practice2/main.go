package main

import (
	"errors"
	"fmt"
	"strings"
	"tools" // ローカルパッケージ
)

func main() {
	//標準入力を取得
	stdin, err := tools.FetchStdin()
	if err != nil {
		panic(err)
	}
	fmt.Println(stdin)

	//問題文に沿ってデータを整形
	pd, err := formatPracticeData(stdin)
	if err != nil {
		panic(err)
	}
	tools.PrintStruct(pd)

	//判定処理
	output := isProductEvenOrOdd(pd)
	fmt.Println(output)
}

// 問題文で指定された入力データ形式
type practiceData struct {
	a int
	b int
}

// 標準出力を整形して問題文で指定された入力形式に整形
func formatPracticeData(stdin []string) (data practiceData, err error) {
	sd := strings.Split(stdin[0], " ")
	data.a = tools.EasyAtoi(sd[0])
	data.b = tools.EasyAtoi(sd[1])

	//検査
	if data.a < 1 || data.a > 10000 {
		err = errors.New("aの値が不正です")
		return
	}
	if data.b < 1 || data.b > 10000 {
		err = errors.New("bの値が不正です")
		return
	}

	return
}

// 入力値の積を判定する
func isProductEvenOrOdd(pd practiceData) (output string) {
	p := pd.a * pd.b
	if (p % 2) == 0 {
		output = "Even"
	} else {
		output = "Odd"
	}

	return
}
