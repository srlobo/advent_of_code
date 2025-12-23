package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	lower int
	upper int
}

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

	var freshIngredientsIDRanges []idRange

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if buff == "" {
			break
		}
		numbers := strings.Split(buff, "-")
		lower, _ := strconv.Atoi(numbers[0])
		upper, _ := strconv.Atoi(numbers[1])

		r := idRange{lower, upper}
		freshIngredientsIDRanges = append(freshIngredientsIDRanges, r)
	}

	printIDRange(freshIngredientsIDRanges)

	count := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()

		number, _ := strconv.Atoi(buff)

		if idInIDRange(freshIngredientsIDRanges, number) {
			count++
		}


	}

	fmt.Println(count)

}

func printIDRange(ingredientsIDRanges []idRange) {
	for _, r := range ingredientsIDRanges {
		fmt.Println(r.lower, "-", r.upper)
	}
}

func idInIDRange(ingredientsIDRanges []idRange, element int) bool {
	for _, r := range ingredientsIDRanges {
		if element >= r.lower && element <= r.upper {
			return true
		}
	}
	return false
}
