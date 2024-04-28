package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type goroutines struct {
	wg
}

type wg struct {
	sync.WaitGroup
}

// горутина проверяет закрылся ли канал пытаясь вычитать из него данные
// если вернулось стандартное значение (не ок), то завершаем функцию
func (g *goroutines) channelClose(ch <-chan int) {
	defer g.Done()
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Вариант 1: Горутина пыталась прочитать из закрытого канала")
			return
		}
		fmt.Printf("Вариант 1: из канала получено значение: %d\n", v)
	}
}

// range завершится сам, как только канал закроется, в принципе примерно тоже что в channelClose
func (g *goroutines) channelCloseRange(ch <-chan int) {
	defer g.Done()
	for v := range ch {
		fmt.Printf("Вариант 2: из канала получено значение: %d\n", v)
	}
	fmt.Println("Вариант 2: Горутина пыталась прочитать из закрытого канала")
}

// используем отдельный Done канал для завершения горутины
func (g *goroutines) channelCloseDoneChannel(ch <-chan int, stop <-chan struct{}) {
	defer g.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf("Вариант 3: из канала получено значение: %d\n", v)
		case <-stop:
			fmt.Println("Вариант 3: Done канал закрыт, горутина остановлена")
			return
		}
	}
}

// используем сигнал из контекста для завершения горутины
func (g *goroutines) channelCloseContext(ctx context.Context, ch <-chan int) {
	defer g.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf("Вариант 4: из канала получено значение: %d\n", v)
		case <-ctx.Done():
			fmt.Println("Вариант 4: Получен сигнал на остановку контекста, горутина остановлена")
			return
		}
	}
}

// запись случайных чисел в канал
// завершается по сигналу контекста, контекст завершается в конце main,
// что бы в остаьлных вариантах отработали сигналы из Done канала
func (g *goroutines) producer(ctx context.Context, ch chan int) {
	defer g.Done()

	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			a := rand.Int()
			ch <- a
			fmt.Printf("Отправили в канал: %d\n", a)
		case <-ctx.Done():
			fmt.Println("ЗАПИСЬ В КАНАЛ ОСТАНОВЛЕНА")
			close(ch)
			ticker.Stop()
			return
		}
	}
}

func main() {
	g := goroutines{}

	// Вариант 1: проверили открыт ли канал через v, ok := <-ch
	fmt.Println("Вариант 1")
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	g.Add(2)
	go g.producer(ctx, ch)
	go g.channelClose(ch)

	time.Sleep(2 * time.Second)
	cancel() //останавливаем продьюсера, продьюсер так же закрывает канал ch
	g.Wait()

	// Вариант 2: читаем данные из канала через range
	fmt.Println("\nВариант 2: читаем данные из канала через range")
	ch2 := make(chan int)
	ctx, cancel = context.WithCancel(context.Background())

	g.Add(2)
	go g.producer(ctx, ch2)
	go g.channelCloseRange(ch2)

	time.Sleep(2 * time.Second)
	cancel() //останавливаем продьюсера, продьюсер так же закрывает канал ch1
	g.Wait()

	// Вариант 3: С закрытием Done канала
	fmt.Println("\nВариант 3: С закрытием Done канала")
	ch3 := make(chan int, 1)
	done := make(chan struct{})
	ctx, cancel = context.WithCancel(context.Background())

	g.Add(2)
	go g.producer(ctx, ch3)
	go g.channelCloseDoneChannel(ch3, done)

	time.Sleep(2 * time.Second)
	close(done)
	cancel() //останавливаем продьюсера, продьюсер так же закрывает канал ch3
	g.Wait()

	// Вариант 4: С контекстом
	fmt.Println("\nВариант 4: С контекстом")
	ctx, cancel = context.WithCancel(context.Background())
	ch4 := make(chan int, 1)

	g.wg.Add(2)
	go g.producer(ctx, ch4)
	go g.channelCloseContext(ctx, ch4)

	time.Sleep(2 * time.Second)
	cancel() //останавливаем продьюсера, продьюсер так же закрывает канал ch4
	g.Wait()
}
