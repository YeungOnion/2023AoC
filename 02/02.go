package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func main() {
	filename := "02/input.txt"

	r_query := 12
	g_query := 13
	b_query := 14

	// stream file by lines
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	var power uint64 = 0
	count := 1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		r_max := findMaxOfColor(line, "red")
		g_max := findMaxOfColor(line, "green")
		b_max := findMaxOfColor(line, "blue")
		if r_max <= r_query && g_max <= g_query && b_max <= b_query {
			sum += count
		}
		power += uint64(r_max * g_max * b_max)
		count++
	}
	fmt.Println("sum is ", sum)
	fmt.Println("power is ", power)
}

func Max[C ~[]T, T constraints.Ordered](s C, null T) T {
	return lo.Reduce(s, func(m, t T, _ int) T {
		if t > m {
			return t
		} else {
			return m
		}
	}, null)
}

func findMaxOfColor(line, color string) int {
	re := regexp.MustCompile(`(\d+) ` + color)
	// will only have one submatch, so this function calls Atoi on the submatch
	convertSubmatch := func(t []string, _ int) int { v, _ := strconv.Atoi(t[1]); return v }

	matches := lo.Map(re.FindAllStringSubmatch(line, -1), convertSubmatch)

	return Max(matches, 0)
}
