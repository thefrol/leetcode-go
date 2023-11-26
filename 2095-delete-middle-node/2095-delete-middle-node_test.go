package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list

// github.com/thefrol/leetcode-go

// дан связный список. Нужно удалить средний элемент,
// типа с индексом middle(list)=floor(len(list)/2)
//
// и приколь в том, что элементы считаются с нулевого.
// то есть в спике из двух элементов, средним будет элемент
// с индексом 1 - последний
//
// 4   3   ~2~  1
// i=0 i=1 i=2 i=3     len=4, middle=2
//
// ДУмаю решение такое. У меня будут два указателя. Один бежит до конца, а
// второй перехоит к следующему элементу каждую вторую итерацию. Таким
// образом он будет указывать на средний элемент списка средка

func deleteMiddle(head *ListNode) *ListNode {
	curr := head
	len := 0

	if head.Next == nil {
		return nil
	}

	// самая большая сложность: чтобы удалить средний элемент,
	// нам надо помнить элемент перед ним
	middle := head
	prev := head // элемент перед средним, мы могли бы сэкономить одну переменную,
	// но тогда как-то менее очевиден стал бы алгоритм, что мы начинаем с 3 элемента
	// или сделать очень большое "встепление к алгоритму в три итерации"
	for curr != nil {
		len++ //
		if len%2 == 0 {
			// каждую вторую итерацию двигаем средний элемент
			prev = middle
			middle = middle.Next
		}
		curr = curr.Next
	}

	//удаляем средний элемент
	prev.Next = prev.Next.Next

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test(t *testing.T) {
	testCases := []struct {
		ints, want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 4, 5}},
		{[]int{1, 2}, []int{1}},
		{[]int{1}, nil},
	}
	for _, tC := range testCases {
		name := fmt.Sprint(tC.ints)
		t.Run(name, func(t *testing.T) {
			l := toList(tC.ints)
			got := deleteMiddle(l)
			assert.Equal(t, tC.want, fromList(got))

		})
	}
}

func toList(arr []int) *ListNode {
	curr := &ListNode{}
	first := curr
	for _, v := range arr {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = v
	}
	return first.Next
}

func fromList(head *ListNode) (arr []int) {
	if head == nil {
		return nil
	}
	arr = append(arr, head.Val)
	for head.Next != nil {
		head = head.Next
		arr = append(arr, head.Val)
	}
	return
}

func TestConvert(t *testing.T) {
	orig := []int{1, 2, 3, 3, 4, 5, 56}
	got := fromList(toList(orig))
	assert.Equal(t, orig, got)
}
