package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-words-in-a-string/

// https://github.com/thefrol/leetcode

// сделаем через встроенные библиотеки, и через указатели, пойдем с конца просто

// если делать двумя массивами относительно просто)

// если попытаться в одни масик простой перестановкой слов короче будет непросто обойтись

// поэтому идем с конца и ищем ссылки пробелы

// func reverseWords(s string) string {

// 	var ss []string
// 	end := -1
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] != ' ' && end == -1 {
// 			end = i + 1
// 		} else if s[i] == ' ' && end != -1 {
// 			ss = append(ss, s[i+1:end])
// 			end = -1
// 		}
// 	}
// 	if end != -1 {
// 		ss = append(ss, s[:end])
// 	}
// 	return strings.Join(ss, " ")
// }

// V2 очень легко через fields
func reverseWords(s string) string {
	ss := strings.Fields(s)

	i, j := 0, len(ss)-1
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--

	}

	return strings.Join(ss, " ")
}

func Test(t *testing.T) {
	testCases := []struct {
		s, want string
	}{
		{"  the sky   is blue", "blue is sky the"},
		{"", ""},
		{"s  ", "s"},
	}
	for _, tC := range testCases {
		t.Run(tC.s, func(t *testing.T) {
			assert.Equal(t, tC.want, reverseWords(tC.s))
		})
	}
}
