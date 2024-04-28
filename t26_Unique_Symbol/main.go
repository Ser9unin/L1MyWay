package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		sample string
		unique bool = true
	)
	uniqueSet := make(map[rune]struct{})

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// читаем до конца строки
	sample = scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// переводим строку в один регистр
	sample = strings.ToLower(sample)

	for _, char := range sample {
		if _, ok := uniqueSet[char]; !ok {
			uniqueSet[char] = struct{}{}
		} else {
			unique = false
		}
	}

	fmt.Println(unique)
}
