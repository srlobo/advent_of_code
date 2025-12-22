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

	total := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		firstNumber := 0
		secondNumber := 0
		for i := 0; i < len(buff)-1; i++ {
			current, _ := strconv.Atoi(string(buff[i]))
			if current > firstNumber {
				firstNumber = current
				secondNumber = 0
			} else if current > secondNumber {
				secondNumber = current
			}
		}
		current, _ := strconv.Atoi(string(buff[len(buff)-1]))
		if secondNumber < current {
			secondNumber = current
		}

		fmt.Println(firstNumber*10 + secondNumber)

		total += firstNumber*10 + secondNumber
	}
	fmt.Println(total)
}
