package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	Name        string
	Files       map[string]int
	Directories map[string]*Directory
	ParentDir   *Directory
	Size        int
}

func main() {
	fileScanner := readFile("input.txt")
	dir := buildDirectoryStructure(fileScanner)
	printTree(dir, 0)
	problemOne(dir)
	problemTwo(dir)

}

func problemTwo(dir Directory) {
	freeSpace := 70000000 - dir.Size
	neededSpace := 30000000 - freeSpace
	smallest := findSmallestDirOverSize(dir, neededSpace, dir.Size)
	fmt.Println("Smallest: ", smallest)
}

func findSmallestDirOverSize(dir Directory, minSize int, currentMin int) int {
	for _, element := range dir.Directories {
		currentMin = findSmallestDirOverSize(*element, minSize, currentMin)
		if element.Size >= minSize && element.Size < currentMin {
			return element.Size
		}
	}
	return currentMin
}

func buildDirectoryStructure(fileScanner *bufio.Scanner) Directory {
	dir := Directory{Directories: make(map[string]*Directory)}
	dir.Directories["/"] = &Directory{
		Name:        "/",
		Directories: make(map[string]*Directory),
		Files:       make(map[string]int),
	}
	currentDir := dir.Directories["/"]
	var args []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if isCommand(line) {
			args = strings.Split(line, " ")
			if args[1] == "cd" {
				switch args[2] {
				case "/": //move back to root dir
					currentDir = dir.Directories["/"]
				case "..": //move up 1 dir
					currentDir = currentDir.ParentDir
				default: //move to new dir
					if newDir, ok := currentDir.Directories[args[2]]; ok {
						//directory exists
						currentDir = newDir

					} else {
						newDir := Directory{
							Name:        args[2],
							Directories: make(map[string]*Directory),
							Files:       make(map[string]int),
							ParentDir:   currentDir,
						}
						currentDir = &newDir
					}
				}
			}
		} else { //always part of a list
			args = strings.Split(line, " ")
			if args[0] == "dir" {
				if _, ok := currentDir.Directories[args[1]]; !ok {
					newDir := Directory{
						Name:        args[1],
						Directories: make(map[string]*Directory),
						Files:       make(map[string]int),
						ParentDir:   currentDir,
					}
					currentDir.Directories[args[1]] = &newDir
				}
			} else { //file and size
				if _, ok := currentDir.Files[args[1]]; !ok {
					currentDir.Files[args[1]], _ = strconv.Atoi(args[0])
					currentDir.Size += currentDir.Files[args[1]]
				}
			}
		}
	}
	calcDirTotals(&dir, 0)
	return dir
}

func problemOne(dir Directory) {

	total := calcTotal(&dir, 0)
	fmt.Println("Grand Total: ", total)

}
func calcTotal(dir *Directory, total int) int {
	for _, element := range dir.Directories {
		total = calcTotal(element, total)
		if element.Size < 100000 {
			total += element.Size
		}
	}
	return total
}

func calcDirTotals(dir *Directory, total int) int {
	for _, element := range dir.Directories {
		dir.Size += calcDirTotals(element, total)
	}

	return dir.Size

}

func printTree(dir Directory, offset int) int {
	for index, element := range dir.Files {
		fmt.Printf("%*s\t%d\n", offset*10, index, element)
	}
	for index, element := range dir.Directories {
		fmt.Printf("%*s - %d\n", offset*10, index, element.Size)
		offset += 1
		offset = printTree(*element, offset)
	}
	offset -= 1
	return offset
}

func isCommand(line string) bool {
	firstChar := line[0:1]
	return firstChar == "$"
}

func readFile(fileName string) *bufio.Scanner {
	readFile, _ := os.Open(fileName)
	return bufio.NewScanner(readFile)
}
