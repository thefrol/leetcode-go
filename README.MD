# Мой литкод

Литкод по большему счету это искусство написания тестов

---

Здесь собраны мои решения задач с литкода на языке `Go`. То, что уже решено лежит в папке `done`, над чем я ещё думаю - лежит в корне. Все что можно улучшить, Или мне хотелось бы, или есть идеи - в папке `ENHANCE`

[Литкод профиль](https://leetcode.com/thefrol/)

![скриншот, где мне апплодирует литкод](./assets/aplause-on-2352.png)
![скриншот окна результатов, где я в топе](./assets/top1.png)
![скриншот окна результатов, где я в топе](./assets/top2.png)

## Что я понял

### Тесты топчик

Иногда можно впасть в бешенство, когда что-то перестает работать. И словно эта верхая башенка внутреннего понимания и уверенности, словно карточный домик начинает рушится, и поймать её невозможно. Если такое случилось, то жизнь просто тускнеет. Тесты будет словно бы опорой, бетонной стеной карточного домика, которая удержит его какое-то время. По крайней мере, станет заметно, что что-то пошло не так.

Стоит один раз написать нормально тесты, как работа над задачкой ускоряется в разы просто. Кучу времени экономит просто сама функция с тестами, например, табличными. Куда можно докидывать тесты в процессе проверки литкодом значений.

Я прокачал и скорость написания тестов, и оформления. Понял на практике какие бывают краевые случаи и что стоит проверять.

Кроме того тесты дают уверенности, они ещё говорят о том, что **все так**, все хорошо.

## Злость и раздражение

Литкод, конечно, научил меня справляться с гневом. Это странное чувство, что я уже все ненавижу, и это заставляет меня продолжать работать, забегать на кухню похрать что плохо леит и бежать обратно. При этом чувствую я себя не очень.

Ничто так сильно не мешает решать задачи, как вот дурные эмоции.

Нужно учиться вылипать из задачи, восстанавливать дыхание, делать растяжку, чтобы хоть как-то трезво взглянуть на мир.

К сожалению, прям как в реальной жизни.

## Разбивай тесты на подзадачи

Основная концепций динамического программирования, что мы делим задачи на подзадачи и решаем снчала более мелкие задачи, которые объединяются в более тяжелые задачи.

С тестами то же самое.

Например, у меня не работала проверка регекс для строки `миссиссиппи`, и меня такой пример испугал поначалу, а потом я сделал вот такое.

Сначала я сделал тест на `мисси`, потом на `миссисси`, потом на `миссиссип`, а значит проблема оставалась в последних двух буквах и я быстро ее нашел.

## Глобальный кеш

Заполняй данные так, чтобы для разных тестов их можно было переиспользовать. Глобальные переменные будут храниться в памяти между тестами! Это преимущество в скорости.

## Понимание тестов говорит о том, что мы понимаем что мы делем

Когда я готовлю тесты, я уже заранее проверяю всякие случаи. Типа а если вот сюда поставить палиндром, а если в конец, а если два палиндрома есть, а между ними подлинней. Как будет работать код?

Иногда мне страшно писать тест от которого мой код может ошибиться.

## Скорость работы

Удивительно, что тесты с похожим принципом работы могут работать медленней на порядки. У меня была функция, которая создавала задачи рекурсивно. Это занимало кучу стека, и видимо она повторяла себя. **Гораздо бысьрее оказалось пройти брутфорсом** по всем значениям - это уксорило работу алгоритма в 300 раз, и видимо изменило даже его сложность вычислительную

## Простая логика выдает топовые результаты

Я даже особо не думал, но простая логика выдала мне 100 скорости 99 памяти. А это я ещё не все оптимизации сделал.

## Иногда дурацкий алгоритм быстрее

Как в `1431`, причем прошлое мое открытие говорит о том же самом собсно. Тупой алгоритм, плюс раннее срабатывание основанное на **конкретных условиях задачи** и топовый результат в кармане.

## Как работать

Сначала пишу мысли, потом тесты

Потом пишу самый тупой алгоритм и обрабатываем не боясь пограничные случаи, это может быть сложно. Зато мы точно знаем, что если что тупой алгоритм сработает как надо.

В конце добавляем быстрый выход.

Так что алгоритм выглядит так

+ инициализация
+ аналитический выход быстро
+ пограничные случаи
+ основной медленный алгоритм

Это вот путь **жадного алгоритма**

## Пограничные случаи

Обрати снимание на самые короткие подзадачи: пустая строка, один символ, два символа. Обычно в них черт и водится

## Литкод дал узнать лучше библиотеку стандартную

например `strings.Fields()` поделит строку и уберет лишние пробелы, а ещё есть `bits.TrailingZeros()` `bits.Div` `math.Frexp()`

## Код работает медленно, если пишем туда же откуда читаем?

238 задача

мой код, без лишних переменных без лишних массивов. считает 17 мс

```go
func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    var i int
    res[0] = 1 // кроме первого(нулевого)
    for i = 1; i < len(nums); i++ {
        res[i] = res[i-1] * nums[i-1]
    }

    buf := 1
    for i = len(nums) - 2; i >= 0; i-- {
        buf *= nums[i+1]
        res[i] *= buf
        // на последней итерации будет записано, но не используется уже
    }

    return res
}
```

а вот код каких-то челиков с двумя массивами лишними работает 7мс, то **в 2,5 раза быстрее**

```go
func productExceptSelf(nums []int) []int {
    prod := make([]int, len(nums))
    revprod := make([]int, len(nums))

    cmult1 := 1
    cmult2 := 1

    for i := 0; i < len(nums); i++ {
        cmult1 = cmult1 * nums[i]
        prod[i] = cmult1

        revindex := len(nums)-1-i
        cmult2 = cmult2 * nums[revindex]
        revprod[revindex] = cmult2
    }

    ans := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        left := 1
        right := 1

        if (i > 0) {
            left = prod[i-1]
        }
        
        if (i < len(nums) -1) {
            right = revprod[i+1]
        }
        ans[i] = left * right
    }

    return ans
}
```

Связано ли это с тем, что мы читаем оттуда же куда пишем и го не может оптимизировать как-то код?

## Иногда написать тесты это почти решить задачу

Пишу тесты, и решим сразу примеров накидать и вот что я вижу для [334](https://leetcode.com/problems/increasing-triplet-subsequence/)

+ `{{12, 3}, false}`, мало чисел, если меньше трех выходим
+ ``{{3,12},false}`, а если их три надо проверку i,j,k сделать
+ `{{1,3,4},true}`
+ `{{1,1,1,4},false}`,  нету среднего числа
+ `{{1,5,4,4},false}`, есть максимум но j>k.

задача сразу словно написалась, а потом словно и распалась. Я увидел кое-какие случаи, где не все так просто и очевидно

PS. Как по итогу оказалось мой запланированный алгоритм вообще не работает. 
НО я тут же придумал другой, более красивый, все **благодаря тестам**

## Строить алгоритм из частного случая может быть ошибкой

[334](https://leetcode.com/problems/increasing-triplet-subsequence/)

```go
// случай 1 - len(N)>0
// случай 2 - index(Max)-index(min)>1 - и частный случай min()<Nk ( последний элемент)
// случай 3 - случай 2 на отрезке (min-max] - мин не включен
// случай 4 - случай 3 на отрезке
// ...
//
```

я стоил алгорит из частного случая 2 и поэтому у меня слишком ветвилось все. Было ужасно. Надо было сразу к ДП перейти.

## Результат очень сильно зависит от размера слайса

Что не удивительно. Там в зависимости от этого создаются бакеты и пересоздаются. Ну больше или меньше коллизий

size=len - 85 мсек, 12 мб
size=100 - 97 мсек, 9 мб
size=20 - 87 мсек, 8,5 мб
size=10 - 103 мсек, 8,4 мб

Вот клевая [задачка 1690](https://leetcode.com/problems/max-number-of-k-sum-pairs/) чтобы понять

## Скорость или память?

```go
// меньше памяти
for {
   pair := k - nums[l]
   //..
}

// быстрее
pair := -1
for {
   pair = k - nums[l]
   //...
}
```

Сфокусируйся на чем-то одном. Выбери быстрота или скорость нужна. Вообще, что ты улучшаешь?

## `for` по структурам

Предположим у нас есть такие данные

```go
type Person struct{
    name, surname string
}
persons:=[]Person{ Person{"Vasya","Pupkin"},{"Dasha","Seksi"}, }
```

Теперь попробуем поменять засекретить имена наших разведчиков

```go
for _,p:=range persons{
    p.surname="***ЗАСЕКРЕЧЕНО***"
}
```

Однако смогли ли мы поменять значения в нашем массиве? Попробуем вывести значения массива `persons`  и получим:

```plain
{Vasya Pupkin}, {Dasha Seksi}
```

А все потому что range передает элементы по значению. Порой нам нужно все-таки изменить значения элементов в цикле, как вот с разведчиками. Что для этого делать?

Первое пришедшее мне на ум решение было сформировать массив из указателей, вместо массива просто элементов. Но это означало бы изменть конструкторы, и весь код, который зависит от конструкторов. Это много.

Не может же язык быть столь опасен, что столь маленькое изменение может потребовать таких изменений.

Но рещение все же есть. Становится понятно зачем в операторе `range` два параметра

```go
for i,v :=range persons{
    persons[i].surname="***ЗАСЕКРЕЧЕНО***" //тогда это сработает
    fmt.Println(v)
}
```

Ну и уберем `v`

```go
for i:=range persons{
    persons[i].surname="***ЗАСЕКРЕЧЕНО***"
    // пробегаем только по индексам, не надо использовать значение
}
```

Вот так можно изменить значения элементов, среди который итерируем

## Иногда важно получше почитать условие

В некоторых задачах явно указано, что чисел не очень много: всего тысяча. А значит можно вместо мапы использовать массив!

А если подумать, то букв-то всего 26. Для многих буквенных задач мапа не нужна

## Проверяй тесты

Пару раз я супер жестко залип лишь оттого, что неправильно переписал тестовые примеры. И пускай мой алгоритм правильно был устроен, тесты все равно падали

## Тесты документируют код

В литкоде есть интересная особенность: все мои решения доступны в открытом виде. Это заставляет приводить код в подрядок. Мало ли кто посмотрит код, пусть ему будет приятно. Может мое решение ему будет самым полезным.

И вот у меня было не самое очевидное, но довольно элегантное решение.
`c-'0'` - что это значит? Так сразу и не скажешь, что результатом будет цифра, содежащаяся в `с`. То есть этот код превращает `c` в цифру.

Мне вообще нужно было проверить свою догадку, поэтому я написал простой тест

```go

func Test_SubtractChar(t *testing.T) {
    assert.Equal(t, int32(1), '1'-'0')
    assert.Equal(t, int32(9), '9'-'0')
}
```

И тут я вдруг понял, а ведь этот тест заодно и задокументировал код. Это же мега офигенно. Да, может чел так сходу и не поймет, но все-таки оставить какую-то подсказку хочется.

Ну и плюс я уверился в своем решении. Такое ощущение, что `литкод для меня это искусство написания тестов` :). Я все лучше и лучше пишу тесты, и задания даются мне проще и проще.

## Задание

Даже не представляю сколько бы я времени сэкономил, если бы чуть повнимательнее читал задания. Всякие ограничения, примеры. Не просто поверхостно глянул и побежал решать, а просто ещё одну секундочку полистать посмотреть.

## Трюк для связного списка

Часто проводить цикл по связному списку не удобно, нам нужно обрабатывать или первый или последний элемент как-то особо, поэтому мы можем воспользоваться трюком с -1 элементом. 

```go
dummy:=&List{0, head}

for {
    // ...
}

return dummy.Next // вернет head, или nil
```
