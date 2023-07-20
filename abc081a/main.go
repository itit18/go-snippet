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

	//fmt.Println(stdin)
	cellData := makeCell(stdin[0])
	//tools.PrintStruct(cellData)
	result := countMarble(cellData)

	fmt.Println(result)
}

type cell struct {
	values [3]bool
}

func makeCell(s string) (c cell) {
	split_s := strings.Split(s, "")
	for i := 0; i < len(c.values); i++ {
		if split_s[i] == "1" {
			c.values[i] = true
		}
	}

	return
}

func countMarble(c cell) (count int) {
	for i := 0; i < len(c.values); i++ {
		if c.values[i] {
			count++
		}
	}

	return
}
