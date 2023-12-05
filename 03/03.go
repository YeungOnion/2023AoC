package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

	// read first two lines
	lines := make([]string, 3, 3)
	_, lines[1] = fileScanner.Scan(), fileScanner.Text()
	_, lines[2] = fileScanner.Scan(), fileScanner.Text()
	lines[0] = strings.Repeat(".", len(lines[1]))

	score := ScoreMiddleLine(lines)

	for fileScanner.Scan() {
		lines[0] = lines[1]
		lines[1] = lines[2]
		lines[2] = fileScanner.Text()

		score += ScoreMiddleLine(lines)
	}
	lines[0] = lines[1]
	lines[1] = lines[2]
	lines[2] = strings.Repeat(".", len(lines[1]))

	fmt.Println("score: ", score)
}

func ScoreMiddleLine(lines []string) int {
	score := 0
	for _, index := range GetNumberIndexes(lines[1]) {
		number, err := strconv.Atoi(lines[1][index[0]:index[1]])
		if err != nil {
			panic("cannot convert digit sequence to number")
		}
		score += number
	}
	return score
}

func HasNeighborSymbol(index []int, neighborLines []string) bool {
	return Any(neighborLines, func(line string) bool {
		return AdjacentColumn(index, line)
	})
}

func GetNumberIndexes(numberRow string) [][]int {
	digitRe := regexp.MustCompile(`\d+`)
	return digitRe.FindAllStringIndex(string(numberRow), -1)
}

// AdjacentColumn checks if the provided string has a symbol in a column adjcent to the location of an unspecified match string
// uses the regexp convention of the match string, s[matchIndex[0]:matchIndex[1]]
func AdjacentColumn(matchIndex []int, searchRow string) bool {
	symbolRe := regexp.MustCompile(`[^\d\.]`)
	for _, index := range symbolRe.FindAllStringIndex(searchRow, -1) {
		if matchIndex[0] <= index[1] &&
			index[0] <= matchIndex[1] {
			return true
		}
	}
	return false
}

func All[T ~[]V, V any](t T, pred func(V) bool) bool {
	for _, v := range t {
		if !pred(v) {
			return false
		}
	}
	return true
}

func Any[T ~[]V, V any](t T, pred func(V) bool) bool {
	for _, v := range t {
		b := pred(v)
		if b {
			return true
		}
	}
	return false
}
