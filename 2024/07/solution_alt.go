package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	total1 := 0
	total2 := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		target, _ := strconv.Atoi(strings.Split(buff, ":")[0])
		seriesStr := strings.Split(strings.Split(buff, ":")[1], " ")
		var series []int
		for i := 0; i < len(seriesStr); i++ {
			num, _ := strconv.Atoi(seriesStr[i])
			series = append(series, num)
		}
		series = series[1:]
		if calculateRecursive1(target, series[0], series[1:]) {
			total1 += target
		}

		if calculateRecursive2(target, series[0], series[1:]) {
			total2 += target
		}

		// fmt.Println("-----")
	}
	fmt.Println("total1", total1)
	fmt.Println("total2", total2)
}

func calculateRecursive1(target, actual int, series []int) bool {
	// fmt.Println("target", target, "actual", actual, "series", series)
	if actual > target {
		return false
	}
	if len(series) == 0 {
		return actual == target
	}
	new := actual + series[0]
	if calculateRecursive1(target, new, series[1:]) {
		return true
	}
	new = actual * series[0]
	if calculateRecursive1(target, new, series[1:]) {
		return true
	}

	return false
}

func calculateRecursive2(target, actual int, series []int) bool {
	// fmt.Println("target", target, "actual", actual, "series", series)
	if actual > target {
		return false
	}
	if len(series) == 0 {
		return actual == target
	}
	new := actual + series[0]
	if calculateRecursive2(target, new, series[1:]) {
		return true
	}
	new = actual * series[0]
	if calculateRecursive2(target, new, series[1:]) {
		return true
	}
	new, _ = strconv.Atoi(strconv.Itoa(actual) + strconv.Itoa(series[0]))
	if calculateRecursive2(target, new, series[1:]) {
		return true
	}

	return false
}
