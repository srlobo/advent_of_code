package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	count := 0

	for fileScanner.Scan() {
		var interval_elf_1 [2]int
		var interval_elf_2 [2]int

		buff := fileScanner.Text()

		elfs_intervals := strings.Split(buff, ",")
		elfs_interval_1 := strings.Split(elfs_intervals[0], "-")
		elfs_interval_2 := strings.Split(elfs_intervals[1], "-")

		interval_elf_1[0], _ = strconv.Atoi(elfs_interval_1[0])
		interval_elf_1[1], _ = strconv.Atoi(elfs_interval_1[1])

		interval_elf_2[0], _ = strconv.Atoi(elfs_interval_2[0])
		interval_elf_2[1], _ = strconv.Atoi(elfs_interval_2[1])

		if !doesNotOverlap(interval_elf_1, interval_elf_2) {
			count += 1
		}

	}
	fmt.Println(count)
}

func doesNotOverlap(interval1 [2]int, interval2 [2]int) bool {
	if (interval1[0] < interval2[0]) && (interval1[1] < interval2[0]) {
		return true
	}

	if (interval2[0] < interval1[0]) && (interval2[1] < interval1[0]) {
		return true
	}

	return false
}
