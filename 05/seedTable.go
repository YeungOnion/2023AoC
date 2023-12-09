package main

import (
	"YeungOnion/2023AoC/avl"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type SeedMap struct {
	src    int
	dest   int
	window int
}

func seedMapCompare(a, b SeedMap) avl.Ordering {
	switch {
	case a.src == b.src:
		return avl.Equal
	case a.src < b.src:
		return avl.Less
	case a.src > b.src:
		return avl.Greater
	default:
		panic("unreachable")
	}
}

func (t SeedMap) String() string {
	return fmt.Sprintf("\t[%d,%d): [%d,%d)", t.src, t.src+t.window, t.dest, t.dest+t.window)
}

func (c SeedMap) Eval(key int) int {
	if key-c.src < c.window {
		return c.dest + key - c.src
	}
	return key
}

func (c SeedMap) EvalInverse(key int) int {
	if key-c.dest < c.window {
		return c.src + key - c.dest
	}
	return key
}

type SeedTable avl.BST[SeedMap]

func (table SeedTable) String() string {
	s := make([]SeedMap, 0, 1<<table.Root.Height)
	avl.InOrderTraversal[SeedMap](table.Root, &s)
	return strings.Join(
		lo.Map(s, func(m SeedMap, _ int) string { return m.String() }),
		"\n",
	)
}

func SeedTableEval(seedTable *avl.BST[SeedMap], keyValue int) int {
	floorNode := seedTable.FloorSearch(SeedMap{src: keyValue})
	if floorNode != nil {
		return floorNode.Value.Eval(keyValue)
	} else {
		return SeedMap{}.Eval(keyValue)
	}
}

func SeedTableRangeEval(seedTable *avl.BST[SeedMap], r Interval) []Interval {
	return nil
}
