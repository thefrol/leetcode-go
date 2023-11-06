package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/increasing-triplet-subsequence/

// https://github.com/thefrol/leetcode-go

// тут нужно найти в массиве чисел подпоследовательность номеров
// i<j<k : num(i)<num(j)<num(k)

// думаю мы просто идём по числам.
// 1. ищем максимум и минимум
// 2. если нашли - сразу возвращаем
// 3. надо ещё третье число

// см тесты, все не так просто

// как насчет работать по методике разделяй и властвуй. Как быстрый поиск.
// выбераем среднее числа. И ищем слева и справа.

// возможно для начала массив стоит отсортировать, чтобы брать числа у которых
// вообще чисто теоретически возможны  такие значения

// а может в процессе сортировки и как бы что-то да обнаружится лол

// если массив очень большой, то предварительная сортировка все оч замедлит

// мы сделаем как бы частный случай сортировки быстрым поиском.

func increasingTriplet(nums []int) bool {
	// попробую массивы не создавать
	// только массив проверенных чисел

	// ещё можем сохранять какой-то массив
	// есть ли у нас числа больше такого-то

	// тут будут граничные кейсы

	//быстрый выходnums
	center, last := len(nums)/2, len(nums)-1
	if nums[0] < nums[center] && nums[center] < nums[last] {
		return true
	}

	for i := 0; i < center; i++ {
		if find(nums, center-i) || find(nums, center+i) {
			return true
		}
	}

	return false

}

// find вернет true елси есть числа меньшие по индексу center и меньшие по значению
// и больше по индекусу и по значению
func find(nums []int, center int) bool {
	// мне не нравится, что у меня не сам алгоритм, а как бы его практическое применение
	// от этого сложнее становится
	if center >= len(nums)-1 && center == 0 {
		return false
	}
	var i int
	for i = 0; i <= center && nums[i] >= nums[center]; i++ {
	}

	if i >= center {
		return false
	}
	for i = center + 1; i < len(nums) && nums[i] <= nums[center]; i++ {
	}

	return i != len(nums) // если уже за последним элементом, то ошибка
}

func Test(t *testing.T) {
	testCases := []struct {
		ints []int
		want bool
	}{
		{[]int{12, 3}, false},
		{[]int{3, 12}, false},
		{[]int{1, 3, 4}, true},
		{[]int{1, 1, 1, 4}, false},
		{[]int{1, 5, 4, 4}, false}, // да уж, k не так-то просто искать будет

		// // i не обязательно всегда минимум глобальный.
		{[]int{6, 7, 1, 9}, true}, // даа жесть

		// но для быстрой работы можно поискать минимум
		// может как-то отсортировать числа даже
		// и постепенно отваливаться к более общему решению

		// но сначала надонаписать общее решение
	}
	for _, tC := range testCases {
		name := fmt.Sprint(tC.ints)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tC.want, increasingTriplet(tC.ints))
		})
	}
}
