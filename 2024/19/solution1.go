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

	count := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if re.MatchString(buff) {
			count++
		}
	}
	fmt.Println(count)
}
