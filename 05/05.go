package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type seedRange struct {
	src  int
	dest int
}

type AocMap []seedRange

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
	fmt.Printf("seeds numbers\n\t>> %v", seedNums)
	fileScanner.Scan() // consume empty newline

	for fileScanner.Scan() { // This scan "eats" the map header

	}

	return
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

func PassthruMap(i int, m AocMap) int {
	panic("unimplemented")
}

func NumberSequenceBuffered(fs *bufio.Scanner) bool {
	re := regexp.MustCompile(`(\d+ *)+`)
	return re.MatchString(fs.Text())
}

// ScanWhile consumes input from scanner via Scan, if Scanned matches the predicate
// Then the text it appended to a string slice output
func ScanWhile(fs *bufio.Scanner, pred func(*bufio.Scanner) bool) []string {
	out := make([]string, 0, 16)
	for fs.Scan() {
		text := fs.Text()
		if !pred(fs) {
			break
		} else {
			out = append(out, text)
		}

	}
	return out
}

func PushKeyThroughMap(key int, m []int)
