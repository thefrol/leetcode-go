package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// можно, чтобы было меньше значений в функциях сделать замыкания
// но по сути мы все сравно слайс по ссылке передаем, не большая проблема

// Можно ещё экранировать значение узла от количества, чтобы там не перемудрить

// Может быть можено как-то оптимизировать при помощи мапы поиск, может даже
// там можно собрать что на что можно умножить, и что получится)
// Вообще хорошая идея. Типа по каждому элементу,
// умноженный элемент, который мы можем получить

func numFactoredBinaryTrees(arr []int) int {
	count := 0

	for _, v := range arr {
		count += node(arr, v)
	}
	return count
}

// node считает, сколько мы можем построить значений от этого листа
func node(arr []int, curr int) int {
	count := 1

	for _, v := range next(arr, curr) {
		count += node(arr, v)
	}
	return count
}

// next выдает список новых значений от этого листа,
// которые получаются умножением этого листа на какой-то другой
// выдается именно уже умноженный, второй элемент нам не оч интересен,
func next(arr []int, item int) []int {
	var list []int
	for _, v := range arr {
		if has(arr, v*item) {
			// получили новый элемент, который тоже в списке
			list = append(list, v*item)
		}
	}
	return list
}

// проверяет, что item в массиве
func has(arr []int, item int) bool {
	for _, v := range arr {
		if item == v {
			return true
		}
	}
	return false
}

// prev находит два множителя при движенит в обратную сторону. Найти для этого элемента из чело можно собрать
func prev(arr []int, curr int) [][]int {
	var items [][]int
	for _, a := range arr {
		for _, b := range arr {
			if has(arr, a*b) {
				items = append(items, []int{a, b})
			}
		}
	}
	return items
}

func TestTree(t *testing.T) {
	tests := []struct {
		arr   []int
		trees int
	}{
		{arr: []int{2, 4}, trees: 3},
		{arr: []int{2, 4, 5, 10}, trees: 7},
		{arr: []int{18, 3, 6, 2}, trees: 12},
	}

	for _, tc := range tests {
		name := fmt.Sprint(tc.arr)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.trees, numFactoredBinaryTrees(tc.arr))
		})
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		arr      []int
		item     int
		expected []int
	}{
		{arr: []int{18, 3, 6, 2}, item: 3, expected: []int{6, 18}},
	}

	for _, tc := range tests {
		name := fmt.Sprint(tc.arr, tc.item)
		t.Run(name, func(t *testing.T) {
			real := next(tc.arr, tc.item)
			assert.ElementsMatchf(t, tc.expected, real, "Was %v but should be %v", real, tc.expected)
		})
	}
}
