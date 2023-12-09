package main

import (
	"YeungOnion/2023AoC/avl"
	"YeungOnion/2023AoC/utils"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Interval struct {
	key    int
	window int
}

func main() {
	filename := os.Args[1]
	if filename == "--" {
		filename = os.Args[2]
	}

	fs, fileCloseHandle := utils.FileScanner(filename)
	defer fileCloseHandle()
	fs.Split(bufio.ScanLines)

	return
}

func ParseAllTables(fs *bufio.Scanner) []*avl.BST[SeedMap] {
	tables := make([]*avl.BST[SeedMap], 7)

	for i := 0; fs.Scan(); i++ { // This scan "eats" the map header
		lines := utils.ScanWhile(fs, numberSequenceBuffered)
		seedTable := avl.NewBST[SeedMap](seedMapCompare)
		ParseRowsToTable(lines, seedTable)
		tables[i] = seedTable
	}

	return tables
}

func PushThroughMaps(tables []*avl.BST[SeedMap], value int) int {
	if len(tables) == 0 {
		return value
	}

	return PushThroughMaps(
		tables[1:],
		SeedTableEval(tables[0], value),
	)
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

	result := make([]int, 0, 16)
	for _, text := range strings.Fields(seedline[0]) {
		if num, err := strconv.Atoi(text); err == nil {
			result = append(result, num)
		}
	}
	return result
}

func numberSequenceBuffered(fs *bufio.Scanner) bool {
	re := regexp.MustCompile(`(\d+ *)+`)
	return re.MatchString(fs.Text())
}

func ParseRowsToTable(rows []string, result *avl.BST[SeedMap]) {
	digitsRe := regexp.MustCompile(`\d+`)
	seedMaps := lo.Map(
		rows,
		func(item string, index int) SeedMap {
			nums := lo.Map(digitsRe.FindAllString(item, -1), MustAtoi)
			return SeedMap{src: nums[1], dest: nums[0], window: nums[2]}
		})

	for _, s := range seedMaps {
		result.Insert(s)
	}
}
