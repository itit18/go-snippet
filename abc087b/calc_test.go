package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPaymentPattarn(t *testing.T) {
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}
	result := CalcPaymentPattarn(7500, haveCoin)
	expect := 266

	assert.Equal(t, expect, len(result))

	haveCoin = CoinPattern{
		coinOf500yen: 2,
		coinOf100yen: 2,
		coinOf50yen:  2,
	}
	result = CalcPaymentPattarn(1550, haveCoin)
	expect = 0

	assert.Equal(t, expect, len(result))

	haveCoin = CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 2,
		coinOf50yen:  0,
	}
	result = CalcPaymentPattarn(1550, haveCoin)
	expect = 0

	assert.Equal(t, expect, len(result))

}

func TestMargePaymentPattern_両替最大数パターン(t *testing.T) {
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}

	mainPattern, _ := Make500yenPaymentPattern(7500, haveCoin)
	subPattern, _ := Make100yenPaymentPattern(0, haveCoin)

	result := MergePaymentPattern(mainPattern, subPattern, haveCoin)
	expect := 266

	assert.Equal(t, expect, len(result))

}

func TestMargePaymentPattern_生成パターンゼロ(t *testing.T) {
	haveCoin := CoinPattern{
		coinOf500yen: 5,
		coinOf100yen: 1,
		coinOf50yen:  0,
	}

	mainPattern, _ := Make500yenPaymentPattern(0, haveCoin)
	subPattern, _ := Make100yenPaymentPattern(150, haveCoin)
	result := MergePaymentPattern(mainPattern, subPattern, haveCoin)
	expect := 0

	assert.Equal(t, expect, len(result))

}

func TestMargePaymentPattern_最小値の正常系(t *testing.T) {
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}
	mainPattern, _ := Make500yenPaymentPattern(500, haveCoin)
	subPattern, _ := Make100yenPaymentPattern(250, haveCoin)
	result := MergePaymentPattern(mainPattern, subPattern, haveCoin)
	expect := 3 + 8 // 500円1枚のパターン数　+ 500円0枚のパターン数の和

	assert.Equal(t, expect, len(result))

	//50円の上限を下げてパターン数が変化するかテスト（2パターン＊2パターンを期待）
	haveCoin.coinOf50yen = 3
	mainPattern, _ = Make500yenPaymentPattern(500, haveCoin)
	subPattern, _ = Make100yenPaymentPattern(250, haveCoin)
	result = MergePaymentPattern(mainPattern, subPattern, haveCoin)
	expect = 2 + 2 // 50円が取れるパターンが2パターンのみになっているはず

	assert.Equal(t, expect, len(result))

	//500円で完全に割り切れるときの挙動チェック
	haveCoin.coinOf50yen = 3
	mainPattern, _ = Make500yenPaymentPattern(500, haveCoin)
	subPattern, _ = Make100yenPaymentPattern(0, haveCoin)
	result = MergePaymentPattern(mainPattern, subPattern, haveCoin)
	expect = 3

	assert.Equal(t, expect, len(result))

	//500円以下の挙動チェック
	haveCoin.coinOf50yen = 3
	mainPattern, _ = Make500yenPaymentPattern(0, haveCoin)
	subPattern, _ = Make100yenPaymentPattern(250, haveCoin)
	result = MergePaymentPattern(mainPattern, subPattern, haveCoin)
	expect = 2

	assert.Equal(t, expect, len(result))
}

func TestMake500yenPaymentPattern(t *testing.T) {
	//手持ちコインはマックス
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}
	data, _ := Make500yenPaymentPattern(1000, haveCoin)

	actual := len(data)
	expect := 6 + 4 + 8
	assert.Equal(t, expect, actual)
	assert.NotEqual(t, data[0], data[1])
}

func TestMake500yenPaymentPattern_入力例テスト(t *testing.T) {
	haveCoin := CoinPattern{
		coinOf500yen: 30,
		coinOf100yen: 40,
		coinOf50yen:  50,
	}
	data, _ := Make500yenPaymentPattern(6000, haveCoin)

	actual := len(data)
	expect := 213
	assert.Equal(t, expect, actual)
	assert.NotEqual(t, data[0], data[1])

	//合計支払額が0円のパターン
	data, _ = Make500yenPaymentPattern(0, haveCoin)
	actual = len(data)
	expect = 0
	assert.Equal(t, expect, actual)

	//合計支払額が手持ちコインを超過してるパターン
	haveCoin = CoinPattern{
		coinOf500yen: 1,
		coinOf100yen: 1,
		coinOf50yen:  1,
	}
	data, flg := Make500yenPaymentPattern(1150, haveCoin)
	actual = len(data)
	expect = 0
	assert.Equal(t, expect, actual)
	assert.Equal(t, true, flg)
}

func TestMake100yenPaymentPattern(t *testing.T) {
	//手持ちコインはマックス
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}
	data, _ := Make100yenPaymentPattern(200, haveCoin)

	actual := len(data)
	expect := 3
	assert.Equal(t, expect, actual)
	assert.NotEqual(t, data[0], data[1])

	//合計支払額が0円のパターン
	data, _ = Make100yenPaymentPattern(0, haveCoin)
	actual = len(data)
	expect = 0
	assert.Equal(t, expect, actual)

	//合計支払い額が手持ちコインを超過しているパターン
	haveCoin = CoinPattern{
		coinOf500yen: 1,
		coinOf100yen: 1,
		coinOf50yen:  1,
	}
	data, flg := Make100yenPaymentPattern(1150, haveCoin)
	actual = len(data)
	expect = 0
	assert.Equal(t, expect, actual)
	assert.Equal(t, true, flg)
}

func TestMakeCoinPattern_手持ちコイン制限(t *testing.T) {
	//500円が不足するパターン
	haveCoin := CoinPattern{
		coinOf500yen: 1,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}

	actual, _ := MakeCoinPattern(1150, haveCoin)
	expect := CoinPattern{}
	expect.coinOf500yen = 1
	expect.coinOf100yen = 6
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//100円が不足するパターン
	haveCoin = CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 1,
		coinOf50yen:  50,
	}

	actual, _ = MakeCoinPattern(750, haveCoin)
	expect = CoinPattern{}
	expect.coinOf500yen = 1
	expect.coinOf100yen = 1
	expect.coinOf50yen = 3

	assert.Equal(t, expect, actual)

}

func TestMakeCoinPattern_手持ちコイン最大(t *testing.T) {
	//手持ちコインはマックス
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}

	//全硬貨1枚づつパターン
	actual, _ := MakeCoinPattern(650, haveCoin)
	expect := CoinPattern{}
	expect.coinOf500yen = 1
	expect.coinOf100yen = 1
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//100円と50円のみ
	actual, _ = MakeCoinPattern(150, haveCoin)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 1
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//50円のみ
	actual, _ = MakeCoinPattern(50, haveCoin)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 0
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//500円のみ
	actual, _ = MakeCoinPattern(500, haveCoin)
	expect.coinOf500yen = 1
	expect.coinOf100yen = 0
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

	//100円のみ
	actual, _ = MakeCoinPattern(100, haveCoin)
	expect = CoinPattern{}
	expect.coinOf500yen = 0
	expect.coinOf100yen = 1
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

	//0円
	actual, _ = MakeCoinPattern(0, haveCoin)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 0
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

}
