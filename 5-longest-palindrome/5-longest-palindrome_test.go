package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Вообще, если мы будем двигаться имеено с краев и кешировать значение IsPalindrome
// на каждом этапе провери, то есть и значения для краев с обрезанными краями, то все будет быстрее работать
// то есть в хеше мы храним не столько ахха=true , но и хх=true, да и вообще
// лучше хранить не строку, а номера символов начала и конца
//
// Если с 6 по 10 символ - палиндром, но нам не надо проверять это ещё раз.
// кроме того это значит, что 8 и девятый образуют палиндром, думаю это поможет значительно в поисках
//
// # Потому что я много раз запускаю поиск по одним и тем же строкам
//
// Плюс постарайся уложиться в стек

// возможно если мы нашли какой-то большой палиндром, то другие более короткие версии заканчиваем искать.
// думаю такой случай может быть, когда мы ищем левый вариант строки, в то время как нашли что-то справа уже
var cache [][]bool
var longestS []byte

func longestPalindrome(s string) string {
	cache = make([][]bool, len(s)) // возможно сделать и сплошной массив, главное потом с индексами не напутать)
	for i := 0; i < len(s); i++ {
		cache[i] = make([]bool, len(s))
	}
	return string(longest([]byte(s), 0, len(s)-1))
}

func longest(source []byte, left, right int) []byte {
	//print("longest", left, right)
	if len(source) < right-left {
		return longestS
	}
	if left == right {
		return source[left : right+1] // +1 потому что правая граница всегда не включается, и мы хотим ее включить
	}
	if isPalindrome(source, left, right) {
		return source[left : right+1]
	}
	leftS := longest(source, left+1, right)
	rightS := longest(source, left, right-1)
	if len(leftS) > len(rightS) {
		return leftS
	} else {
		return rightS
	}
}

func isPalindrome(source []byte, left, right int) bool {
	if cache[left][right] == true {
		return true
	}
	if len(source) == 0 {
		return false
	}
	if len(source) == 1 {
		if len(longestS) < right-left {
			longestS = source[left : right+1]
		}
		return true
	}
	// if val, ok := cache[source]; ok {
	// 	return val
	// }
	for l, r := left, right; r-l >= 1; l, r = l+1, r-1 {
		if cache[l][r] == true {
			break // чтобы записать в кеш
		}
		if source[l] != source[r] {
			return false
		}
	}
	for l, r := left, right; r-l > 1; l, r = l+1, r-1 {
		cache[l][r] = true // возможно если у нас есть более глобальный палиндром, то можно как бы с более мелкими не заморачиваться, которые внутри большого
	}
	return true
}

func TestTooLong(t *testing.T) {
	longestPalindrome("babaddtattarrattatddetartrateedredividerb")
}

func Test(t *testing.T) {
	testCases := []struct {
		input      string
		palindrome string
	}{
		{input: "badad", palindrome: "ada"},
		{input: "cbbd", palindrome: "bb"},
		{input: "sisiooppoosos", palindrome: "ooppoo"},
		{input: "sisiooppoososaaddaaccaaddaa", palindrome: "aaddaaccaaddaa"},
		//{input: "babaddtattarrattatddetartrateedredividerb", palindrome: "123"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			assert.Equal(t, tC.palindrome, longestPalindrome(tC.input))
		})
	}

}
