package tools

import (
	"bufio"
	"fmt"
	"os"
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

// ちゃんとインポートできてるか動作確認用 / そのうち消す
func ImportDebug(name string) {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	fmt.Println(message)
}
