package main

import (
	"bufio"
	"fmt"
	"os"
	// "sort"
	// "strconv"
	"regexp"
)

type Looper[T any] interface {
	Next() T
}

type LoopedSlice[T any] struct {
	s    []T
	size int
	iter int
}

type Network map[string][2]string

func main() {
	filename := os.Args[1]
	if filename == "--" {
		filename = os.Args[2]
	}

	// stream file by lines
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	lrPattern, netwk := ParseFile(fs)
	fmt.Println(lrPattern.s)
	fmt.Println(netwk)

	steps := 0
	for pos := "AAA"; pos != "ZZZ"; pos = netwk.Step(pos, &lrPattern) {
		steps++
	}

	fmt.Println("steps: ", steps)

	return
}

func ParseFile(fs *bufio.Scanner) (LoopedSlice[rune], Network) {
	tripleLetterRe := regexp.MustCompile(`[A-Z]{3}`)

	_, header := fs.Scan(), fs.Text()
	_, lrPattern := fs.Scan(), NewLoopedSlice([]rune(header))
	netwk := make(Network, 750)
	for fs.Scan() {
		line := fs.Text()
		if len(line) == 0 {
			continue
		}
		matches := tripleLetterRe.FindAllString(line, -1)
		netwk[matches[0]] = [2]string(matches[1:3])
	}
	return lrPattern, netwk
}

func NewLoopedSlice[T any](input []T) LoopedSlice[T] {
	return LoopedSlice[T]{
		s:    input,
		size: len(input),
		iter: 0,
	}
}

func (ls *LoopedSlice[T]) Next() T {
	defer func() { ls.iter++ }()
	return ls.s[ls.iter%ls.size]
}

func (n Network) Step(loc string, ls *LoopedSlice[rune]) string {
	dir := ls.Next()
	if dir == rune('R') {
		return n[loc][1]
	} else {
		return n[loc][0]
	}

}
