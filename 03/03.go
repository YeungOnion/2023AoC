package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// for matters of language, it will be helpful to use the terms
// "target" and "peripheral", as:
// part A of the problem asks,
// > for all targets that are digit sequences that have any
// > peripheral symbols, sum those numbers (specified by digits)
// part B instead asks,
// > for all target of the '*' symbol that have two peripheral
// digit sequences, compute the product of the periphal numbers
// and accumulate this product for all targets

func main() {
	filename := "03/input.txt"

	// stream file by words
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fmt.Println("score: ", ScanAndScore(fileScanner))
}

// ScanAndScore computes score for a Scanner
// assumes that scanner has not yet made call to Scan
func ScanAndScore(fs *bufio.Scanner) int {
	lines := make([]string, 3, 3)

	// process top edge / first line
	PrepTopEdge(fs, lines)
	score := ScoreMiddleLine(lines)

	// iterate file linewise
	for fs.Scan() {
		lines[0], lines[1], lines[2] = lines[1], lines[2], fs.Text()
		score += ScoreMiddleLine(lines)
	}

	// process bottom edge / last line
	PrepBottomEdge(fs, lines)
	score += ScoreMiddleLine(lines)

	return score
}

// PrepTopEdge gets input from Scanner and modifies lines slice
// for scoring via ScoreMiddleLine
func PrepTopEdge(fs *bufio.Scanner, lines []string) {
	if len(fs.Bytes()) > 0 {
		panic("PrepTopEdge: expects to call on unscanned Scanner buffer")
	}
	_, lines[1] = fs.Scan(), fs.Text()
	_, lines[2] = fs.Scan(), fs.Text()
	lines[0] = EmptyString(len(lines[1]))
	return
}

// PrepBottomEdge modifies lines slice for scoring via ScoreMiddleLine
// assumes no more input on Scanner, but used as arg for convenience
// panics if there is more input
func PrepBottomEdge(fs *bufio.Scanner, lines []string) {
	// if !(!fs.Scan() && fs.Err() == nil) {
	// 	panic("PrepBottomEdge: expects Scanner buffer to be at EOF without errors")
	// }
	lines[0], lines[1] = lines[1], lines[2]
	lines[2] = EmptyString(len(lines[1]))
	return
}

// EmptyString returns the "empty" string used in this puzzle
// of specified length
func EmptyString(count int) string {
	return strings.Repeat(".", count)
}

// ScoreMiddleLine computes the score of a line given 3 lines
// as a slice - they should be in order of appearance
func ScoreMiddleLine(lines []string) int {
	score := 0
	for _, index := range GetTargetIndexes(lines[1]) {
		number, err := strconv.Atoi(lines[1][index[0]:index[1]])
		if err != nil {
			panic("cannot convert digit sequence to number")
		}
		if HasNeighborPeripherals(index, lines) {
			score += number
		}
	}
	return score
}

// HasNeighborPeripherals identifies if any of the provided lines
// are column adjacent to the indices specified
func HasNeighborPeripherals(index []int, neighborLines []string) bool {
	periphRe := regexp.MustCompile(`[^\d\.]`)
	return Any(neighborLines, func(line string) bool {
		return AdjacentColumn(index, line, periphRe)
	})
}

// GetTargetIndexes returns indices of all matches for the digit sequence
// provided a single string
func GetTargetIndexes(numberRow string) [][]int {

	// target in part A was digit sequences
	targetRe := regexp.MustCompile(`\d+`)

	// // target in part B is literal asterisk
	// targetRe := regexp.MustCompile(`\*`)

	return targetRe.FindAllStringIndex(string(numberRow), -1)
}

// AdjacentColumn checks if the provided string has a peripheral in a column adjcent to the location of an unspecified match string
// uses the regexp convention of the match string, s[matchIndex[0]:matchIndex[1]]
func AdjacentColumn(matchIndex []int, searchRow string, periphRe *regexp.Regexp) bool {
	// periphRe := regexp.MustCompile(`[^\d\.]`)
	for _, index := range periphRe.FindAllStringIndex(searchRow, -1) {
		if matchIndex[0] <= index[1] &&
			index[0] <= matchIndex[1] {
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

// Count returns the number of elements from the iterable that match the predicate
func Count[T ~[]V, V any](t T, pred func(V) bool) int {
	count := 0
	for _, v := range t {
		if pred(v) {
			count++
		}
	}
	return count
}
