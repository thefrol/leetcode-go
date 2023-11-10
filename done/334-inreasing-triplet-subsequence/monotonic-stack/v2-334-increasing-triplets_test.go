package main

// ПОПЫТКА МОНОТОННОГО СТЕКА НЕ УДАЛАСЬ

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/increasing-triplet-subsequence/

// https://github.com/thefrol/leetcode-go

// // тут нужно найти в массиве чисел подпоследовательность номеров
// i<j<k : num(i)<num(j)<num(k)

// а не воспользоваться ли нам монотонным стеком
// для каждого числа, если оно больше, чем
// чем верхушка - добавляем новое. Если меньше,
// то выкидываем все значения и ищем то,
// где новое значение будет больше того, что в стеке

// то есть если мы нашли минимум, но оно полностью
// обновит стек. Если максимум, то достроит
// надо просто глубину проверять

func increasingTriplet(nums []int) bool {
	s := NewMonotonicRising(2) // это монотонный возрастающий стек
	//когда положим три элемента будет переполнение

	var ok bool
	for _, v := range nums {
		ok = s.Push(v) // overflow:=!ok
		if !ok {
			// если стек переполнился
			// значит нашли третий элемент
			return true
		}
	}
	return false
}

// монотонный стек

type MonotonicRising struct {
	s Stack
}

func NewMonotonicRising(n int) MonotonicRising {
	return MonotonicRising{
		s: NewStack(n),
	}
}

func (s *MonotonicRising) Push(x int) bool {
	for {
		if v, ok := s.s.Top(); ok && v > x {
			s.s.Pop()
		} else {
			break
		}

	}
	return s.s.Push(x)
}

func (s MonotonicRising) Size() int {
	return s.s.Size()
}

func TestMonotonic(t *testing.T) {
	s := NewMonotonicRising(3)
	s.Push(3)
	s.Push(0)

	assert.Equal(t, 1, s.Size())

	s.Push(2)
	s.Push(1)

	assert.Equal(t, 1, s.Size())

}

// обычный стек

type Stack struct {
	arr  []int
	curr int
}

func NewStack(n int) Stack {
	return Stack{
		arr: make([]int, n),
	}
}

func (s *Stack) Pop() (val int, found bool) {
	if s.curr > 0 {
		val, s.curr = s.arr[s.curr-1], s.curr-1
		found = true
	}
	return
}

func (s *Stack) Push(x int) (ok bool) {
	if len(s.arr) == s.curr {
		return false // overflow
	}
	s.arr[s.curr] = x
	s.curr++
	return true
}

func (s Stack) Size() int {
	return s.curr
}

func (s Stack) Top() (val int, found bool) {
	if s.curr == 0 {
		return 0, false
	}
	return s.arr[s.curr-1], true
}

func TestStack(t *testing.T) {
	s := NewStack(3)
	ok := s.Push(1)
	assert.True(t, ok)
	tt, ok := s.Top()
	assert.True(t, ok)
	assert.Equal(t, 1, tt)
	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	// more tests
	s.Push(1)
	s.Push(2)
	s.Push(3)
	ok = s.Push(4)
	assert.False(t, ok)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = s.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func Test(t *testing.T) {
	testCases := []struct {
		ints []int
		want bool
	}{
		// {[]int{12, 3}, false},
		// {[]int{3, 12}, false},
		// {[]int{1, 3, 4}, true},
		// {[]int{1, 1, 1, 4}, false},
		// {[]int{1, 5, 4, 4}, false}, // да уж, k не так-то просто искать будет

		// // // i не обязательно всегда минимум глобальный.
		{[]int{6, 7, 1, 9}, true}, // монотонный стек проиграл
		{[]int{1, 5, 0, 4, 1, 3}, true},
		{[]int{6, 7, 1, 2}, false},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 7}, true},

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
