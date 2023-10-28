package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		pattern string
		s       string
		res     bool
	}{
		{pattern: "12313.*11", s: "12313____11", res: true},

		{pattern: "aa", s: "a", res: false},
		{pattern: "a*", s: "aa", res: true},
		{pattern: "mis*is*ip*.", s: "mississippi", res: true},
		// я придумал разбить миссисипи на подзадачки словно бы я тоже машина
		{pattern: "mis*", s: "miss", res: true},
		{pattern: "mis*", s: "missi", res: false},
		{pattern: "mis*is*", s: "mississ", res: true},
		{pattern: "mis*is*ip", s: "mississip", res: true},
		{pattern: "mis*is*ip*", s: "mississipp", res: true},
		/// ага! значит не работают точки
		{pattern: "...", s: "abc", res: true}, /// чиню
		{pattern: "..c", s: "abc", res: true}, // не работает точка в конце
		// неправильное условие стояло
		//
		// господи, спасибо тестам
	}
	for _, tC := range testCases {
		t.Run(tC.pattern+" "+tC.s, func(t *testing.T) {
			assert.Equal(t, tC.res, isMatch(tC.s, tC.pattern))
		})
	}
}
