package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// можно, чтобы было меньше значений в функциях сделать замыкания
// но по сути мы все сравно слайс по ссылке передаем, не большая проблема

// Можно ещё экранировать значение узла от количества, чтобы там не перемудрить

// Может быть можено как-то оптимизировать при помощи мапы поиск,

func numFactoredBinaryTrees(arr []int) int {
	c := 0 // каждый элемент это один вариант
	for _, item := range arr {
		c += count(arr, item)
	}
	return c
}

// count считает сколько мы можем построить деревьев от числа int,
// с элементами из arr
func count(arr []int, item int) int {
	c := 1
	for _, left := range arr {
		for _, right := range arr {
			if left*right == item {
				{
					if left != right {
						c += count(arr, left) * count(arr, right) // Умножение!!! варианты слева умножаются на правые а не складываются
					}
					// можно ускорить, если обрабатывать когда left=right
					if left == right {
						vars := count(arr, left)
						c += vars*(vars-1) + 1
					}

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
	}

	for _, tc := range tests {
		name := fmt.Sprint(tc.arr)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.trees, numFactoredBinaryTrees(tc.arr))
		})
	}
}
