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

	cards := make(map[int]int)

	card_n := 1
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
					card_total += 1
				}
			}
		}
		cards[card_n] = card_total
		card_n += 1
	}

	card_multiplier := make(map[int]int)
	for card := 1; card < card_n; card++ {
		current_card_matching_numbers := cards[card]

		var current_card_multiplier int
		if _, ok := card_multiplier[card]; ok {
			current_card_multiplier = card_multiplier[card] + 1
		} else {
			current_card_multiplier = 1
		}
		card_multiplier[card] = current_card_multiplier
		fmt.Println("card: ", card, "card_matching_numbers: ", current_card_matching_numbers, "card_multiplier: ", current_card_multiplier)

		for i := card + 1; i <= card+current_card_matching_numbers; i++ {
			if _, ok := card_multiplier[i]; ok {
				card_multiplier[i] += current_card_multiplier
			} else {
				card_multiplier[i] = current_card_multiplier
			}
		}
	}

	fmt.Println("cards: ", cards)
	fmt.Println("card_multiplier: ", card_multiplier)
	total := 0
	for k, v := range card_multiplier {
		if k >= card_n {
			continue
		}
		total += v
	}
	fmt.Println(total)
}
