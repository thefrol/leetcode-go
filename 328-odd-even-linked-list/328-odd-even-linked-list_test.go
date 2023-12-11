package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/odd-even-linked-list

// https://github.com/thefrol/leetcode-go

// дан связный список, нужно взять нечетные элементы
// и положить их после четных элементов, и вернуть.
// Ограничение по памяти О(1) по скорости О(н)

// думаю просто собрать два списка по циклу и
// потом один к другому прикрепить, не забыв
// закрыть последний элемент

// нужно ещё посмотреть что быстрее работает
// i%2 или i&1

// по идее самое быстрое, это сразу по два элемента обрабатывать

// V2 можно например odd.Next=odd.NExt.Next
// и можно отсчет делать от even, тогда в конце не надо
// будет дополнительные проверки делать
func oddEvenList(head *ListNode) *ListNode {
	oddHead := &ListNode{}
	evenHead := &ListNode{}

	o, e := oddHead, evenHead
	for head != nil && head.Next != nil {
		o.Next = head
		o = o.Next
		head = head.Next

		e.Next = head
		e = e.Next
		head = head.Next
	}

	if head != nil {
		o.Next = head
		o = o.Next
	}

	o.Next = evenHead.Next
	e.Next = nil

	return oddHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test(t *testing.T) {
	testCases := []struct {
		arr, want []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 3, 2, 4},
		},

		//leetcode

		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 3, 5, 2, 4},
		},

		{
			[]int{2, 1, 3, 5, 6, 4, 7},
			[]int{2, 3, 6, 7, 1, 5, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.arr), func(t *testing.T) {
			got := oddEvenList(ConvertArray(tC.arr))
			assert.Equal(t, tC.want, ConvertList(got))
		})
	}
}

// helpers

func ConvertArray(arr []int) *ListNode {
	dummy := ListNode{}
	i := &dummy
	for _, v := range arr {
		i.Next = &ListNode{
			Val: v,
		}
		i = i.Next
	}
	return dummy.Next
}

func Test_Converting(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	assert.Equal(t, arr, ConvertList(ConvertArray(arr)))
}

func ConvertList(head *ListNode) []int {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}

// tests and benchs

func Test_Binary(t *testing.T) {
	for i := 0; i < 1000; i++ {
		assert.Equal(t, 3%2, 3&1)
	}
}

func Benchmark(b *testing.B) {
	b.Run("with &", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = i & 1
		}
	})

	b.Run("with %", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = i % 2
		}
	})

}
