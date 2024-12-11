package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	total := 0
	cache := make(cache)
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		for _, s := range strings.Split(buff, " ") {
			fmt.Println(s)
			total += cache.process(s, 1)
		}
	}
	fmt.Println(total)
}

func (cache *cache) process(stone string, depth int) int {
	var ret int
	size := len(stone)
	if depth > 75 {
		return 1
	}
	ret, ok := cache.cacheGet(stone, depth)
	if ok {
		return ret
	}

	if stone == "0" {
		// fmt.Println("0 -> 1; depth", depth)
		ret = cache.process("1", depth+1)
	} else if size%2 == 0 {
		firstHalf, _ := strconv.Atoi(stone[:size/2])
		// fmt.Println(stone, "->", firstHalf, "; depth", depth)
		var s string
		s = strconv.Itoa(firstHalf)
		total := cache.process(s, depth+1)
		secondHalf, _ := strconv.Atoi(stone[size/2:])
		// fmt.Println(stone, "->", secondHalf, "; depth", depth)
		s = strconv.Itoa(secondHalf)
		total2 := cache.process(s, depth+1)
		ret = total + total2
	} else {
		num, _ := strconv.Atoi(stone)
		num = num * 2024
		// fmt.Println(stone, "->", num, "; depth", depth)
		s := strconv.Itoa(num)
		total := cache.process(s, depth+1)
		ret = total
	}
	cache.cacheSet(stone, depth, ret)
	return ret
}

func (cache *cache) cacheSet(stone string, depth, total int) {
	(*cache)[cacheItem{stone, depth}] = total
}

func (cache *cache) cacheGet(stone string, depth int) (int, bool) {
	ret, ok := (*cache)[cacheItem{stone, depth}]
	return ret, ok
}

type (
	cache     map[cacheItem]int
	cacheItem struct {
		stone string
		depth int
	}
)
