package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	problemOne(readFile("input.txt"))
	problemTwo(readFile("input.txt"))

}

func problemOne(fileScanner *bufio.Scanner) {
	count := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		groups := strings.Split(line, ",")
		group1 := convertStringsToInts(strings.Split(groups[0], "-"))
		group2 := convertStringsToInts(strings.Split(groups[1], "-"))
		if isAContainedInB(group1, group2) || isAContainedInB(group2, group1) {
			count += 1
		}
	}
	fmt.Println("Answer to problem 1 is: ", count)
}

func problemTwo(fileScanner *bufio.Scanner) {
	count := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		groups := strings.Split(line, ",")
		group1 := convertStringsToInts(strings.Split(groups[0], "-"))
		group2 := convertStringsToInts(strings.Split(groups[1], "-"))
		if isOverlapping(group1, group2) {
			count += 1
		}
	}
	fmt.Println("Answer to problem 2 is: ", count)
}

func isOverlapping(a []int, b []int) bool {
	if a[1] < b[0] || b[1] < a[0] {
		return false
	}
	return true
}

func isAContainedInB(a []int, b []int) bool {
	return b[0] <= a[0] && b[1] >= a[1]
}

func convertStringsToInts(s []string) []int {
	si := make([]int, 0, len(s))
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		si = append(si, i)
	}
	return si
}

func readFile(fileName string) *bufio.Scanner {
	readFile, _ := os.Open(fileName)
	return bufio.NewScanner(readFile)
}
