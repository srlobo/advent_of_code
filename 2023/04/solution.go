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

	empty := struct{}{}
	total := 0

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)

		winning_numbers := map[int]struct{}{}

		buff = strings.Split(buff, ": ")[1]
		buff2 := strings.Split(buff, " | ")
		winning_numbers_str := buff2[0]
		my_numbers_str := buff2[1]

		for _, n := range strings.Split(winning_numbers_str, " ") {
			number, _ := strconv.Atoi(n)
			if number == 0 {
				continue // Weird edge case for the way I'm splicing the string :facepalm:
			}
			winning_numbers[number] = empty
		}
		fmt.Println(winning_numbers)

		card_total := 0
		for _, n := range strings.Split(my_numbers_str, " ") {
			number, _ := strconv.Atoi(n)
			if number == 0 {
				continue // Weird edge case for the way I'm splicing the string :facepalm:
			}
			if _, ok := winning_numbers[number]; ok {
				fmt.Println("number: ", number, "is in")
				if card_total == 0 {
					card_total = 1
				} else {
					card_total = card_total * 2
				}
			}
		}
		fmt.Println("card_total: ", card_total)
		total = total + card_total
	}
	fmt.Println(total)
}
