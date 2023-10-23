package main

import (
	"fmt"
	"math"
	"strconv"
)

//use frexp!!!

func isPowerOfFour(n int) bool {
	// нужно чтобы в двоичной записи
	// число имело четное число последних нулей
	// 4=100, 16=10000
	if n == 0 {
		return false // особый случай
	}
	i := 0
	// сдвигаем направо постепенно, смотрим, что справа
	// вылетают только нолики; как только единицы вылетела
	// заканчиваем цикл
	for n != 0 && n-n>>1<<1 == 0 {
		n, i = n>>1, i+1
	}
	math.Log2(123)

	// цикл мог завершиться из-за вылетевшей слева единицы
	// в этом случае в n записно что-то вроде 11,101,
	// Если это старшая единица, то там записно 1 обязательно
	return i%2 == 0 && n == 1
}

func main() {
	printOut(0)
	printOut(1)
	printOut(2)
	printOut(4)
	printOut(5)
	printOut(8)
	printOut(16)
}

func printOut(i int) {
	fmt.Println(i, i>>2, isPowerOfFour(i), strconv.FormatInt(int64(i), 2))

}
