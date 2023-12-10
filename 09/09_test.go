package main

import (
	"YeungOnion/2023AoC/utils"
	"bufio"
	"fmt"
	"testing"
)

func TestSampleA(t *testing.T) {
	filename := "sample-a.txt"

	fs, closingHandle := utils.FileScanner(filename)
	defer closingHandle()

	fs.Split(bufio.ScanLines)

}

func TestFiniteDifference(t *testing.T) {
	seq := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
	}
	exp := [][]int{
		{3, 3, 3, 3, 3},
		{2, 3, 4, 5, 6},
	}

	for i, s := range seq {
		got := FiniteDiff(s)
		expected := exp[i]
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
			t.Fatalf("not a match, got %v, expected %v", got, expected)
		}
	}

	t.Log(FDtoConst(seq[0]))
}

func TestBinCoef(t *testing.T) {
	BinCoef := BinomialCoefMemo(10)

	tests := []struct {
		N   int
		k   int
		exp int
	}{
		{
			N:   5,
			k:   0,
			exp: 1,
		},
		{
			N:   5,
			k:   1,
			exp: 5,
		},
		{
			N:   5,
			k:   2,
			exp: 10,
		},
		{
			N:   10,
			k:   2,
			exp: 45,
		},
		{
			N:   10,
			k:   4,
			exp: 210,
		},
		{
			N:   10,
			k:   5,
			exp: 252,
		},
		{
			N:   40,
			k:   10,
			exp: 847660528,
		},
	}

	for _, test := range tests {
		if got := BinCoef(test.N, test.k); got != test.exp {
			t.Logf("Bin(N=%d,k=%d)", test.N, test.k)
			t.Fatalf("got %d, expected %d", got, test.exp)
		}
	}
}

func TestLPExtrap(t *testing.T) {
	seq := []int{1, 3, 6, 10, 15, 21}
	bc := BinomialCoefMemo(len(seq))
	for i := 2; i < len(seq); i++ {
		s := seq[:i]
		got := LagrangePolyExtrap(s, bc)
		t.Logf("%v -> %d", s, got)
	}
	return
}
