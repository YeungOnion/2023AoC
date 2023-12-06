package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type seedMap struct {
	src    int
	dest   int
	window int
}

type seedRange [2]seedMap

type seedTable []seedMap

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
		lines := ScanWhile(fileScanner, NumberSequenceBuffered)
		fmt.Println("\t" + strings.Join(lines, "\n\t"))
		table := ParseRowsToTable(lines)

		for i, n := range seedNums {
			nextNums[i] = table.Eval(n)
		}
		seedNums = nextNums
	}

	fmt.Printf("now: %v\n\n", seedNums)
	return
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

	seedline := ScanWhile(fs, NumberSequenceBuffered)
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

func parseMap(s string) {
	panic("unimplemented")
}

func PassthruMap(i int, m seedTable) int {
	panic("unimplemented")
}

func NumberSequenceBuffered(fs *bufio.Scanner) bool {
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

func ParseRowsToTable(rows []string) seedTable {
	digitsRe := regexp.MustCompile(`\d+`)
	s := seedTable(lo.Map(rows, func(item string, index int) seedMap {
		nums := lo.Map(digitsRe.FindAllString(item, -1), MustAtoi)
		return seedMap{src: nums[1], dest: nums[0], window: nums[2]}
	}))

	sort.Sort(s)

	return s
}

// FloorSearch returns the element position that has the largest value of
// seedMap.src not less than the needle
func (c seedTable) FloorSearch(needle int) int {

	if len(c) == 0 || needle < c[0].src {
		// precedes the first (implies all) bin edges
		// alternatively could have all seedTable's have seedMap{0,0,0}
		return -1
	}

	var mid int
	lb, ub := 0, len(c)
	for lb < ub {
		mid = (lb + ub) / 2
		f := c[mid].src
		if needle == f {
			return mid
		} else if needle < f {
			ub = mid
		} else {
			lb = mid + 1
		}
	}

	return mid

}

func (c seedTable) Eval(key int) int {
	index := c.FloorSearch(key)
	if index == -1 {
		return key
	}
	elem := c[index]
	if elem.src+elem.window <= key {
		return key
	}
	return elem.dest + key - elem.src
}

func (c seedTable) Len() int {
	return len(c)
}

func (c seedTable) Less(i, j int) bool {
	return c[i].src < c[j].src
}

func (c seedTable) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
