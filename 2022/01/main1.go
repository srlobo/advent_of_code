package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var elfLoad []int
	currentSum := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if buff == "" {
			elfLoad = append(elfLoad, currentSum)
			currentSum = 0
		} else {
			intVar, err := strconv.Atoi(buff)
			check(err)
			currentSum += intVar
		}
	}
	elfLoad = append(elfLoad, currentSum)
	currentSum = 0

	maxValue := 0
	for _, elf := range elfLoad {
		if elf > maxValue {
			maxValue = elf
		}
	}

	fmt.Println(maxValue)
}
