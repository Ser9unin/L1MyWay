package main

import "fmt"

func main() {
	sample := "главрыба 日本語 test"

	// так как символы могут быть юникод то используем массив рун, а не массив байт
	runeSample := []rune(sample)
	var outString string
	for i := len(runeSample) - 1; 0 <= i; i-- {

		outString += string(runeSample[i])
	}

	fmt.Println(sample, "-", outString)
}
