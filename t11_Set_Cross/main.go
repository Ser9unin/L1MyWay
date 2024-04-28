package main

import "fmt"

func main() {
	set1 := []int{2, 12, 3, 17, 25, 4, 8, 7, 6, 5}
	set2 := []int{13, 15, 17, 1, 4, 12, 10}
	setCross := make(map[int]int)
	var resSet []int

	// проходим по первому массиву и складываем всё в map
	for _, v := range set1 {
		setCross[v] = 1
	}

	// проходим по второму массиву, если нашли значение в map то добавляем его в пересечение resSet,
	// и увеличиваем значение в map для этого ключа что бы не добавлять дублирующиеся значения
	for _, v := range set2 {
		if setCross[v] == 1 {
			resSet = append(resSet, v)
			setCross[v]++
		}
	}

	fmt.Println(resSet)
}
