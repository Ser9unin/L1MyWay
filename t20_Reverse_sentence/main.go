package main

import (
	"fmt"
	"strings"
)

func reverseWithSplit(sample string) {
	sampleSplit := strings.Split(sample, " ")
	var outString string
	// конкатенация через + не эффективный способ под новую строку выделяется новая память
	for i := len(sampleSplit) - 1; 0 <= i; i-- {
		outString += sampleSplit[i]
		if i != 0 {
			outString += " "
		}
	}
	fmt.Println("Вариант 1", outString)
}

func reverseWithFields(sample string) {
	sampleSplit := strings.Fields(sample)
	var outString string
	outString += strings.Join(sampleSplit, " ") //объединяет строки через builder пишет в byte buffer
	fmt.Println("Вариант 2", outString)
}

func main() {
	sample := "snow dog sun"

	reverseWithSplit(sample)
	reverseWithFields(sample)
}
