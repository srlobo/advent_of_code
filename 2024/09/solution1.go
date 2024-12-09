package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	display_length := 20
	total := 0
	fileScanner.Scan()
	buff := fileScanner.Text()
	fmt.Println(len(buff))
	var disk string
	i := 0
	j := len(buff)
	actual := 0
	end_file_id := j / 2
	reminder, _ := strconv.Atoi(string(buff[j-1]))
	id := 0
	for true {
		num, _ := strconv.Atoi(string(buff[i]))
		// position is even, it's a block.
		if i%2 == 0 {
			id = i / 2
			fmt.Println("id", id, "size", num)
			if id >= end_file_id {
				fmt.Println("id: reminder", reminder, "id", id, "end_file_id", end_file_id)
				for k := 0; k < reminder; k++ {
					disk += strconv.Itoa(end_file_id)
					total += end_file_id * actual
				}
				break
			}

			for k := 0; k < num; k++ {
				total += id * actual
				actual++
				disk += strconv.Itoa(id)
			}
		} else { // Position is odd, it's space, let's fill with the latest numbers
			fmt.Println("spaces; size:", num, "end_file_id", end_file_id, "reminder", reminder)
			for k := 0; k < num; k++ {
				total += end_file_id * actual
				disk += strconv.Itoa(end_file_id)
				actual++
				reminder--
				if reminder == 0 {
					j = j - 2
					end_file_id = j / 2
					reminder, _ = strconv.Atoi(string(buff[j-1]))
					fmt.Println("reminder", reminder, "end_file_id", end_file_id)
				}
				if id >= end_file_id {
					fmt.Println("k: reminder", reminder, "id", id, "end_file_id", end_file_id)
					break
				}

			}
		}
		if id >= end_file_id {
			break
		}
		if len(disk) > display_length {
			fmt.Println(disk[len(disk)-display_length:])
		} else {
			fmt.Println(disk)
		}
		i++

	}
	if len(disk) > display_length {
		fmt.Println(disk[len(disk)-display_length:])
	} else {
		fmt.Println(disk)
	}

	fmt.Println(total)
}
