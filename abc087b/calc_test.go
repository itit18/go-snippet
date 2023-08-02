package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	actual := Sum(3, 5)
	expected := 8

	assert.Equal(t, expected, actual)

}

func TestMakeCoinPattern(t *testing.T) {
	//全硬貨1枚づつパターン
	actual := MakeCoinPattern(650)
	expect := CoinPattern{}
	expect.coinOf500yen = 1
	expect.coinOf100yen = 1
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//100円と50円のみ
	actual = MakeCoinPattern(150)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 1
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//50円のみ
	actual = MakeCoinPattern(50)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 0
	expect.coinOf50yen = 1

	assert.Equal(t, expect, actual)

	//500円のみ
	actual = MakeCoinPattern(500)
	expect.coinOf500yen = 1
	expect.coinOf100yen = 0
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

	//100円のみ
	actual = MakeCoinPattern(100)
	expect = CoinPattern{}
	expect.coinOf500yen = 0
	expect.coinOf100yen = 1
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

	//0円
	actual = MakeCoinPattern(0)
	expect.coinOf500yen = 0
	expect.coinOf100yen = 0
	expect.coinOf50yen = 0

	assert.Equal(t, expect, actual)

}
