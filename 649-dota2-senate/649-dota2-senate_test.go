package main

import (
	"container/list"
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
// И так пока не, окажется ноль противников - тогда объявляем победу

// даже так, если радиантов больше чем дайров, то понятно что радианты победят.
// возьмем такое условие. 3Р, 2Д.
//
// Сначала радианты хотят первыми
// РРРДД - победа в первом ходе

// Тепреь по другому их расставим
// ДДРРР
// ДДррР
// дДррр - дай р победил, хммм

// если четыре радианта
// ДДРРРР
// ДДррРР, ддррРР - победили радиант

// короче условие победы - одной фракции в два раза больше.
// то есть мы голосуем по кругу, и пока один не побеждает.

// У меня есть мнение, что одного раунда голосования всегда будет
// достаточно, чтобы определить победителя. Не представляю как это
// математически можно доказать, но мне кажется это так.

// А вообще знаю!
// Если дайров больше чем в два раза, то они победили.
// Соотношение 1 < N(Победители)/(проигравшие) < 2
// в течение раунда победители забанят... а вообще мы не знаем...

// Да в целом нам и не важно. Просто надо писатьтак цикл, чтобы
// как можно раньше срабатывал выход. То есть, мы когда наполним
// кольцевую очередь мы посчитаем количество Радиантов и Дайр,
// а потом если кого-то лишаем права голова, то будем удалять постепенно
// И на каждой итерации если кого-то удаляем, вычитать из
// счетчика по фракциям.

// Теперь как мы баним. У нас есть указатель на текущего человека
// и банящий указатель. Банящий указатель бежит вперед, пока не находит
// противоположную фракцию. Потом переход хода. Если следующий участник:
// противоположной фракции, то указатель обнуляется и начинается поиск
// и голосующего.

// А можно сделать два указателя. Один поиск следующего радианта, другой
// поиска следующего дайра.

// Мы можем даже ещё оптимизировать. Типа в первом раунду вести и подсчет
// состава сената по фракциям, и уже банить. zhestko, kanesh.

// Давай для начала что-то простое напишем.
// например два указателя. Но где очередь тогда в чем прикол очереди тут?)

// время переходить к тестам

// V2 короче эти два указателя о которых я писал... могли бы навести меня на
// некоторые мысли. Но не навели. Я подсмотрел решение. И самое быстрое
// это две очереди. Одна для радиантов, другая для Дайров.
//
// На каждом этапе мы достаем одного радианта и одного дайра.
// кто раньше в очереди тот и побеждает. Проигравший выбывает,
// победитель идёт в конец очереди своей франкции
//
// вот тут хорошее объяснение с картинками
// https://leetcode.com/problems/dota2-senate/solutions/3483399/simple-diagram-explanation/?envType=study-plan-v2&envId=leetcode-75
//

func predictPartyVictory(senate string) string {
	rs := list.New()
	ds := list.New()

	for i, v := range senate {
		switch v {
		case 'R':
			rs.PushBack(i)
		default:
			ds.PushBack(i)
		}
	}

	// интересно могу ли я улучшить этот алгоритм своей системой раннего завершения.
	// типа ещё проверять количество радиантов и дайров

	for {
		// ранний выход
		if rs.Len() == 0 || ds.Len()/rs.Len() >= 2 {
			return "Dire"
		}
		if ds.Len() == 0 || rs.Len()/ds.Len() >= 2 {
			return "Radiant"
		}

		// голосование
		rv := rs.Front().Value.(int)
		dv := ds.Front().Value.(int)
		if rv < dv {
			// тонкий момент, что мы считаем номер в очереди
			// следущего раунда не с начала, а как бы с какого-то числа другого
			// то есть если сенатор с номерком 2 победил, он идет в конце очереди
			// и становится там 100 ( если всего сенаторов меньше 100)
			//
			// И нонечно очевидно(но не оч), что номерок он получает как
			// свой номер + количество сенаторов. Так он получает
			// свой относительный порядок в рамках раунда
			rs.PushBack(rv + len(senate))
		} else {
			ds.PushBack(dv + len(senate))
		}

		rs.Remove(rs.Front())
		ds.Remove(ds.Front())

	}

}

func Test_predictPartyVictory(t *testing.T) {

	tests := []struct {
		senate, want string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
		{"RDRDRDDD", "Dire"},
		{"RDRDRDDDD", "Dire"},

		// leetcode cases
		{"DR", "Dire"},
		{"DRRDRDRDRDDRDRDR", "Radiant"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s->%s", tt.senate, tt.want)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, predictPartyVictory(tt.senate))
		})
	}
}
