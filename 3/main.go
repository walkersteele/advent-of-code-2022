package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// problemOne()
	problemTwo()
}

func problemOne() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for fileScanner.Scan() {
		//get line input
		line := fileScanner.Text()
		items := strings.Split(line, "")
		sack1, sack2 := items[0:len(items)/2], items[len(items)/2:]
		match := '0'
		for _, item := range sack2 {
			if match != '0' {
				break
			}
			for _, item2 := range sack1 {
				if item == item2 {
					if item <= "a" { //is Uppercase
						total += int([]rune(item)[0]) - 38
					} else {
						total += int([]rune(item)[0]) - 96
					}
					match = []rune(item)[0]
					break
				}
			}
		}
	}
	readFile.Close()
	fmt.Println("Problem 1 Answer: ", total)
}

func problemTwo() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for true {
		var items [3][]string
		var matches []string
		for i := 0; i < 3; i++ {
			if fileScanner.Scan() {
				line := fileScanner.Text()
				items[i] = strings.Split(line, "")
				if i == 1 {
					for _, v := range items[i] {
						for _, v1 := range items[0] {
							if v == v1 {
								matches = append(matches, v)
							}
						}
					}
				} else if i == 2 {
					value := 0
					for _, v := range items[i] {
						for _, m := range matches {
							if v == m {
								if v <= "a" { //is Uppercase
									value = int([]rune(v)[0]) - 38
									total += value
								} else {
									value = int([]rune(v)[0]) - 96
									total += value
								}
								break
							}
						}
						if value != 0 {
							break
						}
					}
				}
			} else {
				fmt.Println("Problem 2 Answer: ", total)
				return
			}
		}

	}

}
