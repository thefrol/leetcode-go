package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

// https://github.com/thefrol/leetcode-go

// У детей есть конфеты, и у меня ещё extraCandies штук. Кому если я отдам
// свои конфеты станет самым богатым на конфеты, а кто нет. ответить массивом

// по сути надо отсортировать по убыванию, и вернуть самые старшие числа,
// где конфеты у челика = максимум конфет - мои конфеты

// если мы воспользуемся сортировкой из библиотеки, то мы потеряем связь номера ребенка и
// его конфетами. Но можно хитро потом сверить по количеству конфет

// а можно продемонстрировать владение библиотекой, и сделать сортируемую структуру

// Вообще чет сортировка это оверкилл. Реально можно было посчитать максимум и
// и потом просто по очереди каждое проверить.

// V2 А можно искать максимум и записывать максимальные значения по очередности
// как-то типа от самого большего ... опять сортировка будет ахах

// ИЛи попробовать. Связные список сделать с начало и концом

// да не это опять сортировка, только фенси теперь будет

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	var max int
	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	for i, v := range candies {
		if v+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{candies: []int{4, 2, 1, 1, 2}, extraCandies: 1, want: []bool{true, false, false, false, false}},
		{candies: []int{4, 3, 1, 1, 2}, extraCandies: 1, want: []bool{true, true, false, false, false}},
		{candies: []int{12, 1, 12}, extraCandies: 1, want: []bool{true, false, true}},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v + %d", tt.candies, tt.extraCandies)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, kidsWithCandies(tt.candies, tt.extraCandies))
		})
	}
}

// V1

// type Kid struct {
// 	candy, i int
// }

// type Kids []Kid

// func NewKids(candies []int) (nk Kids) {
// 	for i, c := range candies {
// 		nk = append(nk, Kid{candy: c, i: i})
// 	}
// 	return
// }

// // Len implements sort.Interface.
// func (k Kids) Len() int {
// 	return len(k)
// }

// // Less implements sort.Interface.
// func (k Kids) Less(i int, j int) bool {
// 	return k[i].candy < k[j].candy
// }

// // Swap implements sort.Interface.
// func (k Kids) Swap(i int, j int) {
// 	k[i], k[j] = k[j], k[i]
// }

// var _ sort.Interface = (*Kids)(nil)

// func kidsWithCandies(candies []int, extraCandies int) []bool {
// 	result := make([]bool, len(candies))

// 	// сортируем
// 	kids := NewKids(candies)
// 	sort.Sort(kids)

// 	// проставляем true пока дополнительные кофеы делают ребенка счастливым
// 	max := kids[len(kids)-1].candy
// 	for i := len(kids) - 1; i >= 0 && kids[i].candy+extraCandies >= max; i-- {
// 		// смотрим оригинальный номер ребенка
// 		result[kids[i].i] = true
// 	}
// 	return result
// }
