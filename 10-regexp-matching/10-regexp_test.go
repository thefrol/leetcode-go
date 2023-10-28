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

func Benchmark(b *testing.B) {
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
		//{pattern: "miss*s.s.sasdlnaweiwen*.**", s: "missis.sdfisdfalksd**", res: false},
		/// ага! значит не работают точки
		{pattern: "...", s: "abc", res: true}, /// чиню
		{pattern: "..c", s: "abc", res: true}, // не работает точка в конце
		// неправильное условие стояло
		//
		// господи, спасибо тестам
	}
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			isMatch(v.s, v.pattern)
		}

	}
}

// без кеша
//
//cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
//Benchmark-4   	  229239	      5371 ns/op	    1616 B/op	      64 allocs/op

// сделал проверку продстрок
// cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
// Benchmark-4   	  200332	      5660 ns/op	    1200 B/op	      51 allocs/op
