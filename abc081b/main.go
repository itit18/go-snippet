package main

import (
	"fmt"
	"tools"
)

func main() {
	stdin, err := tools.FetchStdin()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(stdin)

}
