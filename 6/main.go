package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// problemOne(readFile("input.txt"))
	// problemOne(readFile("input_test.txt"))
	problemTwo(readFile("input.txt"))

}

func problemTwo(fileScanner *bufio.Scanner) {
	fileScanner.Scan()
	line := []byte(fileScanner.Text())
	var search []byte
	for i := 0; i < len(line); i++ {
		search = line[i : i+14]
		dict := make(map[byte]int)
		unique := true
		for _, v := range search {
			if _, ok := dict[v]; !ok {
				//unique
				dict[v] = 1
			} else {
				unique = false
				break
			}
		}
		if unique {
			fmt.Println("Unique: ", string(search))
			fmt.Println("At index: ", i+14)
			return
		}

	}
}

func problemOne(fileScanner *bufio.Scanner) {
	fileScanner.Scan()
	line := []byte(fileScanner.Text())
	var search []byte
	for i := 0; i < len(line); i++ {
		search = line[i : i+4]
		dict := make(map[byte]int)
		unique := true
		for _, v := range search {
			if _, ok := dict[v]; !ok {
				//unique
				dict[v] = 1
			} else {
				unique = false
			}
		}
		if unique {
			fmt.Println("Unique: ", string(search))
			fmt.Println("At index: ", i+4)
			return
		}

	}
}

func readFile(fileName string) *bufio.Scanner {
	readFile, _ := os.Open(fileName)
	return bufio.NewScanner(readFile)
}
