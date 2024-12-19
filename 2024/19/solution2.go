package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var empty = struct{}{}

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

	var towels []string

	fileScanner.Scan()
	buff := strings.Split(fileScanner.Text(), ", ")
	for _, b := range buff {
		towels = append(towels, b)
	}

	r := "^(" + strings.Join(towels, "|") + ")*$"

	re := regexp.MustCompile(r)

	fileScanner.Scan()

	sum := 0
	cache := Cache{c: make(map[string]int)}
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if re.MatchString(buff) {
			fmt.Println("Checking ", buff)
			sum += cache.checkTowels(towels, buff)
		}

	}
	fmt.Println(sum)
}

type Cache struct {
	c map[string]int
}

func (cache *Cache) checkTowels(towels []string, objective string) int {
	if r, ok := cache.c[objective]; ok {
		return r
	}

	fmt.Println("Check objective: ", objective)
	if len(objective) == 0 {
		cache.c[objective] = 1
		return 1
	}
	ret := 0
	for _, towel := range towels {
		if len(objective) < len(towel) {
			continue
		}

		if towel == objective[:len(towel)] {
			ret += cache.checkTowels(towels, objective[len(towel):])
		}
	}
	cache.c[objective] = ret
	return ret
}
