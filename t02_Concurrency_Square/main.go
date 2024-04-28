package main

import (
	"fmt"
	"sync"
)

// количество воркеров
var workers = 4

// воркер считывает число из канала и выводит его квадрат
func startWorker(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		fmt.Println("Вывод функции Square с воркерами", n*n)
	}
}

// вариант с использование воркеров
func SquareWithWorkers(nums []int) {
	wg := &sync.WaitGroup{}
	wg.Add(workers)
	input := make(chan int)

	// запускаем воркеров
	for i := 0; i < workers; i++ {
		go startWorker(input, wg)
	}

	// отправляем данные в канал
	for _, v := range nums {
		input <- v
	}

	// закрываем канал, и ждем завершения работы воркеров
	close(input)
	wg.Wait()
}

// вариант с запуском отдельной горутины на элемент
func Square(nums []int) {
	wg := &sync.WaitGroup{}

	for _, v := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println("Вывод функции Square с WaitGroup", num*num)
		}(v)
	}

	wg.Wait()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	Square(nums)
	SquareWithWorkers(nums)
}
