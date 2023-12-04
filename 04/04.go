package main

import (
	"YeungOnion/2023AoC/set"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	filename := "04/input.txt"
	digitsRe := regexp.MustCompile(`\d+`)
	NumberString := func(s string) bool { return digitsRe.MatchString(s) }
	NumberStringer := func(s string) int {
		if val, err := strconv.Atoi(s); err == nil {
			return val
		} else {
			panic(err)
		}
	}

	// stream file by words
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	// lets get this bread
	pointWinnings := 0
	cardCount := make(map[int]int, 200)

	// first two words are "Card", "\d+:", discard both
	fileScanner.Scan()       // consume first "Card"
	for fileScanner.Scan() { // consume "\d+:"
		cardNum, err := strconv.Atoi(digitsRe.FindString(fileScanner.Text()))
		if err != nil {
			panic("error while processing card number " + fileScanner.Text())
		}

		cardCount[cardNum]++

		// consume words that are digit sequences (and the delimiting "|")
		myNumbers := ScanWhile[int](fileScanner, NumberString, NumberStringer)
		// consume words that are digit sequences (as well as "\nCard")
		winningNumbers := ScanWhile[int](fileScanner, NumberString, NumberStringer)

		countMatches := winningNumbers.CountMatches(myNumbers)

		// solves part A
		if countMatches > 0 {
			pointWinnings += 1 << (countMatches - 1)
		}

		// solves part B
		for i := 1; i <= countMatches; i++ {
			cardCount[cardNum+i] += cardCount[cardNum]
		}
	}

	fmt.Println("winnings:", pointWinnings)
	fmt.Println(
		"num of cards",
		lo.Reduce(
			lo.MapToSlice(cardCount, func(k int, v int) int { return v }),
			func(agg int, item int, _ int) int {
				return agg + item
			},
			0),
	)

}

// ScanWhile scans and transforms string input and adds it to a set.Set if predicate matches
// as well as first nonmatch, use case expects this to be "Card" or "|"
func ScanWhile[T comparable](b *bufio.Scanner, pred func(string) bool, transform func(string) T) set.Set[T] {
	out := make(set.Set[T], 0)
	for b.Scan() {
		word := b.Text()
		if !pred(word) {
			break
		}
		out[transform(word)] = struct{}{}
	}
	return out
}
