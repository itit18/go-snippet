package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMargePaymentPattern(t *testing.T) {
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}
	mainPattern := Make500yenPaymentPattern(500, haveCoin)
	subPattern := Make100yenPaymentPattern(250, haveCoin)
	result := MergePaymentPattern(mainPattern, subPattern, haveCoin)

	assert.Equal(t, len(result), 9)

	//50円の上限を下げてパターン数が変化するかテスト（2パターン＊2パターンを期待）
	haveCoin.coinOf50yen = 3
	mainPattern = Make500yenPaymentPattern(500, haveCoin)
	subPattern = Make100yenPaymentPattern(250, haveCoin)
	result = MergePaymentPattern(mainPattern, subPattern, haveCoin)

	assert.Equal(t, len(result), 4)

}

func TestMake500yenPaymentPattern(t *testing.T) {
	//手持ちコインはマックス
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}
	data := Make500yenPaymentPattern(1000, haveCoin)

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
	data := Make500yenPaymentPattern(6000, haveCoin)

	actual := len(data)
	expect := 213
	assert.Equal(t, expect, actual)
	assert.NotEqual(t, data[0], data[1])
}

func TestMake100yenPaymentPattern(t *testing.T) {
	//手持ちコインはマックス
	haveCoin := CoinPattern{
		coinOf500yen: 50,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}
	data := Make100yenPaymentPattern(200, haveCoin)

	actual := len(data)
	expect := 3
	assert.Equal(t, expect, actual)
	assert.NotEqual(t, data[0], data[1])
}

func TestMakeCoinPattern_手持ちコイン制限(t *testing.T) {
	//500円が不足するパターン
	haveCoin := CoinPattern{
		coinOf500yen: 1,
		coinOf100yen: 50,
		coinOf50yen:  50,
	}

	actual := MakeCoinPattern(1150, haveCoin)
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

	actual = MakeCoinPattern(750, haveCoin)
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
	actual := MakeCoinPattern(650, haveCoin)
	expect := CoinPattern{}
	expect.coinOf500yen = 1
	expect.coinOf100yen = 1
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//100円と50円のみ
	actual = MakeCoinPattern(150, haveCoin)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 1
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//50円のみ
	actual = MakeCoinPattern(50, haveCoin)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 0
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//500円のみ
	actual = MakeCoinPattern(500, haveCoin)
	expect.coinOf500yen = 1
	expect.coinOf100yen = 0
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

	//100円のみ
	actual = MakeCoinPattern(100, haveCoin)
	expect = CoinPattern{}
	expect.coinOf500yen = 0
	expect.coinOf100yen = 1
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

	//0円
	actual = MakeCoinPattern(0, haveCoin)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 0
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

}
