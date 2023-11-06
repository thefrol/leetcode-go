package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/product-of-array-except-self/

// https://github.com/thefrol/leetcode-go

// красивая задачка, основной смысл, что на выходе должен быть алгоритм
// O(n) без использования оператора деления. это говорит нам о том, что мы
// можем пройти по циклу лишь однажды, или дважды... Но не цикл в цикле

// я думаю можно умножать и идти по списку, так, чтобы в клетку не показало
// умножение на текущее значение. А потом в обрабную строноу клетну умножаем
// на буфер но не умножанаем на текущее значение

// 5=1*2*3*4(первый проход) + 9*8*7*6(второй проход)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	var i int
	res[0] = 1 // кроме первого(нулевого)
	for i = 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	buf := 1
	for i = len(nums) - 2; i >= 0; i-- {
		buf *= nums[i+1]
		res[i] *= buf
		// на последней итерации будет записано, но не используется уже
	}

	return res
}

func Test(t *testing.T) {
	testCases := []struct {
		ints, want []int
	}{
		{[]int{1, 2, 3}, []int{6, 3, 2}},
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tC := range testCases {
		name := fmt.Sprintf("%+v", tC.ints)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tC.want, productExceptSelf(tC.ints))
		})
	}
}
