package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// // https://leetcode.com/problems/move-zeroes
// //
// // https://github.com/thefrol/leetcode-go
// //
// // будем переносить нули змейкой
// // то есть первое число, где есть ноль меняется местами с первым числом, где нуля нет

// // вообще, я реализовал не совсем то, что имел в виду, у меня два указателя, где
// // второч начинает отсчет постоянно от первого, это тупо. Попробуем, сделать его независящим.
// // а вообще вместо второго указателя, лучше иметь количество нулей

// // Даже так, сохраняем первый ноль в ZeroPtr, а когда меняем местами значения это
// // этот товарищ увеличивается на один

// // V3 избавляемся от указателя, его заменит range

// // V4 сначала найдем первый ноль. Это сократит количество проверок в цикле

func moveZeroes(nums []int) {

	// сначала найдем первый ноль
	zeroPtr := 0
	for zeroPtr = range nums {
		if nums[zeroPtr] == 0 {
			break
		}
	}

	// нет нулей в строке или он в конце
	if zeroPtr >= len(nums)-1 {
		return
	}

	// начинаем искать со следующего после нуля
	for ptr := zeroPtr + 1; ptr < len(nums); ptr++ {
		//ищем ноль.
		if nums[ptr] != 0 {
			nums[zeroPtr], nums[ptr] = nums[ptr], nums[zeroPtr]
			zeroPtr++
		}
	}

	// соберем мусор
	// runtime.GC() - победить по оперативке, проиграть по скорости

}

// V2 моя элегантная версия
// func moveZeroes(nums []int) {
// 	zeroPtr := -1
// 	ptr := 0
// 	for ptr = range nums {
// 		//ищем ноль.
// 		if nums[ptr] == 0 && zeroPtr == -1 {
// 			// впервые наткнулись на ноль, сохраняем в память
// 			// далее нам в целом это не интересно больше
// 			// это можно сделать в процессе подготовки, поиск первого нуля
// 			zeroPtr = ptr
// 		} else if zeroPtr != -1 && nums[ptr] != 0 {
// 			nums[zeroPtr], nums[ptr] = nums[ptr], nums[zeroPtr]
// 			zeroPtr++
// 		}
// 	}

// нашел в инете
// фича в том, что первый свап произойдет 0 и 0, поэтому
// nextIndex переместится дальше, и будет всегда впереди на один.
// Гениально
// func moveZeroes(nums []int) {
// 	nextIndex := 0

//		for i, num := range nums {
//			if num != 0 {
//				nums[i], nums[nextIndex] = nums[nextIndex], nums[i]
//				nextIndex++
//			}
//		}
//	}
func Test(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: nums(1, 0, 2), want: nums(1, 2, 0)},
		{nums: nums(1, 0), want: nums(1, 0)},
		{nums: nums(0), want: nums(0)},
		{nums: nums(0, 1), want: nums(1, 0)},

		{nums: nums(0, 1, 0, 3, 12), want: nums(1, 3, 12, 0, 0)},

		{nums: nums(1, 2, 3, 4), want: nums(1, 2, 3, 4)}, // строка без нулей
		{nums: nums(1, 2, 3, 4, 0), want: nums(1, 2, 3, 4, 0)},
		{nums: nums(1, 2, 3, 0, 4), want: nums(1, 2, 3, 4, 0)}, // быстрый сброс в конец
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
