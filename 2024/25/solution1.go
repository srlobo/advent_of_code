package main

import (
	"bufio"
	"fmt"
	"os"
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

	j := 0

	keyLock := KeyLockRepresentation{}
	var keys, locks [][5]int

	for fileScanner.Scan() {
		buff := fileScanner.Text()
		if buff == "" {
			if keyLock.isKey() {
				keys = append(keys, keyLock.getNumericRepresentation())
			} else {
				locks = append(locks, keyLock.getNumericRepresentation())
			}
			keyLock = KeyLockRepresentation{}
			j = 0
			continue
		}
		for i, r := range buff {
			keyLock[j][i] = r
		}
		j++
	}
	if keyLock.isKey() {
		keys = append(keys, keyLock.getNumericRepresentation())
	} else {
		locks = append(locks, keyLock.getNumericRepresentation())
	}
	keyLock = KeyLockRepresentation{}

	fmt.Println("keys", keys)
	fmt.Println("locks", locks)
	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			if matchKeyLock(key, lock) {
				count++
			}
		}
	}
	fmt.Println(count)
}

type KeyLockRepresentation [7][5]rune

func (kl KeyLockRepresentation) isKey() bool {
	for i := 0; i < 5; i++ {
		if kl[0][i] != '#' {
			return true
		}
	}
	return false
}

func (kl KeyLockRepresentation) getNumericRepresentation() [5]int {
	representation := [5]int{0, 0, 0, 0, 0}
	for i := 0; i < 5; i++ {
		if kl.isKey() {
			for j := 6; j >= 0; j-- {
				if kl[j][i] != '#' {
					representation[i] = 5 - j
					break
				}
			}
		} else {
			for j := 0; j < 7; j++ {
				if kl[j][i] != '#' {
					representation[i] = j - 1
					break
				}
			}
		}
	}
	return representation
}

func (kl KeyLockRepresentation) String() string {
	representation := ""
	for j := 0; j < 7; j++ {
		for i := 0; i < 5; i++ {
			representation += string(kl[j][i])
		}
		representation += "\n"
	}
	return representation
}

func matchKeyLock(key, lock [5]int) bool {
	for i := 0; i < 5; i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}
