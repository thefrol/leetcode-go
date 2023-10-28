package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

//cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
//Benchmark-4   	  715245	      1597 ns/op	       0 B/op	       0 allocs/op
// func isPowerOfThree(n int) bool {
// 	q, r := uint32(n), uint32(0)
// 	for ; q > 0 && r == 0; q, r = bits.Div32(0, q, 3) {
// 	}
// 	return q == 0 && r == 1
// }

//cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
//Benchmark-4   	  721636	      1542 ns/op	       0 B/op	       0 allocs/op
// func isPowerOfThree(n int) bool {
// 	q, r := uint32(n), uint32(0)
// 	for q > 0 && r == 0 {
// 		q, r = bits.Div32(0, q, 3)
// 		if q == 0 && r == 1 {
// 			return true
// 		} else if q == 0 {
// 			return false
// 		}
// 	}
// 	return false
// }

// cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
// Benchmark-4   	   70346	     17262 ns/op	       0 B/op	       0 allocs/op
// func isPowerOfThree(n int) (ret bool) {
// 	var i int
// 	for i = 1; i < n; i *= 3 {
// 	}
// 	return i == n
// }

// cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
// Benchmark-4   	   19503	     60823 ns/op	       0 B/op	       0 allocs/op
// func isPowerOfThree(n int) bool {
// 	return math.Mod(math.Log(float64(n)), math.Log(3)) < 1e-10
// }

// func isPowerOfThree(n int) bool {
// 	if n < 1 {
// 		return false
// 	}
// 	for ; n%3 == 0; n /= 3 {
// 	}
// 	return n == 1
// }

// чужое решение, но я посмотрел разобрался
func isPowerOfThree(n int) bool {
	// 1162261467 - 3^19, оно же самая большая степень тройки,
	// если поделить на наше число недолжно быть остатка, наше число
	// это один из его множителей
	return n > 0 && 1162261467%n == 0
}

func TestPowerOfThree(t *testing.T) {
	tests := []struct {
		val int
		ret bool
	}{
		{val: 0, ret: false},
		{val: 1, ret: true},
		{val: 2, ret: false},
		{val: 7, ret: false},
		{val: 9, ret: true},
		{val: 27, ret: true},
	}

	for _, tc := range tests {
		name := fmt.Sprint(tc.val)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.ret, isPowerOfThree(tc.val))
		})
	}
}

func Benchmark(b *testing.B) {
	var vars []int
	for i := 0; i < 500; i++ {
		vars = append(vars, int(rand.Int63()))
	}

	for i := 0; i < b.N; i++ {
		for _, v := range vars {
			isPowerOfThree(v)
		}
	}
}
