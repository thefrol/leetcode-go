package main

import (
	"fmt"
	"strconv"
)

// func isPowerOfFour(n int) bool {
// 	// n=fraction*e^pow fraction in [0.5,1)
// 	// берем fraction=0,5 но это означает что
// 	// должны быть нечетные степени, например
// 	// 4=0.5*8=0.5*2^3 fraction=0,5 pow=3
// 	fraction, pow := math.Frexp(float64(n))
// 	return pow%2 == 1 && fraction == 0.5
// }

func isPowerOfFour(t int) bool {
	n := int32(t)
	if n < 1 {
		return false
	}
	// просто проверяем что правые два бита из нулей, и сдвигаем всправо
	// как только перестанут быть нулями, цикл остановится
	// чтобы проверить правые биты мы делаем побитовое и с 3=0b11
	for ; n&0b11 == 0; n = n >> 2 {
	}
	return n == 1
}

// в идеале вытесненные биты мы вообще можем из какого-то там регисра процессора достать
// тут наверное можно на asm вставку сделать

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
