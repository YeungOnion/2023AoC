package main

import (
	"YeungOnion/2023AoC/utils"
	"YeungOnion/2023AoC/utils/iter"
	"strings"
)

type Dir int

const (
	East Dir = iota
	North
	West
	South
	Unknown
)

func (d Dir) Inverse() Dir {
	switch d {
	case East:
		return West
	case North:
		return South
	case West:
		return East
	case South:
		return North
	default:
		return Unknown
	}
}

func (d Dir) String() string {
	switch d {
	case East:
		return "➡"
	case North:
		return "⬆"
	case West:
		return "⬅"
	case South:
		return "⬇"
	default:
		panic("Dir.String, unexpected value has no string representation")
	}
}

type Loc [2]int
type TileMap struct {
	tiles [][]Pipe
	start Loc
}

func NewTileMap(in [][]Pipe) *TileMap {
	// check rectangular
	row0Len := len(in[0])
	utils.All(in, func(row []Pipe) bool {
		return len(row) == row0Len
	})

	for i, row := range in {
		for j, p := range row {
			if p == Starting {
				return &TileMap{tiles: in, start: [2]int{i, j}}
			}
		}
	}

	return nil
}

func (tm TileMap) String() string {
	return strings.Join(iter.Map(tm.tiles, PipeRowString), "\n")
}

func (tm TileMap) At(l Loc) (Pipe, bool) {
	nR, nC := len(tm.tiles), len(tm.tiles[0])
	x, y := l[0], l[1]
	if x < nR && y < nC &&
		0 <= x && 0 <= y {
		return tm.tiles[l[0]][l[1]], true
	} else {
		return Pipe(0), false
	}
}

type Pipe rune

const (
	Vertical Pipe = '|'
	Horizont Pipe = '-'
	UppRight Pipe = '7'
	UpprLeft Pipe = 'F'
	LowrLeft Pipe = 'L'
	LwrRight Pipe = 'J'
	OpenSpot Pipe = '.'
	Starting Pipe = 'S'
)

func NewPipe(r rune) Pipe {
	switch r {
	case '-':
		return Horizont
	case '|':
		return Vertical
	case '7':
		return UppRight
	case 'F':
		return UpprLeft
	case 'L':
		return LowrLeft
	case 'J':
		return LwrRight
	case '.':
		return OpenSpot
	case 'S':
		return Starting
	default:
		panic("ScanPipe(rune) Pipe:: unreachable")
	}
}

func PipeRowString(row []Pipe) string {
	return strings.Join(iter.Map(row, Pipe.String), "")
}

func (p Pipe) String() string {
	// return string(p)
	switch p {
	case Horizont:
		return "─"
	case Vertical:
		return "│"
	case UppRight:
		return "┐"
	case UpprLeft:
		return "┌"
	case LowrLeft:
		return "└"
	case LwrRight:
		return "┘"
	case Starting:
		return "S"
	case OpenSpot:
		return "⋅"
	default:
		panic("Pipe.String: this Pipe is not expressable as string")
	}
}

func (in Loc) ValidMove(tm *TileMap, d Dir) bool {
	p, ok0 := tm.At(in)
	n, ok1 := tm.At(in.Add(d))
	if !ok0 || !ok1 || (p == Starting && n == OpenSpot) {
		return false
	}
	return p.oneWayValid(d) && n.oneWayValid(d.Inverse())
}

func (in Pipe) oneWayValid(d Dir) bool {
	switch in {
	case Starting, OpenSpot:
		return true
	case Vertical:
		return d == North || d == South
	case Horizont:
		return d == East || d == West
	case UpprLeft:
		return d == East || d == South
	case UppRight:
		return d == West || d == South
	case LowrLeft:
		return d == East || d == North
	case LwrRight:
		return d == West || d == North
	default:
		panic("Pipe.Valid unrecognized pipe")
	}
}

func ParseStringToPipes(input string) []Pipe {
	return iter.Map([]rune(input), NewPipe)
}

func ParseStringsToPipeTable(input []string) [][]Pipe {
	return iter.Map(input, ParseStringToPipes)
}
