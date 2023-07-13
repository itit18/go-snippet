package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin, err := fetchStdin()
	if err != nil {
		panic(err)
	}
	fmt.Println(stdin)

	//引数を問題文に沿った形で整理する
	formatData, err := formatPracticeData(stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", formatData)

	//a + b + cの計算結果を得る
	sumAbc := sumABbcData(formatData)
	fmt.Println(sumAbc)

	//出力形式に沿って整形する
	output := formatOutput(sumAbc, formatData.s)

	//結果出力
	fmt.Println(output)
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

// 問題文で指定された入力形式
type practiceData struct {
	a int
	b int
	c int
	s string
}

// 標準入力から問題文で指定された入力形式への整形
func formatPracticeData(stdin []string) (formatData practiceData, err error) {
	formatData.a = EasyAtoi(stdin[0])

	split := strings.Split(stdin[1], " ")
	formatData.b = EasyAtoi(split[0])
	formatData.c = EasyAtoi(split[1])

	formatData.s = stdin[2]

	//検査
	if len(formatData.s) == 0 {
		err = errors.New("sの値が空です")
		return
	}
	if len(formatData.s) > 100 {
		err = errors.New("sの値が長すぎます")
		return
	}

	if formatData.a < 1 && formatData.a > 100 {
		err = errors.New("aの値が不正です")
		return
	}
	if formatData.b < 1 && formatData.b > 100 {
		err = errors.New("bの値が不正です")
		return
	}
	if formatData.c < 1 && formatData.c > 100 {
		err = errors.New("cの値が不正です")
		return
	}

	return
}

// strconvのエラー処理付きのシンタックスシュガー
func EasyAtoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return
}

func sumABbcData(pd practiceData) (sum int) {
	sum = pd.a + pd.b + pd.c
	return
}

func formatOutput(i int, s string) (output string) {
	output = fmt.Sprintf("%d %s", i, s)
	return
}
