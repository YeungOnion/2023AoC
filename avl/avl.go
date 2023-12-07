package avl

import (
	"fmt"

	"golang.org/x/exp/constraints"
	"strings"
)

// Node represents a node in the binary search tree with AVL balancing.
type Node[T any] struct {
	Value  T
	Left   *Node[T]
	Right  *Node[T]
	Height int
}

// BST represents a binary search tree with AVL balancing.
type BST[T any] struct {
	Root        *Node[T]
	compareFunc CompareFunc[T]
}

type Ordering int

const (
	Less = iota - 1
	Equal
	Greater
)

// CompareFunc is a custom function to compare two values.
type CompareFunc[T any] func(a, b T) Ordering

// NewBST creates a new BST with a custom comparison function.
func NewBST[T any](cmp CompareFunc[T]) *BST[T] {
	return &BST[T]{compareFunc: cmp}
}

func New[T constraints.Ordered]() *BST[T] {
	return &BST[T]{compareFunc: func(a, b T) Ordering {
		switch {
		case a == b:
			return Equal
		case a < b:
			return Greater
		case a > b:
			return Less
		default:
			panic("unreachable")
		}

	}}
}

func (t BST[T]) String() string {
	vals := make([]T, 0, 10)
	InOrderTraversal(t.Root, &vals)
	lines := make([]string, len(vals))
	for i, v := range vals {
		lines[i] = fmt.Sprintf("%v", v)
	}

	return strings.Join(lines, "\n")
}

// FloorSearch searches and yields the "floor" of the search
// The value in the returned node is the largest that is not more than the search value
func (tr *BST[T]) FloorSearch(value T) *Node[T] {
	return floorSearchNode(tr.Root, value, tr.compareFunc)
}

// Search searches for a value in the BST
func (tr *BST[T]) Search(value T) *Node[T] {
	return searchNode(tr.Root, value, tr.compareFunc)
}

// Insert inserts a value into the AVL-balanced binary search tree.
func (tr *BST[T]) Insert(value T) {
	tr.Root = insertNode(tr.Root, value, tr.compareFunc)
}

// height returns the height of the node.
func height[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// balanceFactor returns the balance factor of the node.
func balanceFactor[T any](node *Node[T]) int {
	return height(node.Left) - height(node.Right)
}

// rotateRight performs a right rotation on the node.
func rotateRight[T any](y *Node[T]) *Node[T] {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

// rotateLeft performs a left rotation on the node.
func rotateLeft[T any](x *Node[T]) *Node[T] {
	y := x.Right
	T2 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T2

	// Update heights
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

// insertNode inserts a value into the AVL-balanced binary search tree.
func insertNode[T any](root *Node[T], value T, compareFunc CompareFunc[T]) *Node[T] {
	// Perform standard BST insert
	if root == nil {
		return &Node[T]{Value: value, Height: 1}
	}

	if compareFunc(value, root.Value) < Equal {
		root.Left = insertNode(root.Left, value, compareFunc)
	} else if compareFunc(value, root.Value) > Equal {
		root.Right = insertNode(root.Right, value, compareFunc)
	} else {
		// Duplicate values are not allowed
		return root
	}

	// Update height of the current node
	root.Height = 1 + max(height(root.Left), height(root.Right))

	// Get the balance factor
	balance := balanceFactor(root)

	// Perform rotations if necessary to maintain balance
	// Left Left Case
	if balance > 1 && compareFunc(value, root.Left.Value) < Equal {
		return rotateRight(root)
	}

	// Right Right Case
	if balance < -1 && compareFunc(value, root.Right.Value) > Equal {
		return rotateLeft(root)
	}

	// Left Right Case
	if balance > 1 && compareFunc(value, root.Left.Value) > Equal {
		root.Left = rotateLeft(root.Left)
		return rotateRight(root)
	}

	// Right Left Case
	if balance < -1 && compareFunc(value, root.Right.Value) < Equal {
		root.Right = rotateRight(root.Right)
		return rotateLeft(root)
	}

	return root
}

// InOrderTraversal performs in-order traversal of the binary search tree.
func InOrderTraversal[T any](root *Node[T], values *[]T) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left, values)
	*values = append(*values, root.Value)
	InOrderTraversal(root.Right, values)
}

// searchNode searches for a value in the AVL tree.
func searchNode[T any](root *Node[T], value T, compareFunc CompareFunc[T]) *Node[T] {
	if root == nil || compareFunc(value, root.Value) == Equal {
		return root
	}

	if compareFunc(value, root.Value) == Less {
		return searchNode(root.Left, value, compareFunc)
	}

	return searchNode(root.Right, value, compareFunc)
}

func floorSearchNode[T any](root *Node[T], value T, compareFunc CompareFunc[T]) *Node[T] {
	if root == nil {
		return root
	}

	switch compareFunc(value, root.Value) {
	case Equal:
		return root
	case Less:
		if f := floorSearchNode(root.Left, value, compareFunc); f == nil {
			return nil
		} else {
			return f
		}
	case Greater:
		if f := floorSearchNode(root.Right, value, compareFunc); f == nil {
			return root.Right
		} else {
			return f
		}
	default:
		panic("comparison yielded unexpected value")
	}
}
