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
		game_result := true
		for _, play := range strings.Split(tmp[1], "; ") {
			fmt.Printf("    play %s\n", play)
			if !check_play(play) {
				game_result = false
				break
			}
		}

		fmt.Printf("%s -> game: %d\n", buff, game)
		if game_result {
			fmt.Printf("Game %d, possible\n", game)
			total += game
		} else {
			fmt.Printf("Game %d, impossible\n", game)
		}

	}
	fmt.Println(total)
}

func check_play(play string) bool {
	result := true
	fmt.Println(play)
	for _, color_n := range strings.Split(play, ", ") {
		fmt.Println(color_n)
		tmp := strings.Split(color_n, " ")
		fmt.Println(tmp)
		n, _ := strconv.Atoi(tmp[0])
		fmt.Printf("Compare: %s - %d\n", tmp[1], n)
		if strings.Compare(tmp[1], "red") == 0 {
			if n > 12 {
				result = false
				break
			}
		} else if strings.Compare(tmp[1], "blue") == 0 {
			if n > 14 {
				result = false
				break
			}
		} else if strings.Compare(tmp[1], "green") == 0 {
			if n > 13 {
				result = false
				break
			}
		}
	}
	return result
}
