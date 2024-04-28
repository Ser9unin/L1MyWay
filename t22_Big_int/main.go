package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(1000000000), big.NewInt(1000000000)

	a.Mul(a, b) // умножение
	fmt.Println("Умножение", a)

	a, b = big.NewInt(1000000000), big.NewInt(1000000000)
	a.Div(a, b) // деление
	fmt.Println("Деление", a)

	a, b = big.NewInt(1000000000), big.NewInt(1000000000)
	a.Add(a, b) // сложение
	fmt.Println("Сложение", a)

	a, b = big.NewInt(1000000000), big.NewInt(1000000000)
	a.Sub(a, b) // вычитание
	fmt.Println("Вычитание", a)
}
