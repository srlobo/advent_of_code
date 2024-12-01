package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	numbers := make(map[int]int)
	leftList := []int{}

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		regex := regexp.MustCompile(" +")
		result := regex.Split(buff, -1)
		left, _ := strconv.Atoi(result[0])
		right, _ := strconv.Atoi(result[1])

		leftList = append(leftList, left)
		numbers[right] = numbers[right] + 1
	}

	fmt.Println(numbers)
	total := 0
	for _, left := range leftList {
		total += left * numbers[left]
	}
	fmt.Println(total)
}
