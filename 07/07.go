package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

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

	ranks := RankHands(hands)
	fmt.Println(bids)
	fmt.Println(ranks)

	winnings := 0
	for i, rank := range ranks {
		winnings += rank * bids[i]
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
	// true if a-b < 0
	aB, bB := []rune(a), []rune(b)
	cB := make([]int, len(a))
	for i := range aB {
		cB[i] = int(aB[i] - bB[i])
	}
	return !Any(cB, func(i int) bool { return i < 0 })
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
func RankHands(hands []Hand) []int {
	indexedHands := NewIndexPairing(hands)
	sort.Sort(indexedHands)
	// N := len(hands)
	result := make([]int, len(hands))
	for i, v := range indexedHands {
		result[i] = v.initPos + 1
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

// Any is a short circuit at first true of Reduce(iterable ~[]V, predicate func(V) bool) bool
func Any[T ~[]V, V any](t T, pred func(V) bool) bool {
	for _, v := range t {
		b := pred(v)
		if b {
			return true
		}
	}
	return false
}

// All is a short circuit at first false of Reduce(iterable ~[]V, predicate func(V) bool) bool
func All[T ~[]V, V any](t T, pred func(V) bool) bool {
	for _, v := range t {
		if !pred(v) {
			return false
		}
	}
	return true
}
