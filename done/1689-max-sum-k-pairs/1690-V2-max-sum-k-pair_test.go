package main_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://github.com/thefrol/leetcode-go

// ВОобще задача из списка задач на два указателя.
// Реально можно так сделать. Стоит обратить внимание
// что нужно проверять текущее число на сумму с самим собой
// можно отсортировать массив и проверять, например, только большие числа
// один указатель идет с начала - второй с конца,
// и сбрасывается когда доходит до нужного число

func maxOperations(nums []int, k int) int {
	slices.Sort(nums)
	counter := 0

	var l, r = 0, len(nums) - 1
	pair := -1
	for {
		pair = k - nums[l]
		for ; r > 0 && nums[r] > pair; r-- {
		}

		if l >= r {
			break
		}

		if nums[r] == pair {
			// пара нашлась
			counter++
			r--
		}

		l++

	}
	return counter
}

func Test_maxOperationsV2(t *testing.T) {

	tests := []struct {
		nums     []int
		k        int
		want     int
		wantNums []int
	}{
		{
			nums: []int{1, 2, 3, 4}, k: 5,
			wantNums: []int{}, want: 2,
		},
		{
			nums: []int{3, 1, 3, 4, 3}, k: 6,
			wantNums: []int{1, 4, 3}, want: 1,
		},
		{
			nums: []int{3, 1, 5, 1, 1, 1, 1, 1, 2, 2, 3, 2, 2}, k: 1,
			wantNums: []int{1, 4, 3}, want: 0, // на литкоде array out of bounds
		},
		{
			nums: []int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, k: 2,
			wantNums: []int{1, 4, 3}, want: 2, // на литкоде array out of bounds
		},
		{
			nums: []int{3, 5, 1, 5}, k: 2,
			wantNums: []int{}, want: 0,
		},

		{
			nums: []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, k: 3,
			wantNums: []int{}, want: 4,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("k=%d; [%v]", tt.k, tt.nums)
		t.Run(name, func(t *testing.T) {
			got := maxOperations(tt.nums, tt.k)
			assert.Equal(t, tt.want, got, "количество операция нужно другое")
		})
	}
}
