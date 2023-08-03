package main

/*
パターンメモ
500円：10枚:5千円を100円と交換可能, 5枚2500円を50円と交換可能
100円：25枚:2500円を50円と交換可能
50円： 50枚～0枚
*/

// 　条件を満たすコインの組み合わせパターン構造体
type CoinPattern struct {
	coinOf500yen int
	coinOf100yen int
	coinOf50yen  int
}

func CalcPaymentPattarn(totalAmont int, haveCoin CoinPattern) (paymentPatterns []CoinPattern) {
	var totals map[string]int = make(map[string]int)

	//合計金額を500で割り切れる数と、端数に分解する
	totals["500"], totals["other"] = divisionWithRemainder(totalAmont, 500)
	//fmt.Println(totals["500"], totals["other"])

	//500円で支払える額について支払いパターンを算出する
	payment500 := Make500yenPaymentPattern(totals["500"], haveCoin)
	//残りの端数について、100円と50円の支払いパターンを算出する
	payment100 := Make100yenPaymentPattern(totals["other"], haveCoin)

	// fmt.Println(len(payment100))
	// fmt.Println(len(payment500))

	//合成して手持ちコイン数を上回ってしまうパターンを枝切り
	paymentPatterns = MergePaymentPattern(payment500, payment100, haveCoin)

	return
}

// 2つの支払いパターンを合成する
func MergePaymentPattern(main []CoinPattern, sub []CoinPattern, haveCoin CoinPattern) (marge []CoinPattern) {
	for _, m := range main {
		for _, s := range sub {
			//合成したパターンを新規作成
			n := CoinPattern{}
			n.coinOf500yen = m.coinOf500yen
			n.coinOf100yen = m.coinOf100yen + s.coinOf100yen
			n.coinOf50yen = m.coinOf50yen + s.coinOf50yen

			//mainとsubのsumが手持ちコイン数を超えていないか確認する
			if haveCoin.coinOf100yen < n.coinOf100yen {
				continue
			}
			if haveCoin.coinOf50yen < n.coinOf50yen {
				continue
			}
			marge = append(marge, n)
		}
	}

	return
}

// 100円で支払えるパターンについてパターンを生成
func Make100yenPaymentPattern(totalAmont int, haveCoin CoinPattern) (paymentPatterns []CoinPattern) {
	basePattern := MakeCoinPattern(totalAmont, haveCoin)
	paymentPatterns = append(paymentPatterns, basePattern)

	//100円→50円の両替
	for i := 0; i < len(paymentPatterns); i++ {
		base := paymentPatterns[i]
		base.coinOf100yen = base.coinOf100yen - 1
		base.coinOf50yen = base.coinOf50yen + 2
		if base.coinOf100yen < 0 {
			continue
		}
		if haveCoin.coinOf50yen < basePattern.coinOf50yen {
			continue
		}

		paymentPatterns = append(paymentPatterns, base)
	}

	return
}

// 500円で支払えるパターンについてパターンを生成
func Make500yenPaymentPattern(totalAmont int, haveCoin CoinPattern) (paymentPatterns []CoinPattern) {
	basePattern := MakeCoinPattern(totalAmont, haveCoin)
	paymentPatterns = append(paymentPatterns, basePattern)

	//rangeでループさせるとループ内のappendの結果を見てくれないので変な実装になってる
	for i := 0; i < len(paymentPatterns); i++ {
		//500→100変換の両替
		base := paymentPatterns[i]
		base.coinOf500yen = base.coinOf500yen - 1
		base.coinOf100yen = base.coinOf100yen + 5
		// 計算後のコイン数が制約違反してないかをチェック
		if base.coinOf500yen < 0 {
			continue
		}
		if haveCoin.coinOf100yen < base.coinOf100yen {
			continue
		}

		paymentPatterns = append(paymentPatterns, base)

		//500→50の両替
		base = paymentPatterns[i]
		base.coinOf500yen = base.coinOf500yen - 1
		base.coinOf50yen = base.coinOf100yen + 10
		// 計算後のコイン数が制約違反してないかをチェック
		if haveCoin.coinOf50yen < base.coinOf50yen {
			continue
		}

		paymentPatterns = append(paymentPatterns, base)
	}

	return
}

// 合計金額を満たすコイン選択パターンを1つ生成する
func MakeCoinPattern(totalAmont int, haveCoin CoinPattern) (result CoinPattern) {
	var remainder int
	//単純な理想値で500円の支払い枚数を算出
	result.coinOf500yen, remainder = divisionWithRemainder(totalAmont, 500)
	//理想値よりも手持ちコインが少ないときの処理
	if haveCoin.coinOf500yen < result.coinOf500yen {
		diff := result.coinOf500yen - haveCoin.coinOf500yen
		result.coinOf500yen = haveCoin.coinOf500yen
		remainder = diff*500 + remainder
	}

	//500円と同じ処理で枚数算出
	result.coinOf100yen, remainder = divisionWithRemainder(remainder, 100)
	if haveCoin.coinOf100yen < result.coinOf100yen {
		diff := result.coinOf100yen - haveCoin.coinOf100yen
		result.coinOf100yen = haveCoin.coinOf100yen
		remainder = diff*100 + remainder
	}

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
