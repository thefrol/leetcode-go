package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-consecutive-ones-iii/

// https://github.com/thefrol/leetcode-go

// дан массив, и число k - сколько нулей можно превратить в единички
// нужно найти самую длинную последовательность единиц, которую можн
// получить, если k ноликов превратить в единички

// действуем скользящим окном. Оно только разной ширины будет. Снача
// наполняем окно, так чтобы в нем было два нуля, а потом:
// - если пришешла единичка, удлинняем окно направо
// - если пришел нолик, сокращаем слева до ближайшего нолика
// таким образом в окне всегда будет два нолика.

// То есть по сути мы
// мы считаем количество нулей в окне переменной длинны. МОжно было бы
// так это решать

// jump0 и pop0 можно объединить

func longestOnes(nums []int, k int) int {
	l, r := 0, 0

	// pop0 скользит левой границей в поисках нуля
	pop0 := func() {
		for ; l < len(nums) && nums[l] == 1; l++ {
		}
		l++
	}

	// push - если справа приходит ноль, но раздвигает направо
	// если нет, то сужает слева в поисках нуля
	push := func() {
		if nums[r] == 0 {
			pop0()
		}
		r++
	}

	// находит ближайший нолик справа и перескакивает через него
	jump0 := func() {
		for ; r < len(nums) && nums[r] == 1; r++ {
		}
		r++
	}

	// двигаем указатель справа, пока не добавим
	// два нуля
	for i := 0; i < k; i++ {
		jump0()
	}

	// основной цикл
	if r >= len(nums) {
		return len(nums)
	}

	max := r - l
	// теперь действуем по алгоритму с добавлением и удалением
	for r < len(nums) {
		push()
		if max < r-l {
			max = r - l
		}
	}
	return max
}

func Test_longestOnes(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{0, 0, 0}, 2, 2},
		{[]int{1, 1, 1}, 2, 3},
		{[]int{1, 1, 1, 0}, 2, 4},
		{[]int{1, 1, 1, 0, 0}, 2, 5},
		// leetcode
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%d@%v", tt.k, tt.nums)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestOnes(tt.nums, tt.k))
		})
	}
}
