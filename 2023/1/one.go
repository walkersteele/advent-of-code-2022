package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"unicode"
)

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var firstDigit rune
		var lastDigit rune
		for _, c := range line {
			if unicode.IsDigit(c) {
				if reflect.ValueOf(firstDigit).IsZero() {
					firstDigit = c
					lastDigit = c
				} else {
					lastDigit = c
				}
			}
		}
		number, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		total += number

		// fmt.Println("Current Total: ", total)
	}
	readFile.Close()
	fmt.Println(total)
}
