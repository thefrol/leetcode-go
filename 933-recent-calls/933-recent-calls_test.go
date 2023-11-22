package main

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-recent-calls/

// github.com/thefrol/leetcode-go

// тут короче надо разработать класс, который прям сложно объянсть
// можно добавлять вызов, который прозошел через t мс, при добавлнее
// вызова надо вернуть количество вызова за последние 3000 мс. t это как бы
// абсолютная величина

// думаю, я сделать на container/list тупо очеред с двумя концами, после вызова
// с начала освобождаю самые старые вызовы

const interval = 3000

type RecentCounter struct {
	*list.List
}

func Constructor() RecentCounter {
	return RecentCounter{
		List: list.New(),
	}
}

func (c *RecentCounter) Ping(t int) int {
	c.PushFront(t)
	for t-c.Back().Value.(int) > interval {
		c.Remove(c.Back())
	}
	return c.Len()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func Test(t *testing.T) {
	testCases := []struct {
		pings, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3001}, []int{1, 2}},

		// leetcode case
		{[]int{1, 100, 3001, 3002}, []int{1, 2, 3, 3}},
	}
	for _, tt := range testCases {
		name := fmt.Sprint(tt.pings)
		t.Run(name, func(t *testing.T) {

			calls := Constructor()
			got := []int{}
			for _, v := range tt.pings {
				got = append(got, calls.Ping(v))
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
