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
