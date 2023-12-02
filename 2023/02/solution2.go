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

	total := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		tmp := strings.Split(buff, ": ")
		tmp2 := strings.Split(tmp[0], " ")
		game, _ := strconv.Atoi(tmp2[1])
		game_power := calculate_game_power(tmp[1])
		fmt.Printf("%s -> game: %d, game_power: %d\n", buff, game, game_power)
		total += game_power
	}
	fmt.Println(total)
}

func calculate_game_power(game string) int {
	min_red := 0
	min_blue := 0
	min_green := 0

	for _, play := range strings.Split(game, "; ") {
		for _, color_n := range strings.Split(play, ", ") {
			fmt.Println(color_n)
			tmp := strings.Split(color_n, " ")
			fmt.Println(tmp)
			n, _ := strconv.Atoi(tmp[0])
			if strings.Compare(tmp[1], "red") == 0 {
				if n > min_red {
					min_red = n
				}
			} else if strings.Compare(tmp[1], "blue") == 0 {
				if n > min_blue {
					min_blue = n
				}
			} else if strings.Compare(tmp[1], "green") == 0 {
				if n > min_green {
					min_green = n
				}
			}
		}
	}
	return min_blue * min_red * min_green
}
