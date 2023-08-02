package main

// これはユニットテストを試運転するための仮コード / TODO: あとで消す
func Sum(a, b int) int {
	return a + b
}

// 　条件を満たすコインの組み合わせパターン構造体
type CoinPattern struct {
	coinOf500yen int
	coinOf100yen int
	coinOf50yen  int
}

func MakeAllPattern(totalAmont int) (pattarns []CoinPattern) {
	basePattern := MakeCoinPattern(totalAmont)

	//500→100変換のパターンを配列へ

	//100→50変換のパターンを配列へ

	return
}

// 500コインを100コインに変換する
func Divide500Coin(basePattern CoinPattern) (newPatterns CoinPattern) {
	newPatterns.coinOf500yen = basePattern.coinOf500yen - 1
	newPatterns.coinOf100yen = basePattern.coinOf100yen + 5
	newPatterns.coinOf50yen = basePattern.coinOf50yen

	return
}

// 100コインを50コインに変換する
func Divide100Coin(basePattern CoinPattern) (newPatterns CoinPattern) {
	newPatterns.coinOf500yen = basePattern.coinOf500yen
	newPatterns.coinOf100yen = basePattern.coinOf100yen - 1
	newPatterns.coinOf50yen = basePattern.coinOf50yen + 2

	return
}

// 合計金額を満たすコイン選択パターンを1つ生成する
func MakeCoinPattern(totalAmont int) (result CoinPattern) {
	var remainder int
	result.coinOf500yen, remainder = divisionWithRemainder(totalAmont, 500)
	result.coinOf100yen, remainder = divisionWithRemainder(remainder, 100)
	result.coinOf50yen, _ = divisionWithRemainder(remainder, 50)

	return
}

// 除算の結果と余りを返す
func divisionWithRemainder(dividend, divisor int) (quotient, remainder int) {
	if dividend == 0 {
		return
	}
	quotient = dividend / divisor
	remainder = dividend % divisor

	return
}

/*
パターンメモ
500円：16枚～0枚
100円：26枚～0枚
100円：50枚～0枚
*/
