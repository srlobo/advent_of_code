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

type crateStruct map[int][]string

func main() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var crate_buff []string
	for {
		if !fileScanner.Scan() {
			panic("Finished before found empty line")
		}
		buff := fileScanner.Text()

		if buff == "" {
			break
		}

		crate_buff = append(crate_buff, buff)
	}

	crate_struct := make(crateStruct)

	for i := len(crate_buff) - 1; i >= 0; i-- {
		// fmt.Println(crate_buff[i])
		for j := 0; j < (len(crate_buff[i])+1)/4; j++ {
			// fmt.Printf("j: %d, index: %d\n", j, (j*4)+1)
			value := crate_buff[i][j*4+1]
			if string(value) != " " {
				crate_struct[j+1] = append(crate_struct[j+1], string(value))
			}
		}
	}

	fmt.Println(crate_struct)

	// Starting with movements
	for fileScanner.Scan() {
		buff := string(fileScanner.Text())
		fmt.Println(buff)
		buffSplitted := strings.Split(buff, " ")

		amount, _ := strconv.Atoi(buffSplitted[1])
		from, _ := strconv.Atoi(buffSplitted[3])
		to, _ := strconv.Atoi(buffSplitted[5])
		// fmt.Printf("%d %d %d\n", ammount, from, to)
		crate_struct.moveCrate(from, to, amount)
		fmt.Println(crate_struct)
	}
	fmt.Println(crate_struct.getTops())
}

func (crate_struct crateStruct) moveCrate(from int, to int, amount int) {
	fmt.Printf("moveCrate %d -> %d; amount: %d", from, to, amount)
	vals := crate_struct[from][len(crate_struct[from])-amount : len(crate_struct[from])]
	fmt.Printf(", vals: %s\n", vals)
	crate_struct[from] = crate_struct[from][:len(crate_struct[from])-amount]
	crate_struct[to] = append(crate_struct[to], vals...)
}

func (crate_struct crateStruct) getTops() string {
	ret := ""

	for i := 1; i <= len(crate_struct); i++ {
		ret = ret + crate_struct[i][len(crate_struct[i])-1]
	}
	return ret
}
