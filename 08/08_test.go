package main

import (
	"testing"
)

func TestStep(t *testing.T) {
	netwk := Network{"AAA": {"BBB", "CCC"}, "BBB": {"DDD", "EEE"}, "CCC": {"ZZZ", "GGG"}, "DDD": {"DDD", "DDD"}, "EEE": {"EEE", "EEE"}, "GGG": {"GGG", "GGG"}, "ZZZ": {"ZZZ", "ZZZ"}}
	ls := NewLoopedSlice[rune]([]rune("RL"))
	expected := []string{"CCC", "ZZZ"}

	pos := netwk.Step("AAA", &ls)
	t.Log("step")
	if got := pos; got != expected[0] {
		t.Fatalf("got %s, expected %s", got, expected)
	}

	t.Log("step")
	pos = netwk.Step(pos, &ls)
	if got := pos; got != expected[1] {
		t.Fatalf("got %s, expected %s", got, expected)
	}
}
