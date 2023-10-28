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
var longestS string
var cache map[string]bool

func longestPalindrome(s string) string {
	longestS = ""
	cache = map[string]bool{}
	longest("", "", s)
	return longestS
}

func longest(curr, left, right string) {
	if isPalindrome(curr) && len(curr) > len(longestS) {
		longestS = curr
	}
	if len(right) > 0 {
		longest(curr+right[0:0], left, right[1:])
	} else if len(left) > 0 {
		longest(left[len(left)-1:len(left)-1]+curr, left[:len(left)-1], right)
	} else {
		return
	}
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if val, ok := cache[s]; ok {
		return val
	}
	bb := []byte(s)
	for i := 0; i < len(bb)/2; i++ {
		if bb[i] != bb[len(bb)-i-1] {
			return false
		}
	}
	return true
}

func Test(t *testing.T) {
	testCases := []struct {
		input      string
		palindrome string
	}{
		{input: "badad", palindrome: "dad"},
		{input: "cbbd", palindrome: "bb"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			assert.Equal(t, tC.palindrome, longestPalindrome(tC.input))
		})
	}
}
