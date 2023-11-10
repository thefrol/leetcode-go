package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length

// https://github.com/thefrol/leetcode-go

// 98 процентов на скорости

// дана строка  и длинна подстроки и надо найти максимальное
// количество гласных в подстроке заданной длинны

// ну опять же скользящее окно, идет с первого до последнего элемента
// смотрим сколько гласных прибыло, а сколько выбыло

// оптимизация тоже будет, имбо мы можем запоминать была ли там именно гласная
// чтобе не проверять опять буквы из исходного массива. МОжно опять кольцевым
// буфером воспользоваться

// теперь самопальным, основанным на массиве)

// v2 оптимизация самопальной isVowel, чтобы int выдавала

func maxVowels(s string, k int) int {
	sum, maxsum := 0, 0
	r := NewRing(k)

	for i := 0; i < len(s); i++ {
		vowels := isVowel(s[i])
		sum += vowels - r.Push(vowels)
		if sum > maxsum {
			maxsum = sum
		}

	}
	return maxsum
}
func isVowel(v byte) int {
	if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
		return 1
	}
	return 0
}

type ring struct {
	curr int
	vals []int
}

func NewRing(n int) ring {
	return ring{
		vals: make([]int, n),
	}
}

// Push кладет новое значение и достает то,
// что было в очереди
func (r *ring) Push(x int) int {
	x, r.vals[r.curr] = r.vals[r.curr], x
	r.curr = (r.curr + 1) % len(r.vals)
	return x
}

func Test_maxVowels(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{s: "abciiidef", k: 3, want: 3},
		{s: "aeiou", k: 2, want: 2},
		{s: "leetcode", k: 3, want: 2},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%d@%s", tt.k, tt.s)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxVowels(tt.s, tt.k))
		})
	}
}

func TestPush(t *testing.T) {
	r := NewRing(3)
	r.Push(1)
	r.Push(2)
	r.Push(3)
	assert.Equal(t, 1, r.Push(-1))
	assert.Equal(t, 2, r.Push(-1))
	assert.Equal(t, 3, r.Push(-1))
	assert.Equal(t, -1, r.Push(-1))
}

// блин какое красивое решение
