package main

import (
	"container/ring"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/dota2-senate/

// https://github.com/thefrol/leetcode-go

// Происходит голосование. Голосующий момент забанить следующего кандидата,
// или если он остался последним -провозгласить себя победителем. Есть две
// фракции: Радиант и Дайр.

// Короче тема такая, что мне кажется если соотношение Р:Д 1:2, то радиантам
// уже никак не победить. НУ или в пределах одного человека.

// Но для начала жадный алгоритм. А там можно в начале каждого цикла проверять
// некое условия быстрого окончания. Будем пользоваться кольцевой очередью
// и запоминать начало. Там даже наверное есть функция, которая пробегает по
// очереди и что-то там меняет, считает.

// короче мы идем по очереди. Текущий сенатор банит противоположную франкцию,
// если видит её справа от себя. Иначе ничего не делает. И так пока не
// окажется ноль противников - тогда объявляем победу

// Конечно, когда у нас 4 радианта и 1 дайр уже понятно кто победит.
// я думаю n(r)/n(d)>2 такое должно быть условие окончания. Примерно.
// Почему два? из-за вот такого кейса "РДРДРДДД". 3 радианта побеждают
// пять(!) дайров благодаря тому что ходят первыми. Но шесть дайров бы победили.

// давай посчитаем
// РДРДРДДДД
// Р: РРРДДД;
// Д: РРДДД;
// Р: РРДД;
// Д: РДД;
// Р: РД;
// Д: Д - вин

// То есть предварительно побеждает тех кого в два раза больше ровно, даже если
// они первыми ходить не будут

// время переходить к тестам

func predictPartyVictory(senate string) string {
	sen := ring.New(len(senate))
	rcount, dcount := 0, 0
	bb := []byte(senate)

	// в нулевой итерации мы заполняем очередь и заодно считаем
	// количество сенаторов
	i := sen
	for _, c := range bb {
		i.Value = c

		// подсчитаем количество сенаторов
		switch c {
		case 'R':
			rcount++
		default:
			dcount++
		}
		i = i.Next()
	}

	// основной этап:
	// сначала проверяем количество сенаторов,
	// если одних в два раза больше чем других - объявить победу,
	// иначе: провести раунд банов.
	for {
		// поиск победителя
		if dcount == 0 || rcount/dcount >= 2 {
			return "Radiant"
		}
		if rcount == 0 || dcount/rcount >= 2 {
			return "Dire"
		}
		// баны
		// мы не можем воспользоваться ring.Do,
		// потому что одновременно и меняем цикл и идем по нему
		// поэтому проитерируем по кольцу
		rcount, dcount = 0, 0
		i := sen
		for {
			if i.Value != i.Next().Value {
				// разные фракцие - удаляем следующего
				// надо обратить внимание, что если он был первым,
				// то у нас новый первый)

				del := i.Unlink(1) // удаляем следующего
				if del == sen {
					// если удалили первого
					// то начинает круг теперь второй, потому что
					// этот уже свой голос использовал
					sen = i.Next()
				}
			}

			switch i.Value.(byte) {
			case 'R':
				rcount++
			default:
				dcount++
			}

			i = i.Next()
			// если пришли в начало
			if i == sen {
				break
			}
		}

	}

}

func Test_predictPartyVictory(t *testing.T) {

	tests := []struct {
		senate, want string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
		{"RDRDRDDD", "Radiant"},
		{"RDRDRDDDD", "Dire"},

		// leetcode cases
		{"DR", "Dire"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s->%s", tt.senate, tt.want)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, predictPartyVictory(tt.senate))
		})
	}
}