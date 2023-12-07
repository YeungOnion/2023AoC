package main

import (
	"YeungOnion/2023AoC/avl"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type seedMap struct {
	src    int
	dest   int
	window int
}

func seedMapCompare(a, b seedMap) avl.Ordering {
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

func main() {
	filename := "05/sample-a.txt"

	// stream file by words
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	seedNums := ScanSeedLine(fileScanner)

	nextNums := make([]int, len(seedNums))
	for fileScanner.Scan() { // This scan "eats" the map header
		fmt.Printf("\nnow: %v\n", seedNums)
		fmt.Println(fileScanner.Text())
		lines := ScanWhile(fileScanner, numberSequenceBuffered)
		fmt.Println("\t" + strings.Join(lines, "\n\t"))

		seedTable := avl.NewBST[seedMap](seedMapCompare)
		ParseRowsToTable(lines, seedTable)

		for i, n := range seedNums {
			nextNums[i] = SeedTableEval(seedTable, n)
		}
		seedNums = nextNums
	}

	fmt.Printf("now: %v\n\n", seedNums)
	return
}

func SeedTableEval(seedTable *avl.BST[seedMap], srcValue int) int {
	floorNode := seedTable.FloorSearch(seedMap{src: srcValue})
	fmt.Println(floorNode)
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

	seedline := ScanWhile(fs, numberSequenceBuffered)
	if len(seedline) != 1 {
		panic("ReadSeedLine: expects one line of seeds only")
	}

	return parseSeedLine(seedline[0])
}

func parseSeedLine(s string) []int {
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

// ScanWhile consumes input from scanner via Scan, if Scanned matches the predicate
// Then the text it appended to a string slice output
func ScanWhile(fs *bufio.Scanner, pred func(*bufio.Scanner) bool) []string {
	textRows := make([]string, 0, 16)
	for fs.Scan() {
		text := fs.Text()
		if !pred(fs) {
			break
		} else {
			textRows = append(textRows, text)
		}

	}

	return textRows
}

func ParseRowsToTable(rows []string, result *avl.BST[seedMap]) {
	digitsRe := regexp.MustCompile(`\d+`)
	seedMaps := lo.Map(rows, func(item string, index int) seedMap {
		nums := lo.Map(digitsRe.FindAllString(item, -1), MustAtoi)
		return seedMap{src: nums[1], dest: nums[0], window: nums[2]}
	})

	for _, s := range seedMaps {
		result.Insert(s)
	}

}

func (t seedMap) String() string {
	return fmt.Sprintf("\t[%d,%d): [%d,%d)", t.src, t.src+t.window, t.dest, t.dest+t.window)
}

func (c seedMap) Eval(key int) int {
	if c.src+c.window <= key {
		return key
	}
	return c.dest + key - c.src
}
