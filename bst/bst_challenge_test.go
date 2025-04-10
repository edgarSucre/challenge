package bst_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/challenge/bst"
	"github.com/stretchr/testify/assert"
)

func getBST(t *testing.T, values ...int) *bst.Node[int] {
	t.Helper()

	var root *bst.Node[int]

	for _, v := range values {
		root = bst.Insert(root, v)
	}

	return root
}

func getLinkedList(t *testing.T, values ...int) *bst.List {
	t.Helper()

	var (
		head *bst.List
		cur  *bst.List
	)

	for _, v := range values {
		if head == nil {
			head = &bst.List{Val: v}
			cur = head

			continue
		}

		cur.Next = &bst.List{Val: v}
		cur = cur.Next
	}

	return head
}

func assertExpected(t *testing.T, expected any, actual *bst.Node[int]) {
	t.Helper()

	if expected == nil {
		assert.Nil(t, actual)
	} else {
		if actual == nil {
			t.Errorf("expected: %v, actual: %v\n", expected, actual)
		}

		assert.Equal(t, expected, actual.Val)
	}
}

func TestInOrderMorrisList(t *testing.T) {
	root := getBST(t, 4, 2, 1, 3, 5)

	expected := []int{1, 2, 3, 4, 5}
	actual := bst.InOrderMorrisList(root)

	assert.Equal(t, expected, actual)
}

func TestInOrderSuccessor(t *testing.T) {
	inOrderSuccessor := bst.InOrderSuccessor{Root: getBST(t, 4, 2, 1, 3, 5)}

	tests := []struct {
		target   int
		expected any
	}{
		{2, 3},
		{1, 2},
		{3, 4},
		{4, 5},
		{5, nil},
		{9, nil},
		{0, nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Morris(%v)", tt.target), func(t *testing.T) {
			actual := inOrderSuccessor.Morris(tt.target)
			assertExpected(t, tt.expected, actual)
		})

		t.Run(fmt.Sprintf("Recursive(%v)", tt.target), func(t *testing.T) {
			actual := inOrderSuccessor.Recursive(tt.target)
			assertExpected(t, tt.expected, actual)
		})

		t.Run(fmt.Sprintf("ReverseRecursive(%v)", tt.target), func(t *testing.T) {
			actual := inOrderSuccessor.ReverseRecursive(tt.target)
			assertExpected(t, tt.expected, actual)
		})

		t.Run(fmt.Sprintf("BST(%v)", tt.target), func(t *testing.T) {
			actual := inOrderSuccessor.BST(tt.target)
			assertExpected(t, tt.expected, actual)
		})
	}
}

func TestInOrderPredecessor(t *testing.T) {
	inOrderPredecessor := bst.InOrderPredecessor{Root: getBST(t, 4, 2, 1, 3, 5)}

	tests := []struct {
		target   int
		expected any
	}{
		{3, 2},
		{2, 1},
		{4, 3},
		{5, 4},
		{1, nil},
		{9, nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("ReverseMorris(%v)", tt.target), func(t *testing.T) {
			actual := inOrderPredecessor.ReverseMorris(tt.target)
			assertExpected(t, tt.expected, actual)
		})

		t.Run(fmt.Sprintf("ReverseRecursive(%v)", tt.target), func(t *testing.T) {
			actual := inOrderPredecessor.ReverseRecursive(tt.target)
			assertExpected(t, tt.expected, actual)
		})

		t.Run(fmt.Sprintf("BST(%v)", tt.target), func(t *testing.T) {
			actual := inOrderPredecessor.BST(tt.target)
			assertExpected(t, tt.expected, actual)
		})
	}
}

func TestSumSmallest(t *testing.T) {
	t1 := bst.Insert(nil, 8)
	t1 = bst.Insert(t1, 7)
	t1 = bst.Insert(t1, 2)
	t1 = bst.Insert(t1, 10)
	t1 = bst.Insert(t1, 9)
	t1 = bst.Insert(t1, 13)

	assert.Equal(t, bst.SumSmallest(t1, 3), 2)

	t2 := bst.Insert(nil, 8)
	t2 = bst.Insert(t2, 5)
	t2 = bst.Insert(t2, 2)
	t2 = bst.Insert(t2, 3)
	t2 = bst.Insert(t2, 7)
	t2 = bst.Insert(t2, 11)

	assert.Equal(t, bst.SumSmallest(t2, 5), 10)
}

