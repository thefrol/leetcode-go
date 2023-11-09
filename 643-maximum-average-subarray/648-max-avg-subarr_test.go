package main

import (
	"container/ring"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-average-subarray-i/

// https://github.com/thefrol/leetcode-go

// задача получить максимальную сумму элементов
// в подмассиве длинной k

// как бы у нас будет идти скользящее окно, прибавляем
// новое значение убираем одно из старых, а чтобы
// хранить значения окна используем кольцевой список

func findMaxAverage(nums []int, k int) float64 {
	r := NewPushRing(k)
	len := len(nums)

	sum := 0.0

	// просто забрасываем первые k элементов в буфер
	i := 0

	var v float64
	for ; i < k; i++ {
		v = float64(nums[i])
		r.push(v)
		sum += v
	}
	// сумма на первой итерации самая большая на данный момент
	maxsum := sum

	for ; i < len; i++ {
		v = float64(nums[i])
		sum += v - r.pushpop(v)
		if sum > maxsum {
			maxsum = sum
		}
	}
	return maxsum / float64(k)
}

type PushRing struct {
	r *ring.Ring
}

func NewPushRing(n int) PushRing {
	if n <= 0 {
		panic("should be >0")
	}

	return PushRing{
		r: ring.New(n),
	}
}

// pushpop - добавляет в очередь элемент, и достает тот,
// что вывалился из нее
func (pr *PushRing) pushpop(new float64) (val float64) {
	val, pr.r.Value = pr.r.Value.(float64), new
	pr.r = pr.r.Next()
	return
}

func (pr *PushRing) push(new float64) {
	pr.r.Value = new
	pr.r = pr.r.Next()
}

func Test_findMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{nums: []int{1, 12, -5, -6, 50, 3},
			k: 4, want: 12.75},
		{nums: []int{5},
			k: 1, want: 5},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%d in %+v", tt.k, tt.nums)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMaxAverage(tt.nums, tt.k))
		})
	}
}
func TestRing(t *testing.T) {
	r := NewPushRing(4)
	r.push(1)
	r.push(2)
	r.push(3)
	r.push(4)
	assert.Equal(t, 1., r.pushpop(-1))
	assert.Equal(t, 2., r.pushpop(-1))
	assert.Equal(t, 3., r.pushpop(-1))
	assert.Equal(t, 4., r.pushpop(-1))
	assert.Equal(t, -1., r.pushpop(-1))
}
