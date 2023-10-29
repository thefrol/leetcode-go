package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Вообще, если мы будем двигаться имеено с краев и кешировать значение IsPalindrome
// на каждом этапе провери, то есть и значения для краев с обрезанными краями, то все будет быстрее работать
// то есть в хеше мы храним не столько ахха=true , но и хх=true, да и вообще
// лучше хранить не строку, а номера символов начала и конца
//
// Если с 6 по 10 символ - палиндром, но нам не надо проверять это ещё раз.
// кроме того это значит, что 8 и девятый образуют палиндром, думаю это поможет значительно в поисках
//
// # Потому что я много раз запускаю поиск по одним и тем же строкам
//
// Плюс постарайся уложиться в стек

// возможно если мы нашли какой-то большой палиндром, то другие более короткие версии заканчиваем искать.
// думаю такой случай может быть, когда мы ищем левый вариант строки, в то время как нашли что-то справа уже

// V2 стоит поменять принцип работы КЕША. Он должен как можно быстрее говорить, что тут палиндрома нет
// Значит если кеш==тру, тут никаких палиндромов

// V3 использую оба кэша

// V4 можем не пользоваться этим странным шагающим окном, по сути мы все равно проверяем
// все пары начала и конца. Сделаем все одним большим циклом, и просто последим, чтобы
// как можно быстрее все это закончилось)))

// ИДЕЯ V5, мне кажется более быстрее будет растить палиндромы, что думаешь? от каждого символа и и от каждых двух

var cache [][]bool
var nopal [][]bool

func longestPalindrome(s string) string {
	cache = make([][]bool, len(s))
	nopal = make([][]bool, len(s)) // возможно сделать и сплошной массив, главное потом с индексами не напутать)
	for i := 0; i < len(s); i++ {  // плюс один опять из-за проблем с левой границей
		cache[i] = make([]bool, len(s))
		nopal[i] = make([]bool, len(s))
	}

	var longest []byte
	source := []byte(s)
	for l := 0; l < len(s); l++ { // можно потом -1 сделать
		if len(s)-l < len(longest) {
			break
		}
		for r := l; r < len(s); r++ {
			if r-l+1 <= len(longest) { // плюс один, потому что, когда l=r они указывают на один символ, по сути длинна сейчас 1
				continue // ранний сброс, если уже что-то да нашли
			}
			if isPalindrome(source, l, r) {
				longest = source[l : r+1]
			}
		}
	}
	return string(longest)
}

func isPalindrome(source []byte, left, right int) bool {
	// сначала проверим кэш
	if nopal[left][right] == true {
		return false
	}
	if cache[left][right] == true {
		return true
	}
	// теперь уже как обычно
	if len(source) == 0 {
		return false
	}
	if len(source) == 1 {
		return true
	}

	for l, r := left, right; r-l >= 1; l, r = l+1, r-1 {
		// сначала проверяем кэш
		if cache[l][r] == true {
			break // чтобы записать в кеш
		}
		if nopal[l][r] == true {
			for l, r := l, r; l == left && r == right; l, r = l-1, r+1 {
				nopal[l][r] = true // возможно если у нас есть более глобальный палиндром, то можно как бы с более мелкими не заморачиваться, которые внутри большого
			}
			return false // чтобы записать в кеш
		}

		if source[l] != source[r] {
			for l, r := l, r; l == left && r == right; l, r = l-1, r+1 {
				nopal[l][r] = true // возможно если у нас есть более глобальный палиндром, то можно как бы с более мелкими не заморачиваться, которые внутри большого
			}

			return false
		}
	}
	for l, r := left, right; r-l > 1; l, r = l+1, r-1 {
		cache[l][r] = true // возможно если у нас есть более глобальный палиндром, то можно как бы с более мелкими не заморачиваться, которые внутри большого
	}
	return true
}

func TestTooLong(t *testing.T) {
	longestPalindrome("babaddtattarrattatddetartrateedredividerb")
}

func TestTable(t *testing.T) {
	testCases := []struct {
		input      string
		palindrome string
	}{
		{input: "a", palindrome: "a"}, // 114 testcase
		{input: "badad", palindrome: "ada"},
		{input: "cbbd", palindrome: "bb"},
		{input: "sisiooppoosos", palindrome: "ooppoo"},
		{input: "sisiooppoososaaddaaccaaddaa", palindrome: "aaddaaccaaddaa"},
		{input: "babaddtattarrattatddetartrateedredividerb", palindrome: "ddtattarrattatdd"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			assert.Equal(t, tC.palindrome, longestPalindrome(tC.input))
		})
	}

}
