package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/max-number-of-k-sum-pairs

// https://leetcode.com/problems/max-number-of-k-sum-pairs/solutions/4266353/map-non-paired-98-speed/

// https://github.com/thefrol/leetcode-go

// тут нужно найти максимальное количество пар, которые
// сумме дают k и удалить их из массива

// Короче, одно число полностью определяет второе
// которое мне нужно для k ( k=5: a=1 => b=4)

// Думаю я нашел элегантное решение:
// Создаем мапу lefties[], в которую записываем оставшиеся числа и
// их количество. Алгоритм такой, мы идем по переданному массиву
// и сверяемся - есть ли второе слагаемое в мапе?
// Если есть - убираем число из мапы и идем дальше, такше увеличиваем счетчик ответа
// если нет - добавляет текущее число

// Возвращаем счетчик

// 98 по скорости, 6 по памяти.
// Все потому что наверное там очень
// большие массивы передаются и у
// меня делаются огроиные мапы думаю

func maxOperations(nums []int, k int) int {
	// оптимизация размера памы. Не будем делать очень большую, но и считать
	// количество неповторяющихся элементов тоже не будем
	mapSize := len(nums)
	if mapSize > 100 {
		mapSize = 100 // вот тут можно посчитать количество уникальных элементов
	}
	lefties := make(map[int]int, mapSize)
	counter := 0

	var pair int
	for _, v := range nums {
		pair = k - v
		if n, ok := lefties[pair]; ok && n != 0 {
			// пара найдена
			lefties[pair]--
			counter++
		} else {
			// пара не найдена
			// добавим в массик без пар
			lefties[v]++
		}
	}

	return counter

}

func Test_maxOperations(t *testing.T) {

	tests := []struct {
		nums     []int
		k        int
		want     int
		wantNums []int
	}{
		{
			nums: []int{1, 2, 3, 4}, k: 5,
			wantNums: []int{}, want: 2,
		},
		{
			nums: []int{3, 1, 3, 4, 3}, k: 6,
			wantNums: []int{1, 4, 3}, want: 1,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("k=%d; [%v]", tt.k, tt.nums)
		t.Run(name, func(t *testing.T) {
			got := maxOperations(tt.nums, tt.k)
			assert.Equal(t, tt.want, got, "количество операция нужно другое")
		})
	}
}
