package main

import (
	"YeungOnion/2023AoC/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type HandType int
type Hand string

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

func main() {
	filename := os.Args[1]
	if filename == "--" {
		filename = os.Args[2]
	}

	// stream file by lines
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	hands, bids := ParseInput(fileScanner)

	handToBids := make(map[Hand]int, len(hands))
	for i := range hands {
		handToBids[hands[i]] = bids[i]
	}
	sort.Sort(hands)

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * handToBids[hand]
	}

	fmt.Println("winnings: ", winnings)
}

func Less(a, b Hand) bool {
	aType, bType := EvaluateHandType(a), EvaluateHandType(b)
	if aType != bType {
		return aType < bType
	} else {
		return LessHandString(a, b)
	}
}

func LessHandString(a, b Hand) bool {
	u := EvaluateHandString(a)
	v := EvaluateHandString(b)
	// fmt.Printf("\n\t%s -> %v", a, u)
	// fmt.Printf("\n\t%s -> %v", b, v)
	return utils.Any(lo.Zip2(u, v), func(tup lo.Tuple2[int, int]) bool { return tup.A < tup.B })
}

func EvaluateHandString(h Hand) []int {
	return lo.Map([]rune(h), func(r rune, _ int) int { return EvaluateCard(r) })
}

func EvaluateCard(r rune) int {
	return strings.IndexRune("23456789TJQKA", r) + 2
}

func EvaluateHandType(h Hand) HandType {
	runes := []rune(h)
	dict := make(map[rune]int)
	for _, symbol := range runes {
		dict[symbol]++
	}

	switch len(dict) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		if len(lo.PickByValues(dict, []int{3})) > 0 {
			return ThreeKind
		} else {
			return TwoPair
		}
	case 2:
		if len(lo.PickByValues(dict, []int{4})) > 0 {
			return FourKind
		} else {
			return FullHouse
		}
	case 1:
		return FiveKind
	default:
		panic("unreachable")
	}
}

func ParseInput(fs *bufio.Scanner) (Hands, []int) {
	hands, bids := make([]Hand, 0), make([]int, 0)
	for fs.Scan() {
		hand := Hand(fs.Text())
		fs.Scan()
		bid, _ := strconv.Atoi(fs.Text())

		hands = append(hands, hand)
		bids = append(bids, bid)
	}
	return Hands(hands), bids
}

type Hands []Hand

func (s Hands) Len() int {
	return len(s)
}

func (s Hands) Less(i, j int) bool {
	return Less(s[i], s[j])
}

func (s Hands) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
