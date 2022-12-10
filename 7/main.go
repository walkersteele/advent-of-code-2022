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
	problemOne(fileScanner)
	// problemTwo(readFile("input.txt"))

}

func problemOne(fileScanner *bufio.Scanner) {
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

	// printTree(dir, 0)
	// fmt.Println()
	// totals := make(map[string]int)
	// _, totals = calcTotals(dir, "", totals)
	// keys := make([]string, 0, len(totals))
	// for k := range totals {
	// 	keys = append(keys, k)
	// }
	// sort.SliceStable(keys, func(i, j int) bool {
	// 	return totals[keys[i]] < totals[keys[j]]
	// })
	// total := 0
	// for _, k := range keys {
	// 	// fmt.Printf("%s: %d\n", k, totals[k])
	// 	if totals[k] <= 100000 {
	// 		total += totals[k]
	// 	}

	// }
	// fmt.Println("Total: ", total)
	calcDirTotals(&dir, 0)
	total := calcTotal(&dir, 0)
	fmt.Println("Grand Total: ", total)
	// printTree(dir, 0)

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
