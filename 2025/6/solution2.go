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
	var numbers []string

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)
		numbers = append(numbers, buff)
	}
	fmt.Println(numbers)
	total := 0

	var columns []int
	var operation byte
	for x := len(numbers[0]) - 1; x >= 0; x-- {
		var column []byte
		finishColumn := true
		for y := 0; y < len(numbers); y++ {
			c := numbers[y][x]
			if c == ' ' {
				continue
			} else if c == '+' || c == '*' {
				operation = c
				finishColumn = false
			} else {
				column = append(column, c)
				finishColumn = false
			}
		}
		parsedColumn, _ := strconv.Atoi(string(column))
		columns = append(columns, parsedColumn)
		if finishColumn {
			fmt.Println(columns, string(operation))
			if operation == '*' {
				subtotal := 1
				for _, num := range columns {
					if num == 0 {
						continue
					}
					subtotal = subtotal * num
				}
				total += subtotal
				fmt.Println(columns, operation, subtotal)
			} else if operation == '+' {
				subtotal := 0
				for _, num := range columns {
					subtotal = subtotal + num
				}
				total += subtotal
				fmt.Println(columns, operation, subtotal)
			} else {
				fmt.Println("Operation error: ", string(operation))
			}
			columns = nil
			operation = ' '

		}
	}
	fmt.Println(columns, string(operation))
	if operation == '*' {
		subtotal := 1
		for _, num := range columns {
			if num == 0 {
				continue
			}
			subtotal = subtotal * num
		}
		total += subtotal
		fmt.Println(columns, operation, subtotal)
	} else if operation == '+' {
		subtotal := 0
		for _, num := range columns {
			subtotal = subtotal + num
		}
		total += subtotal
		fmt.Println(columns, operation, subtotal)
	} else {
		fmt.Println("Operation error: ", string(operation))
	}

	fmt.Println(total)
}
