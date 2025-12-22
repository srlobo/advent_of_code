package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		defer readFile.Close()
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	state := 50
	count := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()

		actualNum, _ := strconv.Atoi(buff[1:])

		if buff[0] == 'L' { // Left
			if actualNum >= state && state != 0 {
				actualNum = actualNum - state
				state = 0
				count++
			}
			count += actualNum / 100
			actualNum = actualNum % 100
			state = state - actualNum
		} else { // Right
			if actualNum >= (100-state) && state != 0 {
				actualNum = actualNum - (100 - state)
				state = 0
				count++
			}
			count += actualNum / 100
			actualNum = actualNum % 100
			state = state + actualNum
		}

		if state < 0 {
			state = 100 + state
		}

		fmt.Println(buff, state, count)

	}
	fmt.Println(count)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
