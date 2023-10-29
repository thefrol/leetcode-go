package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-strings-alternately
//
// https://github.com/thefrol/leetcode-go

// V1 вместо стрингбилдера использовать аппенд
// V2 создать слайс байтов нужной капасити
// V3 вместо string использовать []byte, что по сути одно и то же
// но меньше кринжовых преобразований

// три вот этих пункта исправил, и вместо скорости в 33%, обошел получил 75% сдавших
// работу этим идиоматичным кодом

func mergeAlternately(word1 string, word2 string) string {
	sum := make([]byte, 0, len(word1)+len(word2))

	var lesser, other []byte
	if len(word1) > len(word2) {
		lesser = []byte(word2)
		other = []byte(word1)
	} else {
		lesser = []byte(word1)
		other = []byte(word2)
	}

	// перемешиваем
	var i int
	for i = range lesser {
		sum = append(sum, word1[i], word2[i])
	}

	// добавляем остатки
	if i != len(other)-1 {
		sum = append(sum, other[i+1:]...)
	}

	return string(sum)
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
