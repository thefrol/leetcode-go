package main

import (
	"fmt"
	"strings"
)

// Суть такая, мы преобразуем паттерн в другой тип данных - проверятель
// за него отвечает интерфейс Checker.
//
// Каждый проверятель значет что-то о себе, и следющего проверятеля. В
// зависимости от того, какого типа проверятель, он проверяет является ли ему
// переданная подстрока валидной, при этом он запускает и другие проверятели
//
// Но обычно он отвечает за что-то одно, каждый проверятель это подзадача
// таким образом мы реализуем паттерн ДИНАМИЧЕСКОГО ПРОГРАММИРОВАНИЯ

// По иронии тут быстро очень работает рекурсия как я делал с палиндромами, но я все равно красаучик канеш, что таку штуку сделал, топчик

func isMatch(s string, p string) bool {
	return NewChecker(p).Check(s)
}

type Checker interface {
	fmt.Stringer
	Check(string) bool
}

type Data struct {
	Pred string
	Next Checker
}

type EOF struct {
}

func (p EOF) String() string {
	return "!EOF!"
}

func (p EOF) Check(s string) bool {
	return s == ""
}

type Substring struct {
	Data
}

func (p Substring) Check(s string) bool {
	if len(s) < len(p.Pred) {
		return false
	}
	if end, found := strings.CutPrefix(s, p.Pred); found {
		return p.Next.Check(end)
	}
	return false
}

type Dot struct {
	Data
}

func (p Dot) Check(s string) bool {
	return len(s) >= 1 && p.Next.Check(s[1:])
}

type Star struct {
	Data
}

func (p Star) Check(s string) bool {
	for i := 0; i < len(s)+1; i++ { // i это количество использования символа, может быть ноль, это значит ни одного, поэтому прохдим на один болше чем линна строки
		prefix := strings.Repeat(p.Pred[:1], i)

		if end, found := strings.CutPrefix(s, prefix); !found {
			return false

		} else if p.Next.Check(end) {
			return true
		}

	}
	return false
}

type DotStar struct {
	next Checker
}

func (p DotStar) String() string {
	return "ANY_SYMBOLS" + p.next.String()
}

func (p DotStar) Check(s string) bool {
	for i := 0; i < len(s)+1; i++ {
		if p.next.Check(s[i:]) {
			return true
		}
	}
	return false
}

func NewDotStar(pred string, next string) Checker {
	if pred[0] == '.' {
		return DotStar{next: NewChecker(next)}
	}
	return Star{
		Data: Data{
			Pred: pred,
			Next: NewChecker(next),
		},
	}
}
func (s Data) String() string {
	return fmt.Sprintf("([%s]+[%v])", s.Pred, s.Next)
}

func NewChecker(s string) Checker {
	if s == "" {
		return EOF{}
	}

	if len(s) >= 2 && s[1] == '*' {
		return NewDotStar(s[:2], s[2:])

	} else if s[0] == '.' {
		// именно в этой последовательности после проверки на звездочку
		// тут только одинарные звездочки проверяем
		return Dot{
			Data: Data{
				Pred: s[0:1],
				Next: NewChecker(s[1:]),
			},
		}
	} else {
		stop := -1
		for i := 0; i < len(s); i++ {
			if len(s) > i+1 && s[i+1] == '*' {
				stop = i
				break
			}
			if s[i] == '.' {
				stop = i
				break
			}
		}
		if stop == -1 {
			return Substring{Data: Data{Pred: s, Next: EOF{}}}
		}

		return Substring{
			Data: Data{
				Pred: s[0:stop],
				Next: NewChecker(s[stop:]),
			},
		}
	}
	// а можно strings.Split(s,'*')
}

func main() {
	p := "mis*is*ip*."
	s := "mississippi"
	checker := NewChecker(p)
	fmt.Println(checker)
	fmt.Println(checker.Check(s))
}
