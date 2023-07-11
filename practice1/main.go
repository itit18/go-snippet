package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input_text, err := fetchStdin()
	if err != nil {
		panic(err)
	}

	fmt.Println(input_text)
}

func fetchStdin() (result []string, err error) {
	sc := bufio.NewScanner(os.Stdin)
	if sc.Err() != nil {
		err = sc.Err()
		return
	}

	for sc.Scan() {
		result = append(result, sc.Text())
	}

	return
}
