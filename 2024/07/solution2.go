package main

import (
	"bufio"
	"fmt"
	"math"
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

	total := 0
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
		fmt.Println(target, series)

		length := len(series) - 1
		for i := 0; i < int(math.Pow(3, float64(length))); i++ {
			combinationStr := strconv.FormatUint(uint64(i), 3)
			if len(combinationStr) < length {
				combinationStr = strings.Repeat("0", length-len(combinationStr)) + combinationStr
			}
			// fmt.Println("combinationStr", combinationStr)
			res := calculate(target, series, combinationStr)
			// fmt.Println(res)
			if res {
				fmt.Println("found - ", combinationStr)
				total += target
				break
			}
		}
		fmt.Println("-----")
	}
	fmt.Println("total", total)
}

func calculate(target int, series []int, combination string) bool {
	total := series[0]
	for i, r := range combination {
		if r == '0' {
			total = total + series[i+1]
		} else if r == '1' {
			total = total * series[i+1]
		} else if r == '2' {
			num, _ := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(series[i+1]))
			// fmt.Println(total, series[i+1], "->", num)
			total = num
		}

		if total > target {
			return false
		}
	}
	if total == target {
		return true
	}

	return false
}
