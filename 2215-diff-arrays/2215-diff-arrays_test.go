package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-difference-of-two-arrays/

// http://github.com/thefrol/leetcode-go

// // Нужно найти разницу двух массивов
// // один массива отдать a-b, а другой b-a
// // и вернуть

// что я думаю делать:
//
// создаем мапу.
// Если число из массива а, прибавляем 1
// Если число из b отнимаем 1
//
// по итогу: если 5 содержится в обоих массивах
// то в мапе ноль, значит это число пропускает в
// конце.

// Числа могут повторяться, но это мне не помешает

// Как оказалось, числа не уникальны, это меняет дело.
// Например, если в левом массиве число встречается 3 раза,
// а в правом пять - то в мапе будет +2, хотя это число
// вообще не должно содержаться в выходе

// но тут мне поможет битовая логика.
// я как бы буду записывать как бы последовательность
// булевых значение, но в одном числе
// 0b00 - если нет ни там не там
// 0b01 - если только в правом
// 0b10 - если только в левом
// 0b11 - если в обоих
// проверить содержание можно булевой логикой

// сложность O(n)

const (
	left  = 0b10
	right = 0b01
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	m := make(map[int]int)
	len1, len2 := len(nums1), len(nums2)

	i := 0
	for i = 0; i < min(len1, len2); i++ {
		m[nums1[i]] |= left
		m[nums2[i]] |= right
	}

	////// осталось раскидать хвост длинного массива

	var widerSource []int
	var operation int // тут будет число с которым надо побитово сложить

	if len2 > len1 {
		widerSource, operation = nums2, right
	} else {
		widerSource, operation = nums1, left
	}

	for ; i < len(widerSource); i++ {
		m[widerSource[i]] |= operation // фух это было не просто
	}

	// теперь разберем где что

	r1, r2 := make([]int, 0, len1), make([]int, 0, len2)
	for num, count := range m {
		switch {
		case count == left:
			r1 = append(r1, num)
		case count == right:
			r2 = append(r2, num)
		}
		// в иных случах 0b11 содержится в обоих
	}

	// готово!

	return [][]int{r1, r2}
}

func Test_findDifference(t *testing.T) {

	tests := []struct {
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{[]int{1, 2}, []int{3, 4},
			[][]int{{1, 2}, {3, 4}}},
		{[]int{1, 2}, []int{1, 3},
			[][]int{{2}, {3}}},
		{[]int{1, 2, 4, 3, 16, 25}, []int{1, 3, 9, 4},
			[][]int{{2, 16, 25}, {9}}},

		// leetcode testcase
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2},
			[][]int{{3}, {}}},

		{[]int{26, 48, -78, -25, 42, -8, 94, -68, 26}, []int{61, -17},
			[][]int{{48, -25, 42, -8, 26, 94, -68, -78}, {61, -17}}},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v-%v", tt.nums1, tt.nums2)
		t.Run(name, func(t *testing.T) {
			w1, w2 := tt.want[0], tt.want[1]

			got := findDifference(tt.nums1, tt.nums2)
			g1, g2 := got[0], got[1]
			assert.ElementsMatch(t, w1, g1)
			assert.ElementsMatch(t, w2, g2)
		})
	}
}

func TestBits(t *testing.T) {
	assert.Equal(t, 0b10, 0b10&0b11)
	assert.Equal(t, 0b00, 0b10&0b01)

	assert.Equal(t, 0b10, 0b00|0b10)
}
