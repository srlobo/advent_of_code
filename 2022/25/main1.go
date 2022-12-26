package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Printf("5^0: %d\n", pow(5, 0))
	fmt.Printf("5^1: %d\n", pow(5, 1))
	fmt.Printf("5^2: %d\n", pow(5, 2))

	big_sum := 0
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Printf("%s -> %d\n", buff, Snafu2Decimal(buff))
		big_sum += Snafu2Decimal(buff)
	}
	fmt.Println(big_sum)
	fmt.Println(Decimal2Snafu(big_sum))
}

func Snafu2Decimal(snafu string) int {
	ret := 0
	for i := len(snafu) - 1; i >= 0; i-- {
		power_of_5 := pow(5, len(snafu)-i-1)
		// fmt.Printf("i: %d; snafu: %c; power: %d\n", i, snafu[i], power_of_5)
		switch snafu[i] {
		case '2':
			ret += 2 * power_of_5
		case '1':
			ret += 1 * power_of_5
		case '0':
			ret += 0 * power_of_5
		case '-':
			ret += -1 * power_of_5
		case '=':
			ret += -2 * power_of_5
		}
	}
	return ret
}

func pow(base, exp int) int {
	ret := 1
	for i := 0; i < exp; i++ {
		ret *= base
	}

	return ret
}

func Decimal2Snafu(decimal int) string {
	var ret []string
	carry := 0
	i := 0
	for {
		remind := decimal % 5
		decimal = decimal / 5
		// fmt.Printf("Round %d, decimal: %d, remind: %d, carry: %d, ret: ", i, decimal, remind, carry)
		remind += carry
		carry = 0
		if remind == 0 || remind == 1 || remind == 2 {
			ret = append(ret, fmt.Sprintf("%d", remind))
		} else if remind == 3 {
			carry = +1
			ret = append(ret, "=")
		} else if remind == 4 {
			carry = +1
			ret = append(ret, "-")
		}
		// fmt.Println(ret)
		if decimal == 0 {
			break
		}
		i += 1
	}

	for left, right := 0, len(ret)-1; left < right; left, right = left+1, right-1 {
		ret[left], ret[right] = ret[right], ret[left]
	}
	return strings.Join(ret, "")
}
