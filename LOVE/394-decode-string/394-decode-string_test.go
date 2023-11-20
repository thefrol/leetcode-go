package main

import (
	"bytes"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/decode-string

// https://github.com/thefrol/leetcode-go

// задача состоит в том, чтобы декодировать строку. Она как бы
// сжата вот так:
// 2[a]=aa или 3[abc]=abcabcabc
// причем, если всмотреться в примеры, то можно заметить, что
// декодирование может быть вложенное
// 2[a2[b]]=abbabb

// поначалу мне казалось, что можно сделать что-то типа врайтера,
// с буферами для последнего числа и паттерна, но сложенность
// дает свои плоды. Тут появляется что-то вроде рекурсии.

// А что хорошо выводит рекурсию и не оставеяет следа на строках?
// правильно: стек!

// буду пушить в стек, пока не встретится "]", после чего мы достаем всю
// последовательность типа "(int)[(str)" из стека и дополняем стек
// идем дальше

// по поим расчетам это позволит обойтись без рекурсии, так как внутренний
// паттерн будет уже деподирован к тому времени, как мы начнем декодировать
// внешний

func decodeString(s string) string {
	bb := []byte(s)
	st := Stack{}

	for _, v := range bb {
		if v == ']' {
			// декодируем, и добавляем в стек
			s, n := PopPattern(&st)
			decoded := bytes.Repeat(s, n)
			st.Push(decoded...)
			continue // чтобы "]" не записалась
		}
		st.Push(v)
	}
	return st.GetString()
}

func PopPattern(st *Stack) ([]byte, int) {
	s := []byte{} // это что мы будем умножать
	n := 0        // сколько раз

	// сначала достаем буквы до появления "["
	for st.Len() > 0 { // тут можно не проверять, строка всегду будет содержать [']
		c := st.Pop()
		if c == '[' {
			break
		}
		s = append(s, c)
	}

	// теперь достаем цифры
	// букву можно просто преобразовать в цифры вычитанием
	// '1'-'0'=1
	// см. Test_SubtractChar()
	base := 1 // разряд 1,10,100 итд
	for st.Len() > 0 {
		c := st.Pop()
		if c < '0' || c > '9' {
			// если не число- выходим
			st.Push(c) // но сначала возвращаем обратно
			break
		}

		n += int(c-'0') * base // int('c'-0) это число
		base *= 10
	}

	// обязательно развернуть
	slices.Reverse(s)

	return s, n
}

type Stack struct {
	arr []byte
}

func (s *Stack) Push(c ...byte) {
	s.arr = append(s.arr, c...)
}

func (s *Stack) Pop() (res byte) {
	res = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return
}

func (s Stack) Len() int {
	return len(s.arr)
}

func (s Stack) GetString() string {
	return string(s.arr)
}

func Test_SubtractChar(t *testing.T) {
	assert.Equal(t, int32(1), '1'-'0')
	assert.Equal(t, int32(9), '9'-'0')
}

func Test_decodeString(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"a", "a"},
		{"ab", "ab"},
		{"2[a]", "aa"},
		{"a2[b]", "abb"},
		{"a2[b2[c]]", "abccbcc"},
		{"", ""},

		//leetcode
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			assert.EqualValues(t, tt.want, decodeString(tt.s))
		})
	}
}

func TestPopPattern(t *testing.T) {

	tests := []struct {
		stack string
		s     string
		n     int
	}{
		{"abc23[aa", "aa", 23},
		{"2[abc", "abc", 2},
	}
	for _, tt := range tests {
		t.Run(tt.stack, func(t *testing.T) {
			st := Stack{
				arr: []byte(tt.stack),
			}
			gotS, gotN := PopPattern(&st)

			assert.Equal(t, tt.s, string(gotS))
			assert.Equal(t, tt.n, gotN)
		})
	}
}
