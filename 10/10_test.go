package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	rows := []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}
	expectedStart := Loc([2]int{2, 0})

	tiles := ParseStringsToPipeTable(rows)
	if fmt.Sprintf(strings.Join(rows, "\n")) != fmt.Sprint(TileMap{tiles: tiles}) {
		t.Fatalf("failed parsing TileMap.tiles BOOF")
	}

	tm := NewTileMap(tiles)
	if tm.start != expectedStart {
		t.Fatalf("got: %v, expected: %v", tm.start, expectedStart)
	}

	return
}

func TestMoves(t *testing.T) {
	tiles := [][]Pipe{
		{OpenSpot, OpenSpot, UpprLeft, UppRight, OpenSpot},
		{OpenSpot, UpprLeft, LwrRight, Vertical, OpenSpot},
		{Starting, LwrRight, OpenSpot, LowrLeft, UppRight},
		{Vertical, UpprLeft, Horizont, Horizont, LwrRight},
		{LowrLeft, LwrRight, OpenSpot, OpenSpot, OpenSpot},
	}

	tm := NewTileMap(tiles)
	fmt.Println(tm)
	fmt.Println(AvailMoves(tm, tm.start, tm.start))
}
