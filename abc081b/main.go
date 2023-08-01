package main

import (
	"fmt"
	"strings"
	"tools"
)

func main() {
	stdin, err := tools.FetchStdin()
	if err != nil {
		panic(err)
	}
	fmt.Println(stdin)

	data, err := makeAssignmentData(stdin)
	if err != nil {
		panic(err)
	}
	tools.PrintStruct(data)

	output := countAllHalve(data.values)
	fmt.Println(output)

}

// 問題文で指定された入力値の形式
type assignmentData struct {
	maxvalue int
	values   []int
}

func makeAssignmentData(stdin []string) (result assignmentData, err error) {
	result.maxvalue = tools.EasyAtoi(stdin[0])
	sp := strings.Split(stdin[1], " ")
	for _, d := range sp {
		result.values = append(result.values, tools.EasyAtoi(d))
	}

	return
}

// 「配列内の要素を2で割ったものに置き換える」の処理回数を数える
func countAllHalve(values []int) (counter int) {
	//一応繰り返し動作の上限を決めておく、chatGPTいわく最大数は30らしい
	for i := 0; i < 30; i++ {
		// 先に配列内が全て偶数であるかを確認しておく
		if !verifyAllEven(values) {
			break
		}

		values = halveArrayValues(values)
		counter++
	}
	return

}

// 配列内の要素を2で割ったものに置き換える
func halveArrayValues(values []int) (halveValues []int) {
	for _, v := range values {
		halveValues = append(halveValues, v/2)
	}

	return
}

func verifyAllEven(values []int) (result bool) {
	result = true
	for _, v := range values {
		if v == 0 {
			result = false
			break
		}
		if v%2 == 1 {
			result = false
			break
		}
	}

	return
}
