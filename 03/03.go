package main

import (
	"fmt"
	"regexp"
)

func main() {
	return
}

func GetNumbers(numberRow string) [][]int {
	digitRe := regexp.MustCompile(`\d+`)
	return digitRe.FindAllStringIndex(string(numberRow), -1)

}

// AdjacentColumn checks if the provided string has a symbol in a column adjcent to the location of an unspecified match string
// uses the regexp convention of the match string, s[matchIndex[0]:matchIndex[1]]
func AdjacentColumn(matchIndex []int, searchRow string) bool {

	symbolRe := regexp.MustCompile(`[^\d\.]`)
	for _, index := range symbolRe.FindAllStringIndex(searchRow, -1) {
		if matchIndex[0] < index[1] && index[1] < matchIndex[1] {
			fmt.Println("found near front")
			return true
		}
		if index[0] < matchIndex[1] {
			return true
		}
	}

	return false
}
