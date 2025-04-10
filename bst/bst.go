package bst

import (
	"golang.org/x/exp/constraints"
)

type comp constraints.Ordered

type Node[T comp] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func Insert[T comp](node *Node[T], val T) *Node[T] {
	if node == nil {
		return newNode(val)
	}

	if node.Val > val {
		node.Left = Insert(node.Left, val)
	} else {
		node.Right = Insert(node.Right, val)
	}

	return node
}

func newNode[T comp](val T) *Node[T] {
	return &Node[T]{Val: val}
}
