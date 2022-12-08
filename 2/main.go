package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for fileScanner.Scan() {
		//get line input
		line := fileScanner.Text()
		//Split into array of words
		actions := strings.Fields(line)
		switch actions[0] {
		case "A": //rock
			switch actions[1] {
			case "X":
				total += 4 //1 + 3
			case "Y":
				total += 8 //2 + 6
			case "Z":
				total += 3 //3 + 0
			}
		case "B": //paper
			switch actions[1] {
			case "X":
				total += 1 //1 + 0
			case "Y":
				total += 5 //2 + 3
			case "Z":
				total += 9 //3 + 6
			}
		case "C": //scisorrs
			switch actions[1] {
			case "X":
				total += 7 //1 + 6
			case "Y":
				total += 2 //2 + 0
			case "Z":
				total += 6 //3 + 3
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
	for fileScanner.Scan() {
		//get line input
		line := fileScanner.Text()
		//Split into array of words
		actions := strings.Fields(line)
		switch actions[0] {
		case "A": //rock
			switch actions[1] {
			case "X": //lose
				total += 3 //0 + 3
			case "Y": //draw
				total += 4 //3 + 1
			case "Z": //win
				total += 8 //6 + 2
			}
		case "B": //paper
			switch actions[1] {
			case "X":
				total += 1 //0 + 1
			case "Y":
				total += 5 //3 + 2
			case "Z":
				total += 9 //6 + 3
			}
		case "C": //scisorrs
			switch actions[1] {
			case "X":
				total += 2 //0 + 2
			case "Y":
				total += 6 //3 + 3
			case "Z":
				total += 7 //6 + 1
			}
		}
	}
	readFile.Close()
	fmt.Println("Problem 2 Answer: ", total)
}
