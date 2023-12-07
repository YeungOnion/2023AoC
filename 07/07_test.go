package main

import (
	"testing"
)

func TestCardScore(t *testing.T) {
	c := []rune{'2', '3', '5', '7', 'T', 'K', 'Q', 'J', 'A'}
	v := []int{2, 3, 5, 7, 10, 13, 12, 11, 14}

	for i := range c {
		if got := EvaluateCard(c[i]); got != v[i] {
			t.Fatalf("got %d, expected %d", got, v[i])
		}
	}
	return
}

func TestHandType(t *testing.T) {
	tests := []struct {
		name     string
		hand     Hand
		expected HandType
	}{
		{
			hand:     "12341",
			expected: OnePair,
		},
		{
			hand:     "11243",
			expected: OnePair,
		},
		{
			hand:     "12331",
			expected: TwoPair,
		},
		{
			hand:     "11223",
			expected: TwoPair,
		},
		{
			hand:     "11134",
			expected: ThreeKind,
		},
		{
			hand:     "12121",
			expected: FullHouse,
		},
		{
			hand:     "1T111",
			expected: FourKind,
		},
		{
			hand:     "11111",
			expected: FiveKind,
		},
	}

	for _, test := range tests {
		if got := EvaluateHandType(test.hand); got != test.expected {
			t.Fatalf("got: %d, expected: %d", got, test.expected)
		}
	}
}

func TestLess(t *testing.T) {
	tests := []struct {
		name     string
		left     Hand
		right    Hand
		expected bool
	}{
		{
			left:     "12341",
			right:    "12341",
			expected: false,
		},
		{
			left:     "T55J5",
			right:    "QQQJA",
			expected: true,
		},
	}

	for _, test := range tests {
		if got := Less(test.left, test.right); got != test.expected {
			t.Fatalf("ordering incorrect, expected that %s < %s\n", test.left, test.right)
		}
	}
}
