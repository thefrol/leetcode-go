package main

import (
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

// идеи для V2 в целом мы можем посчитать наибольший общий делить для длинн строк,
// и просто сравнить потом буквы. ПОнятно, что ABCAB и AB не могут иметь общего делителя
// строкового, потому что у них нет общего делителя вообще

// идея для V3 когда проверяем буквы, можно проверять сразу две строки, потому
// что мне в алгоритме надо именно это проверять

func gcdOfStrings(str1 string, str2 string) string {
	min, other := str1, str2
	if len(min) > len(other) {
		min, other = other, min
	}

	var cd string
	for i := len(min); i >= 0; i-- {
		cd = min[0:i]
		if isCD(min, cd) && isCD(other, cd) {
			return min[0:i]
		}
	}
	return ""
}

// а проверим сразу две строки
// func isCDfast(less, large, cd string) bool {
// сложна
// }

// isCD проверяет является ли общим делителем cd делителем строки str
func isCD(str, cd string) bool {
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
