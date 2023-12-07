package avl_test

import (
	"YeungOnion/2023AoC/avl"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// Define a custom comparison function for int values
	intCompare := func(a, b int) avl.Ordering {
		if a < b {
			return avl.Less
		} else if a > b {
			return avl.Greater
		}
		return avl.Equal
	}

	// Create a new avl with the custom comparison function
	tr := avl.NewBST[int](intCompare)

	// Insert values into the AVL tree
	values := []int{5, 3, 7, 1, 4, 6, 8}
	for _, value := range values {
		tr.Insert(value)
	}

	// Display the in-order traversal of the AVL-balanced binary search tree
	sortedValues := make([]int, 0)
	avl.InOrderTraversal(tr.Root, &sortedValues)

	expected := "[1 3 4 5 6 7 8]"
	got := fmt.Sprintf("%v", sortedValues)
	if got != expected {
		t.Fatalf("got: %s, expected %s", got, expected)
	}
}
