package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/increasing-triplet-subsequence/

// https://github.com/thefrol/leetcode-go

// // тут нужно найти в массиве чисел подпоследовательность номеров
// i<j<k : num(i)<num(j)<num(k)

// это задача ДП
// найти кортеж из n последовательних чисел, что
// n1<n2<n3..<nK
//
// случай 1 - len(N)>0
// случай 2 - index(Max)--index(min)>1 - и частный случай min()<Nk ( последний элемент)
// случай 3 - если случай два выполняется на отрезке (min,last]
// случай 4 - случай 3, но на отрезке (min, last] текущее число не включено
//
// Это работает в том члучае если index(Max) больше по индексу среди всех максимумов. Потому что там могут быть одинаковые
// а indexMin наоборот
// а значит максимум считаем только во втором случае

func increasingTriplet(nums []int) bool {
	return hasSequence(nil, 3)
}

// hasSequence проверяет, есть ли в nums 3 последовательно увеличивающихся числа
// условие len(nums)>count
func hasSequence(nums Nodes, count int) bool {
	switch {
	case count <= 1:
		return true
		// ускорить, убрав повторяющиеся числа
	case count == 2:
		// IndexMax(nums)-IndexMin(nums) > 1

	default:
		//imax := IndexMax(nums)
		//inum := IndexMin(nums)

		// seq := nums[IndexMin(nums)+1:] // min not included
		// return hasSequence(seq, count-1)
	}
	return true
}

// IndexMax возвращает индекс максимального значения,
// если максимальных значений несколько вернет самое большое
func IndexMax(nums []int) (index int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[index] {
			index = i
		}
	}
	return
}

type index int

// seq позвращает последовательности индексов по возрастанию значений
// если её переписать на другую сортировку ну нубдет быстрее
// я думаю для разных лоинн интов можно разные функции использовать

type Node struct {
	index int
	val   int
}

type Nodes []Node

func (ns Nodes) Cut(start int, end int) Nodes {
	//nodes := make(Nodes, 0, end-start)
	for n := range []Node(ns) {
		fmt.Printf("n: %v\n", n)
	}
	return nil
}

// Len implements sort.Interface.
func (n Nodes) Len() int {
	return len(n)
}

// Less implements sort.Interface.
func (n Nodes) Less(i int, j int) bool {
	return n[i].val < n[j].val
}

// Swap implements sort.Interface.
func (n Nodes) Swap(i int, j int) {
	n[i], n[j] = n[j], n[i]
}

func NewNodes(nums []int) Nodes {
	ns := make(Nodes, len(nums))
	for i, v := range nums {
		ns[i] = Node{
			index: i,
			val:   v,
		}
	}
	return ns
}

var _ sort.Interface = (*Nodes)(nil)

// IndexMin возващает индекс минимального числа,
// если минимальных чисел несколько возвращает самое маленькое
func IndexMin(nums []int) (index int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return
}

func Test(t *testing.T) {
	testCases := []struct {
		ints []int
		want bool
	}{
		// {[]int{12, 3}, false},
		// {[]int{3, 12}, false},
		// {[]int{1, 3, 4}, true},
		// {[]int{1, 1, 1, 4}, false},
		// {[]int{1, 5, 4, 4}, false}, // да уж, k не так-то просто искать будет

		// // // i не обязательно всегда минимум глобальный.
		// {[]int{6, 7, 1, 9}, true}, // даа жесть
		{[]int{1, 5, 0, 4, 1, 3}, true},

		// но для быстрой работы можно поискать минимум
		// может как-то отсортировать числа даже
		// и постепенно отваливаться к более общему решению

		// но сначала надонаписать общее решение
	}
	for _, tC := range testCases {
		name := fmt.Sprint(tC.ints)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tC.want, increasingTriplet(tC.ints))
		})
	}
}

// // тут нужно найти в массиве чисел подпоследовательность номеров
// // i<j<k : num(i)<num(j)<num(k)

// // думаю мы просто идём по числам.
// // 1. ищем максимум и минимум
// // 2. если нашли - сразу возвращаем
// // 3. надо ещё третье число

// // см тесты, все не так просто

// // как насчет работать по методике разделяй и властвуй. Как быстрый поиск.
// // выбераем среднее числа. И ищем слева и справа.

// // возможно для начала массив стоит отсортировать, чтобы брать числа у которых
// // вообще чисто теоретически возможны  такие значения

// // а может в процессе сортировки и как бы что-то да обнаружится лол

// // если массив очень большой, то предварительная сортировка все оч замедлит

// // мы сделаем как бы частный случай сортировки быстрым поиском.

// // V2 все хорошо, только бы быстрее она до краев доходила

// // V3 все оч просто

// func increasingTriplet(nums []int) bool {
// 	// попробую массивы не создавать
// 	// только массив проверенных чисел

// 	// ещё можем сохранять какой-то массив
// 	// есть ли у нас числа больше такого-то

// 	// тут будут граничные кейсы
// 	if len(nums) < 3 {
// 		return false
// 	}

// 	// базовый случай

// 	if len(nums) == 3 {
// 		return nums[0] < nums[1] && nums[1] < nums[2]
// 	}

// 	// основной алгоритм
// 	// если число больше предыдущего минимума,
// 	// прибавить счетчик
// 	count := 0 // количество последовательных значений, одно больше другого

// 	// главная фишка найти что-то что меньше последнего элемента
// 	// именно отсюда начинается рассчет последовательных значений
// 	// Это гарантирует, что последнее число уже, и мы ищем
// 	// ещё два последовательных с начала

// 	// в ином случае, нам надо было бы сбрасывать счетчик минимального,
// 	// и искать заново

// 	// 7,10,15,8,6(ищем отсюда)

// 	// сращу пришло решение игнорировать опускающиеся последовательности. Типа по
// 	// очереди их вырезать.
// 	// но задача вырезки довольно сложная

// 	// то есть нам надо найти длинну неубывающей подпоследовательности, >3
// 	// с вырезанием все равно не так просто

// 	//last=len(nums)-1
// 	iMin := int(Min(nums))
// 	if len(nums)-iMin < 3 {
// 		return false
// 	}

// 	// тут найдем меньшее
// 	// и в оставшемся отрезке большее, не включая первое

// 	// теперь просто ищем два числа больших, чем lessLast

// 	// Блин вот и решение задачи собсно лол
// 	// чего я там городил
// 	// жадный алгоритм, который мы ограничили
// 	mins := nums[iMin]
// 	for i := int(iMin); i < len(nums) && count < 3; i++ {
// 		if nums[i] > mins {
// 			count = count + 1
// 			mins = nums[i]
// 		}
// 	}
// 	return count >= 2
// }

// type index int

// func Min(nums []int) index {
// 	iMin := 0
// 	for i, v := range nums {
// 		if v < nums[iMin] {
// 			iMin = i
// 		}
// 	}
// 	return index(iMin)
// }
