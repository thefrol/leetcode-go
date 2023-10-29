package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-strings-alternately
//
// https://github.com/thefrol/leetcode-go

func mergeAlternately(word1 string, word2 string) string {
	sum := strings.Builder{}

	var lesser, other *string
	if len(word1) > len(word2) {
		lesser = &word2
		other = &word1
	} else {
		lesser = &word1
		other = &word2
	}

	// перемешиваем
	var i int
	for i = range *lesser {
		sum.WriteByte(word1[i])
		sum.WriteByte(word2[i])
	}

	// добавляем остатки
	if i != len(*other)-1 {
		sum.WriteString((*other)[i+1 : len(*other)])
	}

	return sum.String()
}

func Test(t *testing.T) {
	testCases := []struct {
		word1 string
		word2 string
		want  string
	}{
		{word1: "abc", word2: "xyz", want: "axbycz"},
		{word1: "abc", word2: "xyzi", want: "axbyczi"},
		{word1: "abcdc", word2: "xyz", want: "axbyczdc"},
	}
	for _, tC := range testCases {
		t.Run(tC.word1+"+"+tC.word2, func(t *testing.T) {
			assert.Equal(t, tC.want, mergeAlternately(tC.word1, tC.word2))
		})
	}
}
