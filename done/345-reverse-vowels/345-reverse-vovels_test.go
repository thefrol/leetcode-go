package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string/

// https://github.com/thefrol/leetcode-go

// Тут надо переставить местами гласные, из конца в начало
// гласные это 'a','e','i','o','u'

// делаем по алгоритму двух указателей, слева едем и справа,
// главное аккуратно смотреть пограничные случаи, а особенно начальные задачки
// пустая строка, один или два символа

func reverseVowels(s string) string {
	if len(s) < 2 {
		return s
	}

	bb := []byte(s)
	vowels := []byte("aeiouAEIOU")

	last := len(bb) - 1
	left, right := -1, last+1 //начинаем мы как бы из-за границ, чтобы в цикле пропускать сразу последнюю найденную букву

	for left < right {
		// найдем слева
		for left = left + 1; left < right; left++ {
			if slices.Contains(vowels, bb[left]) {
				break
			}
		}
		// найдем справа
		for right = right - 1; right >= left; right-- {
			if slices.Contains(vowels, bb[right]) {
				break
			}
		}

		// если мы не пошли по второму кругу, то...
		if left < right {
			bb[left], bb[right] = bb[right], bb[left]
		}

	}

	return string(bb)
}

func Test(t *testing.T) {
	testCases := []struct {
		s, want string
	}{
		{
			"hello",
			"holle",
		},
		{
			"leetcode",
			"leotcede",
		},
		{"", ""},
		{" ", " "},
		{"Aa", "aA"},
		{".,", ".,"},
	}
	for _, tC := range testCases {
		t.Run(tC.s, func(t *testing.T) {
			assert.Equal(t, tC.want, reverseVowels(tC.s))
		})
	}
}
