package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/unique-number-of-occurrences/

// https://github.com/thefrol/leetcode-go

// Дан массив чисел, нужно вернуть true, если каждое число
// появляется уникальное число раз

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int) // тут будет хранится количество появлений

	for _, v := range arr {
		m[v]++
	}

	occurences := make(map[int]bool, len(m))

	for _, v := range m {
		if _, found := occurences[v]; found { // если такое число появлений уже есть, то провал
			return false
		}
		occurences[v] = true // запоминаем, что такое количество раз уже появлялось
	}
	return true
}

func Test_uniqueOccurrences(t *testing.T) {

	tests := []struct {
		ints []int
		want bool
	}{
		{[]int{1, 1, 1}, true},
		{[]int{1, 1, 2}, true},
		{[]int{1, 1, 2, 2}, false},
		//lettcode lower

	}
	for _, tt := range tests {
		name := fmt.Sprintf("%+v", tt.ints)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, uniqueOccurrences(tt.ints))
		})
	}
}
