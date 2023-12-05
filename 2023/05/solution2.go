package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	seeds := make([][2]int, 0)
	mega_map := make([][][3]int, 0)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		// fmt.Println(buff)

		if strings.HasPrefix(buff, "seeds:") {
			tmp := strings.Split(buff, ": ")
			tmp = strings.Split(tmp[1], " ")

			for i := 0; 2*i <= len(tmp); i += 2 {
				var s [2]int
				s[0], _ = strconv.Atoi(tmp[i])
				s[1], _ = strconv.Atoi(tmp[i+1])
				seeds = append(seeds, [2]int{s[0], s[1]})
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
			sort.Slice(this_map, func(i, j int) bool {
				return this_map[i][1] < this_map[j][1]
			})
			mega_map = append(mega_map, this_map)
		}
	}
	fmt.Println(mega_map)
	fmt.Println(seeds)

	lower_location := MaxInt
	for _, seed := range seeds {
		location := convertMap(seed, 0, mega_map)
		fmt.Println("seed: ", seed, "location: ", location)
		if location < lower_location {
			lower_location = location
		}
	}
	fmt.Println(lower_location)
}

func convertMap(number_range [2]int, level int, use_map [][][3]int) int {
	// fmt.Println(use_map)
	// fmt.Printf("%d", number_range)
	// fmt.Println("number_range: ", number_range)
	// fmt.Println()
	if level == len(use_map) {
		return number_range[0]
	} else {

		lower := MaxInt
		for _, interval := range divideIntervals(number_range, use_map[level]) {
			n := convertMap(interval, level+1, use_map)
			if n < lower {
				lower = n
			}
		}
		// fmt.Printf(" -> ")
		return lower
	}
}

func divideIntervals(origin [2]int, rules [][3]int) [][2]int {
	// fmt.Println("Enter divideIntervals")
	intervals := make([][2]int, 0)
	// fmt.Println("origin: ", origin)
	// fmt.Println("rules: ", rules)
	beginning := origin[0]
	end := origin[0] + origin[1]
	for _, rule := range rules {
		// fmt.Println("beginning: ", beginning, "end: ", end, "rule: ", rule, "intervals for now: ", intervals)
		if beginning > rule[1]+rule[2] {
			continue
		}
		if beginning < rule[1] { // out of the intervals
			if end < rule[1] {
				intervals = append(intervals, [2]int{beginning, end - beginning})
				break
			} else {
				intervals = append(intervals, [2]int{beginning, rule[1] - beginning})
				beginning = rule[1]
			}
		}
		if end <= rule[1]+rule[2] {
			intervals = append(intervals, [2]int{mapRule(beginning, rule), end - beginning})
			beginning = end
			break
		} else if end > rule[1]+rule[2] {
			intervals = append(intervals, [2]int{mapRule(beginning, rule), rule[1] + rule[2] - beginning})
			beginning = rule[1] + rule[2]
		}
	}
	if beginning != end {
		intervals = append(intervals, [2]int{beginning, end - beginning})
	}

	// fmt.Println("intervals final", intervals)
	return intervals
}

func mapRule(origin int, rule [3]int) int {
	if origin < rule[1] {
		return origin
	} else if origin < rule[1]+rule[2] {
		return rule[0] + (origin - rule[1])
	} else {
		return origin
	}
}
