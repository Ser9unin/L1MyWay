package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// воркер считывает число из канала и выводит его
func startWorker(ctx context.Context, wg *sync.WaitGroup, num int, in <-chan int) {
	defer wg.Done()
	var n int
	for {
		select {
		case n = <-in:
			fmt.Printf("worker № %d, got: %d\n", num, n)
		case <-ctx.Done():
			fmt.Printf("exiting from worker №%d\n", num)
			return
		}
	}
}

func main() {
	var workersNum int
	// считываем количество воркеров из аргументов терминала
	workersNum, err := strconv.Atoi((os.Args[1]))
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}

	// создали входной канал
	input := make(chan int)
	defer close(input)

	// метод оснатновки через контекст выбран так как в более масштабной программе с разными модулями он бы позволил прервать выполнение всех функций
	// куда переадан контекст. Так же метод с done каналом был реализован в задаче 3, решил не дублировать его
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем workersNum воркеров и передаём в них контекст и канал из которого будем читать
	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go startWorker(ctx, wg, i+1, input)
	}

	// запускаем запись случайных чисел в канал с определенной периодичностью
	// завершается по сигналу контекста
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				a := rand.Int()
				input <- a
				fmt.Printf("writer sent: %d\n", a)
			case <-ctx.Done():
				fmt.Println("exiting from writer")
				return
			}
		}
	}()

	// по нажатию ctrl+c отправляем сигнал завершения всем горутинам через контекст
	// и ждем их завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\ngot interrupt signal")
	cancel()
	wg.Wait()
}