func TestSumNthSmallest(t *testing.T) {
	root := getBST(t, 8, 7, 2, 10, 9, 13)
	sumKth := bst.SumKthSmallest{Root: root}

	tests := []struct {
		kth      int
		expected int
	}{
		{0, 0},
		{1, 2},
		{2, 9},
		{3, 17},
		{4, 26},
		{5, 36},
		{6, 49},
		{7, 49},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Recursive(%v)", tt.kth), func(t *testing.T) {
			assert.Equal(t, tt.expected, sumKth.Recursive(tt.kth))
		})

		t.Run(fmt.Sprintf("Morris(%v)", tt.kth), func(t *testing.T) {
			assert.Equal(t, tt.expected, sumKth.Morris(tt.kth))
		})
	}
}

func TestSecondLargest(t *testing.T) {
	root := getBST(t, 8, 7, 2, 10, 9, 13)
	kthLargest := bst.KthLargest{Root: root}

	tests := []struct {
		kth      int
		expected int
	}{
		{1, 13},
		{2, 10},
		{3, 9},
		{4, 8},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Recursive(%v)", tt.kth), func(t *testing.T) {
			actual := kthLargest.Recursive(tt.kth)
			assert.Equal(t, tt.expected, actual)
		})

		t.Run(fmt.Sprintf("Morris(%v)", tt.kth), func(t *testing.T) {
			actual := kthLargest.Morris(tt.kth)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestBSTKeysInRange(t *testing.T) {
	t1 := bst.Insert(nil, 22)
	t1 = bst.Insert(t1, 12)
	t1 = bst.Insert(t1, 8)
	t1 = bst.Insert(t1, 20)
	t1 = bst.Insert(t1, 30)

	//got := bst.BSTKeysInRangeRecursion(t1, 10, 22)
	got := bst.KeysInRangeMorris(t1, 10, 22)
	want := []int{12, 20, 22}

	assert.Equal(t, want, got)

	t2 := bst.Insert(nil, 22)
	t2 = bst.Insert(t2, 12)
	t2 = bst.Insert(t2, 8)
	t2 = bst.Insert(t2, 20)
	t2 = bst.Insert(t2, 30)

	got = bst.KeysInRangeRecursion(t2, 1, 10)
	want = []int{8}

	assert.Equal(t, want, got)
}

func TestBalanceBST(t *testing.T) {
	root := getBST(t, 7, 6, 5, 4, 3, 2, 1)

	balanced := bst.BalanceBST(root)
	_ = balanced
}

func TestBinaryTreeToBST(t *testing.T) {
	root := &bst.Node[int]{Val: 10}

	root.Left = &bst.Node[int]{Val: 2}
	root.Right = &bst.Node[int]{Val: 7}
	root.Left.Left = &bst.Node[int]{Val: 8}
	root.Left.Right = &bst.Node[int]{Val: 4}

	res := bst.BinaryTreeToBST(root)
	_ = res
}

func TestAddGreater(t *testing.T) {
	root := getBST(t, 50, 30, 20, 40, 70, 60, 80)

	bst.AddGreater(root, 0)

	assert.True(t, true)
}

func TestContainsSameElements(t *testing.T) {
	root1 := getBST(t, 15, 10, 5, 12, 20, 25)
	root2 := getBST(t, 15, 12, 5, 10, 20, 25)

	got := bst.ContainsSameElements(root1, root2)

	assert.True(t, got)
}

func TestConstructBSTFromPreorder(t *testing.T) {
	pre := []int{10, 5, 1, 7, 40, 50}

	root := bst.ConstructBSTFromPreorder(pre)
	_ = root
}

func TestSortedLinkedListToBalancedBST(t *testing.T) {
	head := getLinkedList(t, 1, 2, 3, 4, 5, 6, 7, 8)

	root := bst.SortedLinkedListToBalancedBST(head)

	_ = root
}
