package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	Five HandType = iota
	Four
	Full
	Three
	TwoPair
	OnePair
	High
)

type Hand struct {
	cards string
	t     HandType
	bid   int
}

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

	hands := make([]Hand, 0)
	for fileScanner.Scan() {
		buff := fileScanner.Text()
		fmt.Println(buff)

		tmp := strings.Split(buff, " ")
		t := detectHandType(tmp[0])
		bid, _ := strconv.Atoi(tmp[1])
		j := Hand{cards: tmp[0], t: t, bid: bid}
		hands = append(hands, j)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].t > hands[j].t {
			return true
		} else if hands[i].t < hands[j].t {
			return false
		} else { // Same type, let's compare cards
			for c := 0; c < len(hands[i].cards); c++ {
				compare := compareCards(string(hands[i].cards[c]), string(hands[j].cards[c]))
				if compare == 1 {
					return true
				} else if compare == -1 {
					return false
				}
			}
			panic("Same cards")
		}
	})
	total := 0
	for i, hand := range hands {
		hand.PrintHand()
		total = total + (i+1)*hand.bid
	}
	fmt.Println(total)
}

func detectHandType(hand string) HandType {
	counter := make(map[string]int, 0)
	for _, card := range hand {
		counter[string(card)]++
	}
	empty := struct{}{}
	detected_so_far := map[HandType]struct{}{}
	for _, v := range counter {
		if v == 5 {
			return Five
		} else if v == 4 {
			return Four
		} else if v == 3 { // Full or Three
			if _, ok := detected_so_far[OnePair]; ok {
				return Full
			}
			detected_so_far[Three] = empty
		} else if v == 2 {
			if _, ok := detected_so_far[OnePair]; ok {
				return TwoPair
			} else if _, ok := detected_so_far[Three]; ok {
				return Full
			}
			detected_so_far[OnePair] = empty
		}
	}
	if _, ok := detected_so_far[Three]; ok {
		return Three
	} else if _, ok := detected_so_far[OnePair]; ok {
		return OnePair
	} else {
		return High
	}
}

func (hand *Hand) PrintHand() {
	var hand_type_string string
	if hand.t == Five {
		hand_type_string = "Five"
	} else if hand.t == Four {
		hand_type_string = "Four"
	} else if hand.t == Full {
		hand_type_string = "Full"
	} else if hand.t == Three {
		hand_type_string = "Three"
	} else if hand.t == TwoPair {
		hand_type_string = "TwoPair"
	} else if hand.t == OnePair {
		hand_type_string = "OnePair"
	} else if hand.t == High {
		hand_type_string = "High"
	}
	fmt.Println("Hand: ", hand.cards, "Type: ", hand_type_string, "Bid: ", hand.bid)
}

func compareCards(c1, c2 string) int {
	cards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	var c1_index, c2_index int

	fmt.Println(cards)
	for i := range cards {
		if c1 == cards[i] {
			c1_index = i
		}
		if c2 == cards[i] {
			c2_index = i
		}
	}
	if c1_index > c2_index {
		return -1
	} else if c1_index < c2_index {
		return 1
	} else {
		return 0
	}
}
