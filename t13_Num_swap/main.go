package main

import "fmt"

func main() {
	// вариант 1
	a, b := 10, 20
	fmt.Printf("Вариант 1: a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("Вариант 1: Результат: a = %d, b = %d\n", a, b)

	// вариант 2 Сложение вычитание
	a, b = -10, 10
	fmt.Printf("Вариант 2: Сложение вычитание: a = %d, b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("Вариант 2: Результат: a = %d, b = %d\n", a, b)

	// вариант 3 XOR
	a, b = 10, 20 // a = 0000 1010, b = 0001 0100

	fmt.Printf("Вариант 3: XOR: a = %d, b = %d\n", a, b)
	a = a ^ b // в a записываем 0001 1110
	b = b ^ a // в b записываем 0000 1010
	a = a ^ b // в a записываем 0001 0100

	fmt.Printf("Вариант 3: Результат: a = %d, b = %d\n", a, b)
}
