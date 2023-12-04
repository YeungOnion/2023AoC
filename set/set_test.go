package set_test

import (
	"YeungOnion/2023AoC/set"
	"testing"
)

func TestNew(t *testing.T) {
	set.New[int]()
	set.New[rune]()
	return

}
