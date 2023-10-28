package main

import (
	"math/bits"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
// Benchmark-4   	 1219058	      1006 ns/op	       0 B/op	       0 allocs/op
// PASS
// func isPowerOfTwo(n int) bool {
// 	// важен порядок операторов сравнения
// 	if n == 0 {
// 		return false
// 	}
// 	for ; n&1 == 0; n = n >> 1 {
// 	}
// 	return n == 1

// }

//cpu: Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz
//Benchmark-4   	 6587179	       180.3 ns/op	       0 B/op	       0 allocs/op
// func isPowerOfTwo(n int) bool {
// 	return n>>bits.TrailingZeros64(uint64(n)) == 1

// }

func isPowerOfTwo(n int) bool {
	return n>>bits.TrailingZeros32(uint32(n)) == 1
}

func TestPower(t *testing.T) {
	testCases := []struct {
		val int
		res bool
	}{
		{val: 0, res: false},
		{val: 1, res: true},
		{val: 2, res: true},
		{val: 3, res: false},
		{val: 5, res: false},
		{val: 8, res: true},
		{val: 15, res: false},
		{val: 256, res: true},
	}
	for _, tC := range testCases {
		t.Run("val:"+strconv.Itoa(tC.val), func(t *testing.T) {
			assert.Equal(t, tC.res, isPowerOfTwo(tC.val))
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
			isPowerOfTwo(v)
		}
	}
}
