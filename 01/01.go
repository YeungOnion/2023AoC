package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var forwardDigitRe, reverseDigitRe string

func main() {
	filename := "01/input.txt"

	// stream file by lines
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	// send each line to channel to get digits
	initBothRegex()
	lineValues := make([]int, 0, 100)
	for fileScanner.Scan() {
		s := fileScanner.Text()
		if s != "" {
			lineValues = append(lineValues, extractNumberByDigitOrDigitName(fileScanner.Text()))
		}
	}

	// sum
	sum := 0
	for _, v := range lineValues {
		sum += v
	}

	fmt.Println(sum)
}

func Map[T, V any](t []T, fn func(T) V) []V {
	v := make([]V, len(t))
	for i, elem := range t {
		v[i] = fn(elem)
	}
	return v
}

func initBothRegex() {
	digitNames := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitNamesReverse := Map(digitNames, func(fwd string) string { return reverseString(fwd) })

	forwardDigitRe = strings.Join(digitNames, "|") + `|\d`
	reverseDigitRe = strings.Join(digitNamesReverse, "|") + `|\d`
}

func canonicalizeDigit(s string) (string, error) {
	switch s {
	case "0", "zero", "orez":
		return "0", nil
	case "1", "one", "eno":
		return "1", nil
	case "2", "two", "owt":
		return "2", nil
	case "3", "three", "eerht":
		return "3", nil
	case "4", "four", "ruof":
		return "4", nil
	case "5", "five", "evif":
		return "5", nil
	case "6", "six", "xis":
		return "6", nil
	case "7", "seven", "neves":
		return "7", nil
	case "8", "eight", "thgie":
		return "8", nil
	case "9", "nine", "enin":
		return "9", nil
	default:
		return "", errors.New(fmt.Sprintf("%s is not a valid digit or digitname", s))
	}
}

func reverseSliceInPlace[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	reverseSliceInPlace(runes)
	return string(runes)
}

// findDigitIfAny will return a string with a single character for a base 10 digit
// based on the regexp provided. Note that a conversion must exist from the regex match
// to the single character string for the digit
// this function will panic if one is not found via the regex or cannot convert
func findDigitIfAny(line string, re regexp.Regexp) string {

	lineCopy := line
	if found := re.FindString(line); found == "" {
		panic(fmt.Sprintf("%s did not find match", lineCopy))
	} else {
		v, err := canonicalizeDigit(found)
		if err != nil {
			panic(fmt.Sprintf("error processing: %s\n%s", lineCopy, err))
		} else {
			return v
		}
	}
}

func extractNumberByDigitOrDigitName(line string) int {
	var digits string

	digits = findDigitIfAny(line, *regexp.MustCompile(forwardDigitRe))
	digits += findDigitIfAny(reverseString(line), *regexp.MustCompile(reverseDigitRe))

	value, err := strconv.Atoi(digits)

	if err != nil {
		fmt.Printf("while processing input:\n\t>> %s", digits)
		panic(err)
	}
	return value

}
