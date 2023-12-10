package main

import (
	"YeungOnion/2023AoC/utils"
	"YeungOnion/2023AoC/utils/iter"
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/samber/lo"
)

func main() {
	filename := os.Args[1]
	if filename == "--" {
		filename = os.Args[2]
	}

	fs, closingHandle := utils.FileScanner(filename)
	defer closingHandle()

	fs.Split(bufio.ScanLines)

	accum := 0
	for fs.Scan() {
		seq := ParseLine(fs.Text())
		nextVal := ExtrapFromSeq(seq)
		accum += nextVal
	}
	fmt.Println("total: ", accum)
	return
}

func ExtrapFromSeq(seq []int) int {
	accum, _ := lo.Last(seq)
	allSame := func(s []int) bool {
		return lo.EveryBy(s, func(x int) bool { return x == s[0] })
	}

	for !allSame(seq) {
		seq = FiniteDiff(seq)
		val, _ := lo.Last(seq)
		accum += val
	}

	return accum
}

func ParseLine(line string) []int {
	digitsRe := regexp.MustCompile(`\d+`)
	seqDigits := digitsRe.FindAllString(line, -1)
	return lo.Map(seqDigits, utils.MustAtoi)
}

func FiniteDiff(in []int) []int {
	if in == nil {
		return nil
	}
	return iter.Zip2With(
		in[:len(in)-1],
		in[1:],
		func(left, right int) int {
			return right - left
		},
	)
}

func FDtoConst(in []int) [][]int {
	allSame := func(s []int) bool {
		return lo.EveryBy(s, func(x int) bool { return x == s[0] })
	}
	seqs := make([][]int, 0, len(in))
	seqs = append(seqs, in)

	for i := 0; !allSame(seqs[i]); i++ {
		seqs = append(seqs, FiniteDiff(seqs[i]))
	}
	return seqs
}

func BinomialCoefMemo(bufferSize int) func(int, int) int {
	memo := make(map[int][]int, bufferSize)
	var BinCoef func(N, k int) int
	BinCoef = func(N, k int) int {
		// if row NOT memoed, allocate row as []int
		if _, ok := memo[N]; !ok {
			memo[N] = make([]int, N/2+2)
		}
		if k > N/2 {
			k = N - k
		}

		if k == 0 || N < 2 {
			return 1
		} else if k == 1 {
			return N
		} else if memo[N][k] == 0 {
			memo[N-1][k-1] = BinCoef(N-1, k-1)
			memo[N-1][k] = BinCoef(N-1, k)
			memo[N][k] = memo[N-1][k] + memo[N-1][k-1]
		}
		return memo[N][k]
	}
	return BinCoef
}

func LagrangePolyExtrap(seq []int, BinCoef func(int, int) int) int {
	k := len(seq) + 1
	if BinCoef != nil {
		BinCoef = BinomialCoefMemo(k)
	}

	coefs := lo.Map(lo.Range(k), func(j, _ int) int {
		var sign int
		if (k-j)%2 == 1 {
			sign = 1
		} else {
			sign = -1
		}
		return sign * BinCoef(k, j)
	})

	return lo.Sum(iter.Zip2With(coefs, seq, func(c, s int) int {
		return c * s
	}))
}
