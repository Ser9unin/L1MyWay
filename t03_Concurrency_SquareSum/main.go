package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// с помощью атомарных операций,
// для простых вычислений подходит лучше чем мьютекс,
// так как энергоэффективнее
func AtomicWay(nums []int) (res int64) {
	wg := sync.WaitGroup{}

	for _, v := range nums {
		wg.Add(1)
		go func(v int) {
			//прибавляет к res квадрат числа из массива
			atomic.AddInt64(&(res), int64(v*v))
			wg.Done()
		}(v)
	}

	wg.Wait()
	return
}

// с помощью мьютексов
func MutexWay(nums []int) (res int) {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	// На каждое число из массива запускаем свою горутину
	wg.Add(len(nums))
	for _, v := range nums {

		v := v
		go func() {
			// прибавляет к res квадрат числа из массива, мьютекс блокирует действия над одним res в других горутинах
			// если не применить мьютекс в данном случае, каждая из созданных горутин может вязть для вычислений одино и тоже значение res и изменять его
			// конечный результат будет непредсказуем.
			mx.Lock()
			res += v * v
			mx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return
}

// метод с каналами и запуском нескольких горутин для обработки потока входных данных
// позволяет
func ChanWay(nums []int) (res int) {
	fmt.Println("Запустили расчет через канал")
	wg := sync.WaitGroup{}
	inputChan := make(chan int)
	squareChan := make(chan int)
	defer close(squareChan)

	// канал с пустой структурой для остановки работы горутин, работает с пустой структурой, которая ничего не принимает
	doneChan := make(chan struct{})

	//записываем данные в канал input через отдельную горутину, запись в канал в той же горутине где он создан недопустима
	wg.Add(1)
	go func(inputChan chan int, doneChan chan struct{}) {
		defer wg.Done()
		defer close(inputChan)
		for _, v := range nums {
			select {
			case inputChan <- v:
			case <-doneChan:
				fmt.Println("Горутина Input остановлена")
				return
			}
		}
	}(inputChan, doneChan)

	// запускаем 31 грутину
	// все вычисления будут произведены параллельно в 3 горутинах в данном случае,
	// но каждая получит свой элемент входного массива для обработки
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func(inputChan chan int, squareChan chan int, doneChan chan struct{}) {
			defer wg.Done()
			// чтение данных из канала позволяет не использовать мьютекс.
			// если doneChan закрыт то в select первым произойдёт чтение из него и выход из горутин
			fmt.Printf("Запущена Горутина %d\n", i)
			for {
				v, ok := <-inputChan
				if !ok {
					fmt.Printf("Нечего вычислять горутине %d \n", i)
				}
				v = v * v
				fmt.Printf("Вычисляет горутина %d\n", i)
				select {
				case squareChan <- v:
				case <-doneChan:
					fmt.Printf("остановлена Горутина %d\n", i)
					return
				}
			}
		}(inputChan, squareChan, doneChan)
	}

	for i := 0; i < len(nums); i++ {
		v, ok := <-squareChan
		if !ok {

			break
		}
		res += v
	}

	// закрытие done канала завершает запущенные горутины
	close(doneChan)
	wg.Wait()

	return res
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	fmt.Println("Результат на Atomic", AtomicWay(nums))
	fmt.Println("Результат на Mutex", MutexWay(nums))
	fmt.Println("Результат на Chan", ChanWay(nums))
}
