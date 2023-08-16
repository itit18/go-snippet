package main

// 　条件を満たすコインの組み合わせパターン構造体
type CoinPattern struct {
	coinOf500yen int
	coinOf100yen int
	coinOf50yen  int
}

func CalcPaymentPattarn(totalAmont int, haveCoin CoinPattern) (paymentPatterns []CoinPattern) {
	var totals map[string]int = make(map[string]int)

	//手持ち金額が支払額を上回っているかをチェック
	if !isPay(totalAmont, haveCoin) {
		return
	}

	//合計金額を500で割り切れる数と、端数に分解する
	totals["500"], totals["other"] = divisionWithRemainder(totalAmont, 500)
	// fmt.Println(totals["500"], totals["other"])

	//500円で支払える額について支払いパターンを算出する
	payment500, failFlg := Make500yenPaymentPattern(totals["500"]*500, haveCoin)
	if failFlg {
		return
	}

	//残りの端数について、100円と50円の支払いパターンを算出する
	payment100, failFlg := Make100yenPaymentPattern(totals["other"], haveCoin)
	if failFlg {
		return
	}

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
			//重複チェック
			if existsPttern(marge, n) {
				continue
			}
			// tools.PrintStruct(n)
			marge = append(marge, n)
		}
	}

	//空配列があったときの処理パターン
	if len(main) == 0 {
		marge = sub
	}
	if len(sub) == 0 {
		marge = main
	}

	return
}

// 100円で支払えるパターンについてパターンを生成
// failFlg = 論理的に支払いすることが不可能なパターンを伝えるフラグ
func Make100yenPaymentPattern(totalAmont int, haveCoin CoinPattern) (paymentPatterns []CoinPattern, failFlg bool) {
	failFlg = false
	if totalAmont == 0 {
		return
	}

	basePattern, failFlg := MakeCoinPattern(totalAmont, haveCoin)
	if failFlg {
		return
	}
	paymentPatterns = append(paymentPatterns, basePattern)

	//100円→50円の両替
	for i := 0; i < len(paymentPatterns); i++ {
		base, fail := exchange100yenTo50yen(paymentPatterns[i], haveCoin)
		if fail {
			continue
		}

		paymentPatterns = append(paymentPatterns, base)
	}

	return
}

// 500円で支払えるパターンについてパターンを生成
// failFlg = 論理的に支払いすることが不可能なパターンを伝えるフラグ
func Make500yenPaymentPattern(totalAmont int, haveCoin CoinPattern) (paymentPatterns []CoinPattern, failFlg bool) {
	failFlg = false

	if totalAmont == 0 {
		return
	}

	basePattern, failFlg := MakeCoinPattern(totalAmont, haveCoin)
	if failFlg {
		return
	}

	// tools.PrintStruct(basePattern)
	paymentPatterns = append(paymentPatterns, basePattern)

	//rangeでループさせるとループ内のappendの結果を見てくれないのでlen()でループさせている
	for i := 0; i < len(paymentPatterns); i++ {
		//500→100変換の両替
		exchange100, exchangeFail := exchange500yenTo100yen(paymentPatterns[i], haveCoin)
		// すでに同じパターンが登録済みでないかをチェック
		if existsPttern(paymentPatterns, exchange100) {
			exchangeFail = true
		}

		if !exchangeFail {
			paymentPatterns = append(paymentPatterns, exchange100)
		}

		//500→50の両替
		exchange50, exchangeFail := exchange500yenTo50yen(paymentPatterns[i], haveCoin)
		// すでに同じパターンが登録済みでないかをチェック
		if existsPttern(paymentPatterns, exchange50) {
			exchangeFail = true
		}

		if !exchangeFail {
			paymentPatterns = append(paymentPatterns, exchange50)
		}
	}

	//生成したパターンからさらに100円→50円の両替パターンを生やしていく
	var addPatteerns []CoinPattern
	for i := 0; i < len(paymentPatterns); i++ {
		if paymentPatterns[i].coinOf100yen == 0 {
			continue
		}

		exchange, exchangeFail := exchange100yenTo50yen(paymentPatterns[i], haveCoin)
		if existsPttern(addPatteerns, exchange) {
			exchangeFail = true
		}
		if existsPttern(paymentPatterns, exchange) {
			exchangeFail = true
		}

		if !exchangeFail {
			paymentPatterns = append(paymentPatterns, exchange)
		}
	}

	return
}

