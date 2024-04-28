package main

import (
	"fmt"
	"time"
)

// Ожидает время "duration" и выводит результат
func SleepAfter(duration int) {
	fmt.Println("Время из SleepAfter", <-time.After(time.Duration(duration)*time.Second))
}

// Ожидает время "duration" и выводит результат
func SleepTick(duration int) {
	fmt.Println("Время из SleepTick", <-time.Tick(time.Duration(duration)*time.Second))
}

// По сути тоже что и SleepTick
func TickerC(duration int) {
	ticker := time.NewTicker(time.Duration(duration) * time.Second)
	fmt.Println("Время из Ticker", <-ticker.C)
}

func main() {
	fmt.Println(time.Now())
	duration := 2

	SleepAfter(duration)
	SleepTick(duration)
	TickerC(duration)
}
