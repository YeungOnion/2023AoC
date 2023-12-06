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

	partScore, gearScore := ScanAndScore(fileScanner)
	fmt.Println("score: ", partScore)
	fmt.Println("gear: ", gearScore)
}

// ScanAndScore computes score for a Scanner
// assumes that scanner has not yet made call to Scan
func ScanAndScore(fs *bufio.Scanner) (int, int) {
	var partScore, gearScore int
	lines := make([]string, 3, 3)

	// process top edge / first line
	PrepTopEdge(fs, lines)
	partScore = ScorePartsMiddleLine(lines)
	gearScore = ScoreGearRatioMiddleLine(lines)

	// iterate file linewise
	for fs.Scan() {
		lines[0], lines[1], lines[2] = lines[1], lines[2], fs.Text()
		partScore += ScorePartsMiddleLine(lines)
		gearScore += ScoreGearRatioMiddleLine(lines)
	}

	// process bottom edge / last line
	PrepBottomEdge(fs, lines)
	partScore += ScorePartsMiddleLine(lines)
	gearScore += ScoreGearRatioMiddleLine(lines)

	return partScore, gearScore
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
	if !(!fs.Scan() && fs.Err() == nil) {
		panic("PrepBottomEdge: expects Scanner buffer to be at EOF without errors")
	}
	lines[0], lines[1] = lines[1], lines[2]
	lines[2] = EmptyString(len(lines[1]))
	return
}

// EmptyString returns the "empty" string used in this puzzle
// of specified length
func EmptyString(count int) string {
	return strings.Repeat(".", count)
}

// ScorePartsMiddleLine computes the score of a line given 3 lines
// as a slice - they should be in order of appearance
func ScorePartsMiddleLine(lines []string) int {
	score := 0
	periphRe := regexp.MustCompile(`[^\d\.]`)
	targetRe := regexp.MustCompile(`\d+`)
	for _, index := range GetTargetIndexes(lines[1], targetRe) {
		number, err := strconv.Atoi(lines[1][index[0]:index[1]])
		if err != nil {
			panic("cannot convert digit sequence to number")
		}
		if HasNeighborPeripherals(index, lines, periphRe) {
			score += number
		}
	}
	return score
}

// ScoreGearRatioMiddleLine computes the score of a line given 3 lines
// as a slice - they should be in order of appearance
func ScoreGearRatioMiddleLine(lines []string) int {
	score := 0
	periphRe := regexp.MustCompile(`\d+`)
	targetRe := regexp.MustCompile(`\*`)
	// for each target...
	for _, index := range GetTargetIndexes(lines[1], targetRe) {
		nb := NeighborPeripherals(index, lines, periphRe)
		// where there are two peripherals
		if len(nb) == 2 {
			a, _ := strconv.Atoi(lines[nb[0][2]][nb[0][0]:nb[0][1]])
			b, _ := strconv.Atoi(lines[nb[1][2]][nb[1][0]:nb[1][1]])
			score += a * b
		}
	}
	return score
}

// NeighborPeripherals returns slice of peripherals index and line, peripheral described by regex
// compared to provided lines
// regex match corresponds to lines via
// ```
// nbs := NeighborPeripherals(targetIdx, lines, periphRe)
//
//	for _, nb := range nbs {
//	  lines[nb[2]][nb[0]:nb[1]]
//	}
//
// ```
// Note that this matches the convention of indices returned by regexp in that,
// 0 and 1 are half open start and end positions of the match
func NeighborPeripherals(index []int, lines []string, perhiphRe *regexp.Regexp) [][]int {
	indices := make([][]int, 0, 2)
	for i, line := range lines {
		for _, periphIndex := range perhiphRe.FindAllStringIndex(line, -1) {
			if AdjacentColumn(index, periphIndex) {
				indices = append(indices, append(periphIndex, i))
			}
		}
	}

	return indices

}

// HasNeighborPeripherals identifies if any of the provided lines have
// peripherals defined by regex are column adjacent to the indices specified
func HasNeighborPeripherals(index []int, lines []string, periphRe *regexp.Regexp) bool {
	return Any(lines, func(l string) bool {
		return AnyAdjacentColumn(index, l, periphRe)
	})
}

// CountNeighborPeripherals returns number of peripherals defined
// by regex given lines and indices of target for middle line
func CountNeighborPeripherals(index []int, lines []string, periphRe *regexp.Regexp) int {
	count := 0
	for _, l := range lines {
		count += Count(periphRe.FindAllStringIndex(l, -1), func(periphIdx []int) bool {
			return AdjacentColumn(index, periphIdx)
		})
	}
	return count
}

// GetTargetIndexes returns indices of all matches for the digit sequence
// provided a single string
func GetTargetIndexes(numberRow string, targetRe *regexp.Regexp) [][]int {
	return targetRe.FindAllStringIndex(string(numberRow), -1)
}

func AdjacentColumn(targetIndex []int, peripheralIndex []int) bool {
	return targetIndex[0] <= peripheralIndex[1] &&
		peripheralIndex[0] <= targetIndex[1]
}

// AnyAdjacentColumn checks if the provided string has a peripheral in a column adjcent to the location of an unspecified match string
// uses the regexp convention of the match string, s[matchIndex[0]:matchIndex[1]]
func AnyAdjacentColumn(targetIndex []int, searchRow string, periphRe *regexp.Regexp) bool {
	// periphRe := regexp.MustCompile(`[^\d\.]`)
	for _, index := range periphRe.FindAllStringIndex(searchRow, -1) {
		if AdjacentColumn(targetIndex, index) {
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
