package main

import (
	"fmt"
	"sync"
)

type CounterMutex struct {
	sync.Mutex
	val int
}

// конструктор структуры CounterMutex
func NewCounter() *CounterMutex {
	return &CounterMutex{
		val: 0,
	}
}

func (c *CounterMutex) increment() {
	c.Lock()
	c.val++
	c.Unlock()
}

// с помощью мьютексов
func MutexWay(counter *CounterMutex) int {
	wg := sync.WaitGroup{}
	// На каждое число из массива запускаем свою горутину
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// увеличивает counter, мьютекс блокирует действия над одним counter в других горутинах
			// если не применить мьютекс в данном случае, каждая из созданных горутин может взять для вычислений одно и тоже значение counter и изменять его
			// конечный результат будет непредсказуем.
			counter.increment()
			wg.Done()
		}()
	}

	wg.Wait()
	return counter.val
}

func main() {
	counterMutex := NewCounter()

	fmt.Println("Результат на Mutex", MutexWay(counterMutex))
}
