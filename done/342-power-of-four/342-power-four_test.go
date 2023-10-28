package main

import (
	"math/bits"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func isPowerOfFour(n int) bool {
// 	// n=fraction*e^pow fraction in [0.5,1)
// 	// берем fraction=0,5 но это означает что
// 	// должны быть нечетные степени, например
// 	// 4=0.5*8=0.5*2^3 fraction=0,5 pow=3
// 	fraction, pow := math.Frexp(float64(n))
// 	return pow%2 == 1 && fraction == 0.5
// }

func isPowerOfFour(n int) bool {
	trail := bits.TrailingZeros32(uint32(n))
	return n == 1 || n != 0 && trail%2 == 0 && n>>bits.TrailingZeros32(uint32(n)) == 1
}

func TestPower(t *testing.T) {
	testCases := []struct {
		val int
		res bool
	}{
		{val: 0, res: false},
		{val: 1, res: true},
		{val: 2, res: false},
		{val: 4, res: true},
		{val: 5, res: false},
		{val: 8, res: false},
		{val: 16, res: true},
		{val: 256, res: true},
	}
	for _, tC := range testCases {
		t.Run("val:"+strconv.Itoa(tC.val), func(t *testing.T) {
			assert.Equal(t, tC.res, isPowerOfFour(tC.val))
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
			isPowerOfFour(v)
		}
	}
}
