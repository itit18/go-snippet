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
