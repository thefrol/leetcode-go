package main

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-two-numbers/
//

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://github.com/thefrol/leetcode-go

func Append(n *ListNode, val int) *ListNode {
	new := ListNode{
		Val: val,
	}

	if n == nil {
		return &new
	}

	n.Next = &new
	return &new
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r, root *ListNode

	var v, leftover uint
	for {
		// условие выхода
		if l1 == nil && l2 == nil {
			if leftover != 0 {
				Append(r, int(leftover))
			}
			break
		}

		sum := leftover

		r = Append(r, 0)
		// на случай первого прохода
		if root == nil {
			root = r
		}
		if l2 != nil {
			sum += uint(l2.Val)
			l2 = l2.Next
		}
		if l1 != nil {
			sum += uint(l1.Val)
			l1 = l1.Next
		}

		// тут очень забавный момент, что наше значение это по сути остаток от деления, можно легко перепутать
		leftover, v = bits.Div(0, sum, 10) // может тут можно ускорить через 32 разрядность, а может как-то можно использовать и два разряда верхний и нижний
		r.Val = int(v)

	}
	return root
}

// For Tests
func (n *ListNode) Int() int {
	var res int
	for i := 1; n != nil; n, i = n.Next, i*10 {
		res += n.Val * i
	}
	return res
}

func (n *ListNode) Validate() error {
	var v *ListNode
	for v = n; v != nil; v = v.Next {
		if v.Next == nil && v.Val == 0 {
			return fmt.Errorf("trailing zero")
		}
	}
	return nil
}

func NewList(digits ...int) *ListNode {
	var root, r *ListNode
	for _, v := range digits {
		if root == nil {
			root = &ListNode{Val: v}
			r = root
		} else {
			r = Append(r, v)
		}
	}
	return root
}

func Test(t *testing.T) {
	testCases := []struct {
		V1   *ListNode
		V2   *ListNode
		want int
	}{
		// {V1: NewList(1), V2: NewList(2), want: 3},
		// {V1: NewList(1, 2), V2: NewList(2, 2), want: 43},
		// {V1: NewList(7), V2: NewList(9), want: 16},
		{V1: NewList(2, 4, 3), V2: NewList(5, 6, 4), want: 807},
	}
	for _, tC := range testCases {
		name := fmt.Sprintf("%v+%v=%v", tC.V1.Int(), tC.V2.Int(), tC.want)
		t.Run(name, func(t *testing.T) {

			actual := addTwoNumbers(tC.V1, tC.V2)
			assert.NoError(t, actual.Validate())
			assert.Equal(t, tC.want, actual.Int())

		})
	}
}