// 合計金額を満たすコイン選択パターンを1つ生成する
func MakeCoinPattern(totalAmont int, haveCoin CoinPattern) (result CoinPattern, failFlg bool) {
	var remainder int
	failFlg = false

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
	if haveCoin.coinOf50yen < result.coinOf50yen {
		result.coinOf50yen = haveCoin.coinOf50yen
	}

	//支払いパターンが存在しなかったときの処理
	sumCoin := (result.coinOf500yen * 500) + (result.coinOf100yen * 100) + (result.coinOf50yen * 50)
	if totalAmont != sumCoin {
		failFlg = true
		result = CoinPattern{} //支払額に満たない場合、初期化して返す
	}

	return
}

///////////////////////////////
//  Privete
///////////////////////////////

// 除算の結果と余りを返す
func divisionWithRemainder(dividend, divisor int) (quotient, remainder int) {
	if dividend == 0 {
		return
	}
	quotient = dividend / divisor
	remainder = dividend % divisor

	return
}

// 重複するパターンが登録済みかをチェックする
func existsPttern(patterns []CoinPattern, check CoinPattern) bool {
	for _, p := range patterns {
		if equalCoinPattern(p, check) {
			return true
		}
	}
	return false
}

// CoinPatternが同一かをチェックする
func equalCoinPattern(a CoinPattern, b CoinPattern) bool {
	if a.coinOf500yen != b.coinOf500yen {
		return false
	}
	if a.coinOf100yen != b.coinOf100yen {
		return false
	}
	if a.coinOf50yen != b.coinOf50yen {
		return false
	}
	return true
}

// 両替処理 / 500円→100円 / 両替不可能な場合はfailFlgで検知
func exchange500yenTo100yen(c CoinPattern, haveCoin CoinPattern) (r CoinPattern, failFlg bool) {
	r = c
	r.coinOf500yen = c.coinOf500yen - 1
	r.coinOf100yen = c.coinOf100yen + 5

	// 計算後のコイン数が制約違反してないかをチェック
	failFlg = false
	if r.coinOf500yen < 0 {
		failFlg = true
		r = CoinPattern{} //結果が違反してるので初期化して空にしておく
	}
	if haveCoin.coinOf100yen < r.coinOf100yen {
		failFlg = true
		r = CoinPattern{}
	}

	return
}

// 両替処理 / 500円→100円 / 両替不可能な場合はfailFlgで検知
func exchange500yenTo50yen(c CoinPattern, haveCoin CoinPattern) (r CoinPattern, failFlg bool) {
	r = c
	r.coinOf500yen = c.coinOf500yen - 1
	r.coinOf50yen = c.coinOf50yen + 10

	failFlg = false
	if r.coinOf500yen < 0 {
		failFlg = true
		r = CoinPattern{} //結果が違反してるので初期化して空にしておく
	}
	if haveCoin.coinOf50yen < r.coinOf50yen {
		failFlg = true
		r = CoinPattern{}
	}

	return
}

// 両替処理 / 100円→　50円
func exchange100yenTo50yen(c CoinPattern, haveCoin CoinPattern) (r CoinPattern, failFlg bool) {
	r = c
	r.coinOf100yen = c.coinOf100yen - 1
	r.coinOf50yen = c.coinOf50yen + 2

	failFlg = false
	if r.coinOf100yen < 0 {
		failFlg = true
		r = CoinPattern{} //結果が違反してるので初期化して空にしておく
	}
	if haveCoin.coinOf50yen < r.coinOf50yen {
		failFlg = true
		r = CoinPattern{}
	}

	return
}

// 支払い能力の可否チェック
func isPay(totalAmont int, haveCoin CoinPattern) bool {
	sumCoin := (haveCoin.coinOf500yen * 500) + (haveCoin.coinOf100yen * 100) + (haveCoin.coinOf50yen * 50)
	return totalAmont <= sumCoin
}
