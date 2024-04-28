package main

import (
	"fmt"
	"sort"
)

func main() {
	// Сортировка встроенноей функцией библиотекой sort, зависит от длины массива
	// если длина слайса 0 или 1 получим панику (индекс вне диапазона)
	//
	unsorted := []int{10, 7, 8, 9, 1, 5}
	sort.Slice(unsorted, func(i, j int) bool { return unsorted[i] < unsorted[j] })
	fmt.Println("Sorted array:", unsorted)

	// сортировка написанная руками

}
