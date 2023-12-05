package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
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

	seeds := make([]int, 0)
	mega_map := make([][][3]int, 0)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)

		if strings.HasPrefix(buff, "seeds:") {
			tmp := strings.Split(buff, ": ")
			tmp = strings.Split(tmp[1], " ")

			for _, s := range tmp {
				n, _ := strconv.Atoi(s)
				seeds = append(seeds, n)
			}

			continue
		}

		if strings.Contains(buff, ":") {
			this_map := make([][3]int, 0)
			for fileScanner.Scan() {
				buff := fileScanner.Text()
				if buff == "" {
					break
				}
				tmp := strings.Split(buff, " ")
				var tmp_n [3]int
				for i, s := range tmp {
					n, _ := strconv.Atoi(s)
					tmp_n[i] = n
				}
				this_map = append(this_map, tmp_n)

			}
			mega_map = append(mega_map, this_map)
		}
	}
	fmt.Println(seeds)
	fmt.Println(mega_map)

	lower_location := MaxInt
	for _, seed := range seeds {
		location := convertMap(seed, 0, mega_map)
		if location < lower_location {
			lower_location = location
		}
	}
	fmt.Println(lower_location)
}

func convertMap(number, level int, use_map [][][3]int) int {
	// fmt.Println(use_map)
	fmt.Printf("%d", number)
	// fmt.Println()
	if level == len(use_map) {
		// fmt.Println("level: ", level, "len: ", len(use_map))
		fmt.Println()
		return number
	} else {
		for _, m := range use_map[level] {
			// fmt.Println("Compare m: ", m)
			if number >= m[1] && number < (m[1]+m[2]) {
				fmt.Printf(" -> ")
				return convertMap(m[0]+(number-m[1]), level+1, use_map)
			}
		}
		fmt.Printf(" -> ")
		return convertMap(number, level+1, use_map)

	}
}
