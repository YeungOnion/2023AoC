package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var forwardDigitRe, reverseDigitRe string

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
	N := len(hands)

	fmt.Println(hands)
	fmt.Println(bids)

	ranks := RankHands(hands)

	winnings := 0
	for i, position := range ranks {
		winnings += (N - i) * bids[position]
	}

	fmt.Println("winnings: ", winnings)

	return
}

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

func Less(a, b Hand) bool {
	aType, bType := EvaluateHandType(a), EvaluateHandType(b)

	if aType != bType {
		return aType < bType
	} else {
		return a < b
	}
}

func EvaluateHandType(h Hand) HandType {

	panic("unimplemented")
}

func ParseInput(fs *bufio.Scanner) ([]Hand, []int) {
	hands, bids := make([]Hand, 0), make([]int, 0)
	for fs.Scan() {
		hand := Hand(fs.Text())
		fs.Scan()
		bid, _ := strconv.Atoi(fs.Text())

		hands = append(hands, hand)
		bids = append(bids, bid)
	}
	return hands, bids
}

// RankHands returns n element slice from (0,n) such that
// hands[RankHands[i-1]] is the ith largest (i=0 -> largest)
func RankHands(hands []Hand) []int {
	indexedHands := NewIndexPairing(hands)
	sort.Sort(indexedHands)
	result := make([]int, len(hands))
	for i, v := range indexedHands {
		result[v.initPos] = i
	}
	return result
}

type indexPair struct {
	hand    Hand
	initPos int
}

type indexPairSlice []indexPair

func NewIndexPairing(s []Hand) indexPairSlice {
	result := make([]indexPair, len(s))

	for i, v := range s {
		result[i] = indexPair{hand: v, initPos: i}
	}
	return result
}

func (s indexPairSlice) Len() int {
	return len(s)
}

func (s indexPairSlice) Less(i, j int) bool {
	return Less(s[i].hand, s[j].hand)
}

func (s indexPairSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
