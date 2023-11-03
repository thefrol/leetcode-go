package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/greatest-common-divisor-of-strings/

// https://github.com/thefrol/leetcode-go

// Будем делать как поиск выбором по стратегии разделяй и властвуй
// Находим самую короткую строку, делим её пополам. Если это делитель, то
// увеличиваем до половины оставшегося интервала. и так далее

// Но это так не работает. Возьмем ABC и ABCABCABC. Если мы поделим ABC пополам,
// то получим A, он не является делителем. Значит алгоритм попробует делить ещё сильнее
// и в итоге не найдет вообще общего делителя

// V2 в целом мы можем посчитать наибольший общий делить для длинн строк,
// и просто сравнить потом буквы. ПОнятно, что ABCAB и AB не могут иметь общего делителя
// строкового, потому что у них нет общего делителя вообще. ТУт нам поможет алгоритм евклида
// для начала выходить из алгоритма если общий делитель не существует

// ИДЕЯ для V3, начинать не с полной длинны, а с НОД по алгоритму ЕВКЛИДА

// идея для V; когда проверяем буквы, можно проверять сразу две строки, потому
// что мне в алгоритме надо именно это проверять

func gcdOfStrings(str1 string, str2 string) string {
	lesser, larger := str1, str2
	if len(lesser) > len(larger) {
		lesser, larger = larger, lesser
	}

	var cd string
	for i := GCD(len(larger), len(lesser)); i >= 0; i-- {
		cd = lesser[0:i]
		if isCD(lesser, cd) && isCD(larger, cd) {
			return lesser[0:i]
		}
	}
	return ""
}

// GCD считает НОД для двух чисел, обязательно надо передать большее и меньшее
func GCD(larger, lesser int) int {
	if larger%lesser == 0 {
		return lesser
	}
	return GCD(lesser, larger%lesser)
}

// isCD проверяет является ли общим делителем cd делителем строки str
func isCD(str, cd string) bool {
	// тут тоже можно сначала проверять что строки делятся друг на друга
	if len(cd) == 0 {
		return true
	}

	if len(str)%len(cd) != 0 {
		return false
	}

	ptr := 0
	for {

		if ptr+len(cd) > len(str) {
			return false
		}
		if str[ptr:ptr+len(cd)] == cd {
			ptr += len(cd)
		} else {
			return false
		}

		if ptr == len(str) {
			return true
		}
	}
	// эта функция могла бы оптимизировать поиск
	// типа сообщать что общий длитель в два раза короче
	// на какой цифре сбился поиск, пока не понимаю как это использовать тока
}

// а проверим сразу две строки
// func isCDfast(less, large, cd string) bool {
// сложна
// }

func Test_gcdOfStrings(t *testing.T) {

	tests := []struct {
		str1, str2 string
		want       string
	}{
		//{str1: "ABABABA", str2: "ABAB", want: ""},
		//{str1: "ABABABAB", str2: "ABAB", want: "AB"},
		{str1: "ABCABCABC", str2: "ABC", want: "ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.str1+"+"+tt.str2, func(t *testing.T) {
			assert.Equal(t, tt.want, gcdOfStrings(tt.str1, tt.str2))
		})
	}
}

func Test_isCD(t *testing.T) {

	tests := []struct {
		str, cd string
		want    bool
	}{
		{str: "ABAB", cd: "AB", want: true},
		{str: "ABAB", cd: "ABA", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.str+"/"+tt.cd, func(t *testing.T) {
			assert.Equal(t, tt.want, isCD(tt.str, tt.cd))
		})
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		larger, lesser int
		want           int
	}{
		{5, 3, 1},
		{10, 5, 5},
		{15, 10, 5},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("NOD(%d,%d)", tt.larger, tt.lesser)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, GCD(tt.larger, tt.lesser))
		})
	}
}
