package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/is-subsequence

// https://github.com/thefrol/leetcode-go

// надо найти подстроку в строке
// face - ace = true
// aciole -ace = true

// думаю последовательно иду по строке и ищу буквы

// просто ищу первое попадание буквы из списка. Как только нахожу переключасб
// когда буквы в s кончатся, это победа, если цикл дошел до конца - поражение
func isSubsequence(s string, t string) bool {
	i := -1         // при создании первого цикла будет сразу инкремент
	tb := []byte(t) // просто чтобы бегать по байтам
loop:
	for _, c := range []byte(s) {
		for i++; i < len(tb); i++ {
			if tb[i] == c {
				continue loop
			}
		}
		// условие остановки,
		// тут мы окажемся только если цикл подошел к концу,
		// но буква не нашлась
		return false
	}
	// закончились буквы в массике
	return true
}

func Test(t *testing.T) {
	testCases := []struct {
		s, t string
		want bool
	}{
		//{"ace", "face", true},
		{"aaaaaa", "bbaaaaa", false},
	}
	for _, tC := range testCases {
		t.Run(tC.s+"<"+tC.t, func(t *testing.T) {
			assert.Equal(t, tC.want, isSubsequence(tC.s, tC.t))
		})
	}
}
