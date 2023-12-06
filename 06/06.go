package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

var forwardDigitRe, reverseDigitRe string

func main() {
	filename := os.Args[1]
	if filename == "--" {
		filename = os.Args[2]
	}
	digitsRe := regexp.MustCompile(`\d+`)

	// stream file by lines
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	times := lo.Map(digitsRe.FindAllString(fileScanner.Text(), -1), MustAtoi)

	fileScanner.Scan()
	distances := lo.Map(digitsRe.FindAllString(fileScanner.Text(), -1), MustAtoi)

	fmt.Printf("times: %v\ndists: %v\n", times, distances)
	numRaces := len(times)

	numWinning := make([]int, numRaces)
	for i := 0; i < numRaces; i++ {
		numWinning[i] = countSolns(times[i], distances[i])
	}

	fmt.Println("numWins", numWinning)
	fmt.Println(lo.Reduce(numWinning, Prod, 1))

	return
}

func MustAtoi(s string, _ int) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func Prod(prod int, i int, _ int) int {
	return prod * i
}

func countSolns(raceDuration int, record int) int {
	isSoln := func(t int) bool { return (raceDuration-t)*t > record }

	T, dRecord := float64(raceDuration), float64(record)
	delta := math.Sqrt(T*T/4 - dRecord)
	lb, ub := T/2-delta, T/2+delta
	fmt.Printf("lb,ub = (%.2f, %.2f)\n", lb, ub)
	numWins := int(math.Floor(ub) - math.Ceil(lb) + 1)
	if !isSoln(int(math.Floor(ub))) {
		numWins--
	}
	if !isSoln(int(math.Ceil(lb))) {
		numWins--
	}
	// if isSoln(int(math.Floor(ub + 1))) {
	// 	numWins++
	// }
	// if isSoln(int(math.Ceil(lb - 1))) {
	// 	numWins++
	// }
	return numWins
}
