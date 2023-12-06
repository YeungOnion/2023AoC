package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestAdjacentColumn(t *testing.T) {
	tests := []struct {
		in  string
		pos []int
		cmp string
		exp bool
	}{
		{
			cmp: ".......", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: false,
		},
		{
			cmp: "..123*.", // [][]int{{5,6}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: true,
		},
		{
			cmp: "*....?.", // [][]int{{0,1},{5,6}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: true,
		},
		{
			cmp: "......?", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: false,
		},
		{
			cmp: "?......", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: false,
		},
		{
			cmp: ".?.....", // [][]int {{0,1},{6,7}}
			in:  "..123*.",
			pos: []int{2, 5},
			exp: true,
		},
	}

	for _, test := range tests {
		// got := cal
		got := AnyAdjacentColumn(test.pos, test.cmp, regexp.MustCompile(`[^\d\.]`))
		if got != test.exp {
			t.Fatalf("\n%s\n%s\n", test.in, test.cmp)
			t.Fatalf("got %v, expected %v\n", got, test.exp)
		}
	}
}

func TestHasNeighborsPartA(t *testing.T) {
	tests := []struct {
		in       []string
		expected bool
	}{
		{
			in: []string{
				".......",
				"..123*.",
				"*....?.",
			},
			expected: true,
		},
		{
			in: []string{
				".......",
				"..123..",
				"*....?.",
			},
			expected: true,
		},
		{
			in: []string{
				".......",
				"..123..",
				"*.....?",
			},
			expected: false,
		},
		{
			in: []string{
				".......",
				"..123*.",
				".......",
			},
			expected: true,
		},
	}

	periphRe := regexp.MustCompile(`[^\d\.]`)
	targetRe := regexp.MustCompile(`\d+`)
	for _, test := range tests {
		// assert number of matches since testing
		index := GetTargetIndexes(test.in[1], targetRe)[0]
		if HasNeighborPeripherals(index, test.in, periphRe) != test.expected {
			out := fmt.Sprintf("expected match? = %v, assertion failed with input\n", test.expected)
			out += strings.Join(test.in, "\n")
			t.Fatal(out)
		}
	}

}

func TestCountNeighborsPartB(t *testing.T) {
	tests := []struct {
		in       []string
		expected int
	}{
		{
			in: []string{
				".......",
				"..123*.",
				".....?.",
			},
			expected: 1,
		},
		{
			in: []string{
				".......",
				"*....?.",
				"..123..",
			},
			expected: 0,
		},
		{
			in: []string{
				".132...",
				"*.....?",
				"..123..",
			},
			expected: 1,
		},
		{
			in: []string{
				"..123..",
				"..123*.",
				".......",
			},
			expected: 2,
		},
		{
			in: []string{
				"..123.1",
				"..123*.",
				".......",
			},
			expected: 3,
		},
	}

	periphRe := regexp.MustCompile(`\d+`)
	targetRe := regexp.MustCompile(`\*`)
	for _, test := range tests {
		// assert number of matches since testing
		index := GetTargetIndexes(test.in[1], targetRe)[0]
		got := CountNeighborPeripherals(index, test.in, periphRe)
		if got != test.expected {
			out := fmt.Sprintf("expected %d, got %d\n", test.expected, got)
			out += strings.Join(test.in, "\n")
			t.Fatal(out)
		} else {
			t.Log("passed case")
		}
	}

}

func TestNeighborPeripherals(t *testing.T) {
	tests := []struct {
		pos      []int
		in       []string
		expected [][]int
	}{
		{
			pos: []int{5, 6},
			in: []string{
				".......",
				"..123*.",
				".....?.",
			},
			expected: [][]int{{2, 5, 1}},
		},
		{
			pos: []int{0, 1},
			in: []string{
				".......",
				"*....?.",
				"..123..",
			},
			expected: [][]int{},
		},
		{
			pos: []int{0, 1},
			in: []string{
				".132...",
				"*.....?",
				"..123..",
			},
			expected: [][]int{{1, 4, 0}},
		},
		{
			pos: []int{5, 6},
			in: []string{
				"..123..",
				"..123*.",
				".......",
			},
			expected: [][]int{{2, 5, 0}, {2, 5, 1}},
		},
		{
			pos: []int{5, 6},
			in: []string{
				"..123..",
				"..123*.",
				"..123..",
			},
			expected: [][]int{{2, 5, 0}, {2, 5, 1}, {2, 5, 2}},
		},
		{
			pos: []int{3, 4},
			in: []string{
				"467..114",
				"...*....",
				"..35..63",
			},
			expected: [][]int{{0, 3, 0}, {2, 4, 2}},
		},
	}

	for i, test := range tests {
		periphRe := regexp.MustCompile(`\d+`)
		got := NeighborPeripherals(test.pos, test.in, periphRe)
		// t.Logf("case %d\n\tgot     : %v\n\texpected: %v", i, got, test.expected)
		if len(got) != len(test.expected) {
			t.Fatal("number of matches not as expected")
		}
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", test.expected) {
			t.Fatalf("case %d\n\tgot     : %v\n\texpected: %v", i, got, test.expected)
		}
	}

}

func TestSamplePartA(t *testing.T) {
	exp := 4361
	infile := "sample.txt"

	// stream file by words
	file, err := os.Open(infile)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	gotPart, _ := ScanAndScore(fileScanner)
	if exp != gotPart {
		t.Fatalf("expected %d, got %d", exp, gotPart)
	}

}

func TestSamplePartB(t *testing.T) {
	exp := 467835
	infile := "sample-b.txt"

	// stream file by words
	file, err := os.Open(infile)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	_, gotGear := ScanAndScore(fileScanner)
	if exp != gotGear {
		t.Fatalf("expected %d, got %d", exp, gotGear)
	}

}

func TestPartA(t *testing.T) {
	exp := 539590
	infile := "input.txt"

	// stream file by words
	file, err := os.Open(infile)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	gotPart, _ := ScanAndScore(fileScanner)
	if exp != gotPart {
		t.Fatalf("expected %d, got %d", exp, gotPart)
	}
}
