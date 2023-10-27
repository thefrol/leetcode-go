package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var longestS string
var cache map[string]bool

func longestPalindrome(s string) string {
	longestS = ""
	cache = map[string]bool{}
	return longest(s)
}

func longest(s string) string {
	if len(s) < len(longestS) {
		return ""
	}
	if isPalindrome(s) {
		if len(s) > len(longestS) {
			longestS = s
		}
		longestS = s
		return s
	}
	res1 := longest(s[:len(s)-1])
	res2 := longest(s[1:])
	if len(res1) > len(res2) {
		return res1
	}
	return res2
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
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
