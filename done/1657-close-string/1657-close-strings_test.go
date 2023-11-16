package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/determine-if-two-strings-are-close

// https://github.com/thefrol/leetcode-go

// Долго объяснить условие. Нужно сказать "близки" ли строки. Близость строк -
// когда одна из другой получается перестановкой двух букв или заменой набора букв,
// между собой. И все это неограниченное число раз. Лучше почитай там

// мне кажется что строки могут быть близки если состоят из одинакового набора букв
// и одинакового набора количеств букв аббввв(набор букв: а, б, в; набор количеств 1,2,3)

// итого считаем мапу, из нее собираем наборы букв и количеств. Сравниваем их

// v2: букв то у нас всего 26, это можно в массиве хранить!
// в мапе 114 мс, без 14 мс. Жесть

// Но и памяти массивы больше занимают по итогу

func isClose(o1, o2 [26]int) bool {
	for i, v := range o1 {
		if (v != 0) != (o2[i] != 0) {
			return false
		}
	}
	// sorting!
	sort.Ints(o1[:])
	sort.Ints(o2[:])

	i := 0
	for ; i < 26 && o1[i] == o2[i]; i++ {
	}

	return i == 26
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	w1 := []byte(word1)
	w2 := []byte(word2)

	var o1, o2 [26]int
	for i := 0; i < len(w1); i++ {
		o1[w1[i]-'a']++
		o2[w2[i]-'a']++
	}

	return isClose(o1, o2)
}

func Test_closeStrings(t *testing.T) {
	tests := []struct {
		w1, w2 string
		want   bool
	}{
		{"ab", "ba", true},
		{"aab", "bba", true},
		// leetcode
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s~%s", tt.w1, tt.w2)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, closeStrings(tt.w1, tt.w2))
		})
	}
}
