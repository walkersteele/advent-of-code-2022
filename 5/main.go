package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// problemOne(readFile("input.txt"))
	// problemTwo(readFile("input_test.txt"))
	problemTwo(readFile("input.txt"))

}

func problemOne(fileScanner *bufio.Scanner) {

	stacks := readDrawing(fileScanner)
	fmt.Println(stacks)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		re := regexp.MustCompile(`\d+`)
		chars := re.FindAllString(line, -1)
		fmt.Println(chars)
		var nums [3]int
		for k, v := range chars {
			nums[k], _ = strconv.Atoi(v)
		}
		for i := 0; i < nums[0]; i++ {
			lastIndex := len(stacks[nums[1]-1]) - 1
			toMove := stacks[nums[1]-1][lastIndex]
			stacks[nums[1]-1] = append(stacks[nums[1]-1][:lastIndex], stacks[nums[1]-1][lastIndex+1:]...) //remove from stack
			stacks[nums[2]-1] = append(stacks[nums[2]-1], toMove)                                         //append to new location
		}
	}
	answerString := ""
	for _, stack := range stacks {
		answerString += stack[len(stack)-1]
	}
	fmt.Println(answerString)

}

func problemTwo(fileScanner *bufio.Scanner) {

	stacks := readDrawing(fileScanner)
	fmt.Println(stacks)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		re := regexp.MustCompile(`\d+`)
		chars := re.FindAllString(line, -1)
		fmt.Println(chars)
		var nums [3]int
		for k, v := range chars {
			nums[k], _ = strconv.Atoi(v)
		}

		lastIndex := len(stacks[nums[1]-1]) - 1
		toMove := stacks[nums[1]-1][lastIndex-nums[0]+1:]
		stacks[nums[1]-1] = stacks[nums[1]-1][:lastIndex-nums[0]+1]
		stacks[nums[2]-1] = append(stacks[nums[2]-1], toMove...) //append to new location

	}
	answerString := ""
	for _, stack := range stacks {
		answerString += stack[len(stack)-1]
	}
	fmt.Println(answerString)

}

func readDrawing(fileScanner *bufio.Scanner) [][]string {
	var numStacks int
	var stacks [][]string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !strings.Contains(line, "[") {
			return stacks
		}
		if numStacks == 0 {
			numStacks = numberOfStacks(len(line))
			stacks = make([][]string, numStacks)
		}
		r := strings.Split(line, "")
		for i := 0; i < numStacks; i++ {
			if r[4*i+1] != " " {
				stacks[i] = append([]string{r[4*i+1]}, stacks[i]...)
			}

		}
	}
	return stacks
}

func numberOfStacks(length int) int {
	//3n + (n-1) = length
	return (length + 1) / 4
}

func readFile(fileName string) *bufio.Scanner {
	readFile, _ := os.Open(fileName)
	return bufio.NewScanner(readFile)
}
