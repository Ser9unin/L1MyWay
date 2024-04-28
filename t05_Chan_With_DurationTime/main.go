package main

/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.*/

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	//читаем время работы из терминала
	seconds, err := strconv.Atoi((os.Args[1]))
	if err != nil {
		panic(err)
	}

	// создаём done канал, закроем его по истечении времени
	doneChannel := make(chan struct{})

	// запускаем таймер после которого закроем done канал и остановим передачу данныйх в каналы
	timer := time.AfterFunc(time.Second*time.Duration(seconds), func() {
		close(doneChannel)
	})
	defer timer.Stop()

	// канал в который будем писать и из него же читать
	channel := make(chan int)
	defer close(channel)

	data := 500

	// так как в горутине канал, она запустится и будет работать без Wait Group
	go func() {
		// так же бесконечно читаем данные из канала
		for {
			data := <-channel //reading "data" from channel
			fmt.Println(data)
		}
	}()

	// в бесконечном цикле пишем и читаем данные
	for {
		time.Sleep(300 * time.Millisecond) // поставим таймер что бы не слишком стремительно увеличивать счетчик
		data += 1
		select {
		case channel <- data: //writing "data" to channel
		case <-doneChannel: // получили сигнал о закрытии done канала
			fmt.Println("Finish")
			return
		}
	}
}
