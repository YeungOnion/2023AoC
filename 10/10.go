package main

import (
	"YeungOnion/2023AoC/utils"
	"YeungOnion/2023AoC/utils/iter"
	"fmt"
	"os"

	"github.com/samber/lo"
)

func main() {
	fname := os.Args[1]
	if fname == "--" {
		fname = os.Args[2]
	}
	fs, closeFileHandle := utils.FileScanner(fname)
	defer closeFileHandle()
	lines := utils.ScanWhile(fs, nil)
	tm := NewTileMap(ParseStringsToPipeTable(lines))
	path := make([]Loc, 1, 100)
	path[0] = tm.start

	fmt.Println(tm)

	firstDir := AvailMoves(tm, tm.start, tm.start)[0]
	path = append(path, FindPath(tm, tm.start, tm.start.Add(firstDir))...)

	fmt.Println("count:", (len(path)+1)/2)
	return
}

// AttemptMove does not fail, but will return same location if move invalid
func AttemptMove(tm *TileMap, curr Loc, dir Dir) Loc {
	proposed := curr.Add(dir)

	if curr.ValidMove(tm, dir) {
		return proposed
	} else {
		return curr
	}
}

func AvailMoves(tm *TileMap, prev, curr Loc) []Dir {
	return lo.Filter([]Dir{East, North, West, South}, func(dir Dir, _ int) bool {
		newPos := AttemptMove(tm, curr, dir)
		return !(newPos == curr || newPos == prev)
	})
}

func (l Loc) Add(d Dir) Loc {
	row, col := l[0], l[1]
	switch d {
	case East:
		return Loc{row + 0, col + 1}
	case North:
		return Loc{row - 1, col + 0}
	case West:
		return Loc{row + 0, col - 1}
	case South:
		return Loc{row + 1, col + 0}
	default:
		panic("Loc.Add: not a value for Dir" + d.String())
	}
}

func Next(tm *TileMap, prev, curr Loc) Loc {
	locs := iter.Map(AvailMoves(tm, prev, curr), func(d Dir) Loc { return curr.Add(d) })
	if len(locs) != 1 {
		fmt.Println("prev:", prev, ",  curr:", curr)
		fmt.Println("dirs:", locs)
		panic("main.Next: more than one move available")
	} else if len(locs) == 0 {
		panic("main.Next: no available moves")
	}
	return locs[0]
}

func FindPath(tm *TileMap, prev, curr Loc) []Loc {
	next := Next(tm, prev, curr)
	if p, ok := tm.At(next); ok && p == Starting {
		return []Loc{curr}
	}

	return append(FindPath(tm, curr, next), curr)
}
