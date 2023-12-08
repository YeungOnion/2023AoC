package main

import (
	"YeungOnion/2023AoC/avl"
	"YeungOnion/2023AoC/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type SeedMap struct {
	src    int
	dest   int
	window int
}

func main() {
	filename := os.Args[1]
	if filename == "--" {
		filename = os.Args[2]
	}

	// stream file by words
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	nextNums := ScanSeedLine(fileScanner)

	for fileScanner.Scan() { // This scan "eats" the map header
		lines := utils.ScanWhile(fileScanner, numberSequenceBuffered)

		seedTable := avl.NewBST[SeedMap](seedMapCompare)
		ParseRowsToTable(lines, seedTable)

		lo.Map(nextNums, func(n int, _ int) int { return SeedTableEval(seedTable, n) })
	}

	fmt.Printf("now: %v\n\n", nextNums)
	return
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

func SeedTableEval(seedTable *avl.BST[SeedMap], srcValue int) int {
	floorNode := seedTable.FloorSearch(SeedMap{src: srcValue})
	return floorNode.Value.Eval(srcValue)
}

func MustAtoi(s string, _ int) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func ScanSeedLine(fs *bufio.Scanner) []int {
	if len(fs.Bytes()) > 0 {
		panic("ReadSeedLine: expects to call at start of file")
	}

	seedline := utils.ScanWhile(fs, numberSequenceBuffered)
	if len(seedline) != 1 {
		panic("ReadSeedLine: expects one line of seeds only")
	}

	return ParseSeedLine(seedline[0])
}

func ParseSeedLine(s string) []int {
	out := make([]int, 0, 16)
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		if num, err := strconv.Atoi(scan.Text()); err == nil {
			out = append(out, num)
		}
	}
	return out
}

func numberSequenceBuffered(fs *bufio.Scanner) bool {
	re := regexp.MustCompile(`(\d+ *)+`)
	return re.MatchString(fs.Text())
}

func ParseRowsToTable(rows []string, result *avl.BST[SeedMap]) {
	digitsRe := regexp.MustCompile(`\d+`)
	seedMaps := lo.Map(rows, func(item string, index int) SeedMap {
		nums := lo.Map(digitsRe.FindAllString(item, -1), MustAtoi)
		return SeedMap{src: nums[1], dest: nums[0], window: nums[2]}
	})

	for _, s := range seedMaps {
		result.Insert(s)
	}

}

func (t SeedMap) String() string {
	return fmt.Sprintf("\t[%d,%d): [%d,%d)", t.src, t.src+t.window, t.dest, t.dest+t.window)
}

func (c SeedMap) Eval(key int) int {
	if c.src+c.window <= key {
		return key
	}
	return c.dest + key - c.src
}
