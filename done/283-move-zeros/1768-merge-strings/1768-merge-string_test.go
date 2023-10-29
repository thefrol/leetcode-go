package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-strings-alternately
//
// https://github.com/thefrol/leetcode-go

// https://leetcode.com/problems/merge-strings-alternately/solutions/4224142/go-fast-beautiful-ideomatic/

// V1 вместо стрингбилдера использовать аппенд
// V2 создать слайс байтов нужной капасити
// V3 вместо string использовать []byte, что по сути одно и то же
// но меньше кринжовых преобразований

// три вот этих пункта исправил, и вместо скорости в 33%, обошел получил 75% сдавших
// работу этим идиоматичным кодом

// V4 Установить меньшую и болшую строку рандомно и проверить, если все так)

func mergeAlternately(word1 string, word2 string) string {
	sum := make([]byte, 0, len(word1)+len(word2))

	lesser, larger := []byte(word1), []byte(word2)
	if len(lesser) > len(larger) {
		lesser, larger = larger, lesser
	}

	// перемешиваем
	var i int
	for i = range lesser {
		sum = append(sum, word1[i], word2[i])
	}

	// добавляем остатки
	if i != len(larger)-1 {
		sum = append(sum, larger[i+1:]...)
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
