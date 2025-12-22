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
		mul := 1
		buff := fileScanner.Text()
		if buff[0] == 'L' {
			mul = -1
		}

		actualNum, _ := strconv.Atoi(buff[1:])

		state = (state + (actualNum * mul)) % 100
		if state == 0 {
			count++
		}
		if state < 0 {
			state = state + 100
		}
		fmt.Println(buff, state)

	}
	fmt.Println(count)
}
