package main

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/can-place-flowers

// https://githib.com/thefrol/leetcode-go

// Тут нужно сажать цветы в клумбы, чтобы они поместились.
// Сделаем по умному, сначала напишем дурацкий алгоритм,
// который проходит все клумбы по одному, а потом будем улучшать

// 1. Когда мы точно сможем посадить все цветы?
// 2. Когда мы точно не посадим
// В остальных случаях проверять

// Ускорить алгоритм "тупой" проверки. проверяя сразу по три грядки
// возможно встречая в следующей ячейки посадку мы можем прыгнуть сразу не сколько ячеек. На три типа или две

// конечно тут дофига ускоренный алгоритм сделал, но господи сколько исключений просто с ума сойти

// ну 98%/98% того стоит, это я ещё не обрабатывал всякие интересные ускорялки

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true // это ловко они в тестах
	}

	// Если у нас один горшок то он должен быть пустой
	if len(flowerbed) == 1 {
		return flowerbed[0] == 0
	}

	i := 1
	last := len(flowerbed) - 1

	//вместо первой итерации
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		n--
		i = 2
		if n <= 0 {
			return true
		}
	}

	// и вместо последней
	if flowerbed[last] == 0 && flowerbed[last-1] == 0 {
		flowerbed[last] = 1
		n--
		if n <= 0 {
			return true
		}
	}

	// основной алгоритм
	for i < len(flowerbed)-2 {
		switch {
		case flowerbed[i+1] == 1:
			// если в следующей клетке, то можем сразу сильно прыгнуть
			i += 3
		case flowerbed[i] == 1:
			// если под нами цветок, то уже через две клетки может быть свободно
			i += 2
		case flowerbed[i-1] == 1:
			i++ // именно в такой последовательности чтобы быстрее прыгать
		default:
			//сажаем цветок
			n--
			i += 2 //прыгаем на две клетки
			if n <= 0 {
				return true
			}
		}

	}

	return false
}

func Test(t *testing.T) {
	testCases := []struct {
		flowerbed string
		plants    int
		want      bool
	}{
		{"10001", 1, true},
		{"10001", 2, false},
		{"100001", 2, false},
		{"1000001", 2, true},
		{"0100001", 2, false},
		{"10101010", 0, true}, // тесты литкода
		{"001010", 1, true},   // пограничный кейс
		{"1000100", 2, true},  // и еще один
	}
	for _, tC := range testCases {
		t.Run(tC.flowerbed+"x"+strconv.Itoa(tC.plants), func(t *testing.T) {
			assert.Equal(t, tC.want, canPlaceFlowers(Ints(tC.flowerbed), tC.plants))

		})
	}
}

func Ints(s string) (ints []int) {
	for _, c := range s {
		switch c {
		case '1':
			ints = append(ints, 1)
		case '0':
			ints = append(ints, 0)
		default:
			log.Fatalf("строка может состоять только из 1 и 0, а тут %v (%s)", c, s)
		}

	}
	return
}
