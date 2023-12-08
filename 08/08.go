package main

import (
	"bufio"
	"fmt"
	"os"

	// "sort"
	// "strconv"
	"regexp"

	"github.com/samber/lo"
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

	// obtain all starts
	startPositionRe := regexp.MustCompile(`\w{2}A`)
	endPositionRe := regexp.MustCompile(`\w{2}Z`)
	endsInA := func(key string, _ int) bool { return startPositionRe.MatchString(key) }
	startPositions := lo.Filter(lo.Keys(netwk), endsInA)

	// count steps for each in isolation, then use least common multiple
	steps := make([]int, len(startPositions))
	for i, start := range startPositions {
		lr := lrPattern
		for pos := start; !endPositionRe.MatchString(pos); pos = netwk.Step(pos, &lr) {
			steps[i]++
		}
	}

	fmt.Println("all steps: ", steps)
	ghostSteps := Lcm(steps)

	fmt.Println("ghost steps: ", ghostSteps)

	return
}

func ParseFile(fs *bufio.Scanner) (LoopedSlice[rune], Network) {
	tripleLetterRe := regexp.MustCompile(`\w{2}[A-Z]`)

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

func (n Network) Step(pos string, ls *LoopedSlice[rune]) string {
	dir := ls.Next()
	if dir == rune('R') {
		return n[pos][1]
	} else {
		return n[pos][0]
	}

}

func Lcm(nums []int) int {
	multiples := lo.SliceToMap(nums, func(i int) (int, int) { return i, i })
	for len(multiples) > 1 {
		// fmt.Println("pre lcm step", multiples)
		minMult := lo.Min(lo.Keys(multiples))

		baseNum := multiples[minMult]
		delete(multiples, minMult)

		if _, ok := multiples[minMult+baseNum]; ok {
			multiples[minMult+baseNum] = minMult + baseNum
		} else {
			multiples[minMult+baseNum] = baseNum
		}
		// fmt.Println("post lcm step", multiples)
	}

	return lo.Keys(multiples)[0]
}
