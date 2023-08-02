package main

import (
	"fmt"
	"tools"
)

func main() {
	var output int

	stdin, err := tools.FetchStdin()
	if err != nil {
		panic(err)
	}
	fmt.Println(stdin)

	input, err := makeAssignmentData(stdin)
	if err != nil {
		panic(err)
	}
	tools.PrintStruct(input)

	fmt.Println(output)
}

// 問題文で定義されたデータ構造
type assignmentData struct {
	coinOf500yen int
	coinOf100yen int
	coinOf50yen  int
	totalAmount  int
}

// 標準入力を問題文データに整形
func makeAssignmentData(stdin []string) (data assignmentData, err error) {
	data.coinOf500yen = tools.EasyAtoi(stdin[0])
	data.coinOf100yen = tools.EasyAtoi(stdin[1])
	data.coinOf50yen = tools.EasyAtoi(stdin[2])
	data.totalAmount = tools.EasyAtoi(stdin[3])

	return
}
