package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// отсортируем по жирноте, и при каждом добавляемом числе пытаться посчитать количество новых возможных деревьев, приечем
// причем раз новое число больше остальных, то на прошлые короткие деревье оно не может повлиять

// можем содать какой-то кеш с количеством вариков

var m map[int][2]int = make(map[int][2]int)

func numFactoredBinaryTrees(arr []int) int {
	c := 0 // каждый элемент это один вариант
	sort.Ints(arr)
	var curr []int
	for _, item := range arr {
		curr = append(curr, item)
		c += 1 + count(curr, item)
	}
	return c
}

// count считает сколько мы можем построить деревьев от числа int,
// с элементами из arr
func count(arr []int, item int) int {
	c := 0
	for _, left := range arr {
		for _, right := range arr {
			if left*right == item {
				{
					c += count(arr, left)*count(arr, right) + count(arr, right) + count(arr, right) + 1 // Умножение!!! варианты слева умножаются на правые а не складываются

				}

			}
		}
	}
	return c
}

func TestTree(t *testing.T) {
	tests := []struct {
		arr   []int
		trees int
	}{
		{arr: []int{2, 4}, trees: 3},
		{arr: []int{2, 4, 8}, trees: 8},
		{arr: []int{2, 4, 5, 10}, trees: 7},
		{arr: []int{18, 3, 6, 2}, trees: 12},
		{trees: 777, arr: []int{45, 42, 2, 18, 23, 1170, 12, 41, 40, 9, 47, 24, 33, 28, 10, 32, 29, 17, 46, 11, 759, 37, 6, 26, 21, 49, 31, 14, 19, 8, 13, 7, 27, 22, 3, 36, 34, 38, 39, 30, 43, 15, 4, 16, 35, 25, 20, 44, 5, 48}},
		{trees: 666, arr: []int{46, 144, 5040, 4488, 544, 380, 4410, 34, 11, 5, 3063808, 5550, 34496, 12, 540, 28, 18, 13, 2, 1056, 32710656, 31, 91872, 23, 26, 240, 18720, 33, 49, 4, 38, 37, 1457, 3, 799, 557568, 32, 1400, 47, 10, 20774, 1296, 9, 21, 92928, 8704, 29, 2162, 22, 1883700, 49588, 1078, 36, 44, 352, 546, 19, 523370496, 476, 24, 6000, 42, 30, 8, 16262400, 61600, 41, 24150, 1968, 7056, 7, 35, 16, 87, 20, 2730, 11616, 10912, 690, 150, 25, 6, 14, 1689120, 43, 3128, 27, 197472, 45, 15, 585, 21645, 39, 40, 2205, 17, 48, 136}},
	}

	for _, tc := range tests {
		name := fmt.Sprint(tc.arr)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.trees, numFactoredBinaryTrees(tc.arr))
		})
	}
}
