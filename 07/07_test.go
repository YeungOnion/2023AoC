package main

import (
	"fmt"
	"testing"
)

func TestIndexing(t *testing.T) {
	hands := []Hand{
		"12345", // 2
		"01234", // 1
		"00000", // 0
	}

	ranks := RankHands(hands)
	fmt.Println(hands[ranks[0]])
	fmt.Println(hands[ranks[1]])
	fmt.Println(hands[ranks[2]])
	return
}
