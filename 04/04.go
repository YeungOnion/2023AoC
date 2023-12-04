package main

import (
	"YeungOnion/2023AoC/set"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := "03/input.txt"
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
	winnings := 0

	// first two words are "Card", "\d+:", discard both
	fileScanner.Scan()       // consume first "Card"
	for fileScanner.Scan() { // consume "\d+:"
		// consume words that are digit sequences (and the delimiting "|")
		myNumbers := ScanWhile[int](fileScanner, NumberString, NumberStringer)
		// consume words that are digit sequences (as well as "\nCard")
		winningNumbers := ScanWhile[int](fileScanner, NumberString, NumberStringer)
		if count := winningNumbers.CountMatches(myNumbers); count > 0 {
			winnings += 1 << (count - 1)
		}
	}

	fmt.Println(winnings)

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
