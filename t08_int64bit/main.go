package main

import (
	"fmt"
)

func setBit(num int64, bitIndex int) int64 {
	return num | 1<<bitIndex
}

func main() {
	var numMaxInt64 int64 = 9223372036854775807 // Пример числа для демонстрации, выбрал максимальное значение для int64
	bitIndex := 63                              // Индекс бита, который хотим установить, выбрал 64-й что бы перевести в диапазон отрицательных чисел

	newNum := setBit(numMaxInt64, bitIndex)
	fmt.Println("Исходное число:", numMaxInt64)
	fmt.Println("Число после установки бита:", newNum)

	var num int64 = 8
	bitIndex = 1
	newNum = setBit(num, bitIndex)
	fmt.Println("Исходное число:", num)
	fmt.Println("Число после установки бита:", newNum)
}
