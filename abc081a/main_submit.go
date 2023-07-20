// コード提出用のファイル
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin, err := FetchStdin()
	if err != nil {
		panic(err)
	}

	cellData := makeCell(stdin[0])
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

/********************************/
// 汎用的な処理をまとめておく場所
/********************************/

// 標準入力からデータを受け取る
func FetchStdin() (result []string, err error) {
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

// strconv.Atoiのエラー処理付きのシンタックスシュガー
func EasyAtoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return
}

// 構造体をデバッグ出力するときのシンタックスシュガー
func PrintStruct(st interface{}) {
	fmt.Printf("%+v\n", st)
}

// ちゃんとインポートできてるか動作確認用 / そのうち消す
func ImportDebug(name string) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Println(message)
}
