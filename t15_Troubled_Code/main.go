/*
Данный фрагмент кода может привести к нескольким негативным последствиям:
Использование глобальной переменной: В данном коде используется глобальная переменная justString,
что может привести к неожиданным результатам, если эта переменная будет изменена в других частях программы.
var justString string так же в данном случае justString нигде не используется

func someFunc() {
Создание огромной строки: Функция createHugeString создает строку размером 1024 байта, что может вызвать проблемы с памятью, особенно если таких строк создается много.
  v := createHugeString(1 << 10)
  Неэффективное использование памяти: Строки в Go являются неизменяемыми, поэтому при присвоении части строки justString = v[:100] создается новая строка,
  что может привести к неэффективному использованию памяти.
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func createHugeString(size int) string {
	// используем буфер для эффективной конкатенации строк
	var b strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "界")
	}

	return b.String()
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) //создаём строку 1024 байт
	// руна может занимать не один байт
	fmt.Println(utf8.RuneLen('界'))

	// в данном случае мы срезаем по количеству байт, а не по количеству рун,
	// это приведет к непредвиденному результату, какой-то символ может быть срезан посередине
	justString = v[:100]

	// преобразовываем строку в слайс рун
	r := []rune(v)
	// в даннам случае мы срезаем по количеству рун
	justString = string(r[:100])

	// для примера выведем justString в stdout
	fmt.Println(justString)
}

func main() {
	someFunc()
}
