package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	setTemp := make(map[int][]float64)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal("Ошибка ввода", err)
	}

	line := scanner.Text()
	line = strings.ReplaceAll(line, ",", "")
	lines := strings.Fields(line)

	for _, line := range lines {
		temp, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
		}

		lim := int(temp/10) * 10

		setTemp[lim] = append(setTemp[lim], temp)
		lim = 0
	}
	fmt.Println(setTemp)
}
