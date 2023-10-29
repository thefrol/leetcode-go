package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/move-zeroes
//
// https://github.com/thefrol/leetcode-go
//
// будем переносить нули змейкой
// то есть первое число, где есть ноль меняется местами с первым числом, где нуля нет

func moveZeroes(nums []int) {

	var zero, nonzero int
	for zero = 0; zero < len(nums); zero++ {
		//ищем ноль
		if nums[zero] == 0 {
			for nonzero = zero + 1; nonzero < len(nums); nonzero++ {
				if nums[nonzero] != 0 {
					//меняем местами
					nums[zero], nums[nonzero] = nums[nonzero], nums[zero]
					break // теперь опять начинаем искать ноль
					// метку использую специально, чтобы
				}
			}
			// раннее срабатываение выхода
			if nonzero >= len(nums) {
				return
			}
		}

	}
}

func Test(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		// {nums: nums(1, 0, 2), want: nums(1, 2, 0)},
		// {nums: nums(1, 0), want: nums(1, 0)},
		// {nums: nums(0), want: nums(0)},
		// {nums: nums(0, 1), want: nums(1, 0)},

		{nums: nums(0, 1, 0, 3, 12), want: nums(1, 3, 12, 0, 0)},
	}
	for _, tC := range testCases {
		name := fmt.Sprintf("%+v", tC.nums)
		t.Run(name, func(t *testing.T) {
			moveZeroes(tC.nums)
			assert.EqualValues(t, tC.want, tC.nums)
		})
	}
}

func nums(ints ...int) []int {
	return ints
}
