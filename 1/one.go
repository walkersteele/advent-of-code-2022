package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, _ := os.Open("input.txt");
	fileScanner := bufio.NewScanner(readFile)
	max, total := 0, 0
	for fileScanner.Scan() {
		line := fileScanner.Text();
		intLine, err := strconv.Atoi(line)
		if(err == nil){
			total += intLine
		}else{
			if total > max {
				max = total
			}
			total = 0
		}

	}
	readFile.Close();
	fmt.Println(max);
}
