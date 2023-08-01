package main

import (
	"fmt"
	"tools"
)

func main() {
	fmt.Println("aaa")
	stdin, err := tools.FetchStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(stdin)
}
