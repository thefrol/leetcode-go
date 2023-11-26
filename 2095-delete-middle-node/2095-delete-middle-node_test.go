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

// V2 подсмотрел прикольное решение, где мы всегда прыгаем на два шага

func deleteMiddle(head *ListNode) *ListNode {
	// очень часто испльзуется этот трюк с минус первым элементом
	// так уж тут сложены циклы, что нам как бы надо начинать их
	// с элемента перед нулевым, чтобы они были красивыми, и не надо
	// было слишком сильно обрабатывать начальные условия
	dummy := &ListNode{0, head}

	// самая большая сложность: чтобы удалить средний элемент,
	// нам надо помнить элемент перед ним
	prev, curr := dummy, head // один указывает на дамми, а курсор на хед
	// то есть прев изначально указывает на -1 элемент
	for curr != nil && curr.Next != nil {
		prev = prev.Next
		curr = curr.Next.Next
	}

	//удаляем средний элемент
	prev.Next = prev.Next.Next

	return dummy.Next // мы могли удалить нулевой элемент через prev
	// поэтому мы не можем вернуть head и пользуемся триксом через
	// -1 элемент дамми
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
