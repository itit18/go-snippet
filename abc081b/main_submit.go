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
	stdin, err := FetchStdin()
	if err != nil {
		panic(err)
	}
	// fmt.Println(stdin)

	data, err := makeAssignmentData(stdin)
	if err != nil {
		panic(err)
	}
	PrintStruct(data)

	output := countAllHalve(data.values)
	fmt.Println(output)

}

// 問題文で指定された入力値の形式
type assignmentData struct {
	maxvalue int
	values   []int
}

func makeAssignmentData(stdin []string) (result assignmentData, err error) {
	result.maxvalue = EasyAtoi(stdin[0])
	sp := strings.Split(stdin[1], " ")
	for _, d := range sp {
		result.values = append(result.values, EasyAtoi(d))
	}

	if result.maxvalue > 200 || result.maxvalue < 1 {
		err = errors.New("maxvalueが不正な値です")
		return
	}

	if len(result.values) != result.maxvalue {
		err = errors.New("maxvalueとvaluesが一致しません")
		return
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
