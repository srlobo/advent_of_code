package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	sort.Ints(elfLoad)

	sumThreeLastElements := elfLoad[len(elfLoad)-1] + elfLoad[len(elfLoad)-2] + elfLoad[len(elfLoad)-3]

	fmt.Println(sumThreeLastElements)
}
