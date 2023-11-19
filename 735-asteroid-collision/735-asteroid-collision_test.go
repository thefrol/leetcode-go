package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/asteroid-collision

// https://github.com/thefrol/leetcode-go

// Тут суть такая, есть массив чисел. Если число больше нуля
// то это масса астероида летящего вправо, а меньше нуля
// налево. При столкновении самый толстный астероид разбивает
// мелкий и летит дальше. Нужно посчитать что останется в конец
//
// 3,1,-2 станет 3. Так как -2 разобьет 1, а 3 разобьет -2
// -2,3 так и останется
//
// И ещё один важный момент, если они одинаковые по массе, то уничтожаются оба

// Решать буду стеком, как собсно завещано. Стеком решаеются задачки,
// когда вдруг надо вернуться на шаг назад.

func asteroidCollision(asteroids []int) []int {
	st := Stack{
		items: make([]int, len(asteroids)),
	}

	for _, v := range asteroids {
		// достаем астероиды по одному
		if v > 0 {
			st.Push(v)
		}

		// если отрицательный, залезаем в стек и по одному там всех убиваем, пока не встретим кого-то сильнее его
		for {
			p, ok := st.Pop()
			if !ok {
				// отрицательный астероид дошел до начала
				st.Push(v)
				break
			}
			// елси астероид отрицальный в стеке, то значит дальше нет положительных больше, эти все улетят налево
			if p < 0 {
				st.Push(p) // возвращаем что достали
				st.Push(v)
				break
			}

			// какой-то положительный астероид взорвал отрицательный
			if p > -v { // больше по модулю

				st.Push(p)
				break
			}
			// если они равны по массе, то взрываются оба, в стек не пишем
			if p == -v {

				break
			}
		}

	}

	return st.Get()
}

type Stack struct {
	items  []int
	cursor int32
}

func (st *Stack) Pop() (res int, found bool) {
	// в рамках некущей задачи возврашщать не надо
	if st.cursor == 0 {
		return 0, false
	}
	st.cursor--
	return st.items[st.cursor], true
}

func (st *Stack) Push(c int) {
	st.items[st.cursor] = c
	st.cursor++
}

func (st Stack) Get() []int {
	return st.items[:st.cursor]
}

func Test_asteroidCollision(t *testing.T) {
	tests := []struct {
		asteroids, want []int
	}{
		{[]int{3, 2, 1}, []int{3, 2, 1}},
		{[]int{3, -2, 1}, []int{3, 1}},
		{[]int{3, 1, -2}, []int{3}},
		{[]int{3, -2, 1, -4}, []int{-4}},
		{[]int{-1, 2}, []int{-1, 2}},

		//leetcode
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%+v", tt.asteroids)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, asteroidCollision(tt.asteroids))
		})
	}
}
