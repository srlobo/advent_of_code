package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	conditions := make(Condiions)

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if buff == "" {
			break
		}
		splitBuff := strings.Split(buff, "|")
		before, _ := strconv.Atoi(splitBuff[0])
		after, _ := strconv.Atoi(splitBuff[1])
		if _, ok := conditions[before]; !ok {
			conditions[before] = Condition{}
		}
		if _, ok := conditions[after]; !ok {
			conditions[after] = Condition{}
		}

		b := conditions[before]
		b.after = append(b.after, after)
		conditions[before] = b

		a := conditions[after]
		a.before = append(a.before, before)
		conditions[after] = a
	}
	cmp := func(a, b int) int {
		if c, ok := conditions[a]; ok {
			for i := range c.after {
				if b == c.after[i] {
					// fmt.Printf("a: %v, b: %v, conditions[a].after\n", a, b)
					return -1
				}
			}
		}

		if c, ok := conditions[a]; ok {
			for i := range c.before {
				if b == c.before[i] {
					// fmt.Printf("a: %v, b: %v, conditions[a].before\n", a, b)
					return 1
				}
			}
		}

		if c, ok := conditions[b]; ok {
			for i := range c.before {
				if a == c.before[i] {
					// fmt.Printf("a: %v, b: %v, conditions[b].before\n", a, b)
					return -1
				}
			}
		}

		if c, ok := conditions[b]; ok {
			for i := range c.after {
				if a == c.after[i] {
					// fmt.Printf("a: %v, b: %v, conditions[b].after\n", a, b)
					return 1
				}
			}
		}

		return 0
	}

	fmt.Println(conditions)

	total := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		splitBuff := strings.Split(buff, ",")
		numberSeq := make([]int, len(splitBuff))
		numberSeqSorted := make([]int, len(splitBuff))
		for i := range splitBuff {
			c, _ := strconv.Atoi(splitBuff[i])
			numberSeq[i] = c
			numberSeqSorted[i] = c
		}

		slices.SortFunc(numberSeqSorted, cmp)
		fmt.Printf("%v = %v\n", numberSeq, numberSeqSorted)

		if sliceCmp(numberSeq, numberSeqSorted) {
			middle := numberSeq[len(numberSeq)/2]
			fmt.Printf("YES - %d\n", middle)
			total += middle
		} else {
			fmt.Println("NO")
		}
	}
	fmt.Println(total)
}

func sliceCmp(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type Condition struct {
	before []int
	after  []int
}

type Condiions map[int]Condition
