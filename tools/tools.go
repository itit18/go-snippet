package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
