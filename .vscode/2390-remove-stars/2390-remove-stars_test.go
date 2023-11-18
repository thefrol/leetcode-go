package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/removing-stars-from-a-string

// https://github.com/thefrol/leetcode-go

// дана строка, надо удалить все звездочки. Звезда
// удаляется сразу вместе с прошлым незвездным элементом
// ab*c -> ac

// Ну стек делаем самопальный, причем на той же строке что нам дана)

func removeStars(s string) string {
	bb := []byte(s)
	st := Stack{
		s: bb,
	}

	for _, c := range bb {
		switch c {
		case '*':
			st.Pop()
		default:
			st.Push(c)
		}
	}

	return string(st.Get())
}

type Stack struct {
	s      []byte
	cursor int32
}

func (st *Stack) Pop() {
	// в рамках некущей задачи возврашщать не надо
	st.cursor--
}

func (st *Stack) Push(c byte) {
	st.s[st.cursor] = c
	st.cursor++
}

func (st Stack) Get() []byte {
	return st.s[:st.cursor]
}

func Test_removeStars(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"", ""},
		{"a", "a"},
		{"a*", ""},
		{"ab*c", "ac"},
		//leetcode
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.want, removeStars(tt.s))
		})
	}
}
