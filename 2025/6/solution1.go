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
	var numbers [][]int

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		if buff[0] == '*' || buff[0] == '+' {
			break
		}
		var row []int
		pendingNumber := ""
		for _, val := range buff {
			if val == ' ' && pendingNumber != "" {
				num, _ := strconv.Atoi(pendingNumber)
				row = append(row, num)
				pendingNumber = ""
			} else if val == ' ' {
				continue
			} else {
				pendingNumber = pendingNumber + string(val)
			}
		}
		if pendingNumber != "" {
			num, _ := strconv.Atoi(pendingNumber)
			row = append(row, num)
		}

		numbers = append(numbers, row)
	}
	fmt.Println(numbers)
	buff := fileScanner.Text()
	c := 0
	total := 0

	for _, val := range buff {
		if val == ' ' {
			continue
		} else if val == '+' {
			subtotal := 0
			for y := 0; y < len(numbers); y++ {
				fmt.Println(y, c)
				subtotal += numbers[y][c]
			}
			fmt.Println(subtotal, c)
			total += subtotal
			c++
		} else if val == '*' {
			subtotal := 1
			for y := 0; y < len(numbers); y++ {
				fmt.Println(y, c)
				subtotal = numbers[y][c] * subtotal
			}
			fmt.Println(subtotal, c)
			total += subtotal
			c++
		}
	}

	fmt.Println(total)
}
