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

		size_interval_1 := interval_elf_1[1] - interval_elf_1[0]
		size_interval_2 := interval_elf_2[1] - interval_elf_2[0]

		if size_interval_1 >= size_interval_2 {
			if (interval_elf_1[0] <= interval_elf_2[0]) && (interval_elf_1[1] >= interval_elf_2[1]) {
				count += 1
			}
		} else {
			if (interval_elf_2[0] <= interval_elf_1[0]) && (interval_elf_2[1] >= interval_elf_1[1]) {
				count += 1
			}

		}

	}
	fmt.Println(count)
}
