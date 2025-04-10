package bst

import (
	"math"
	"sort"
)

func InOrderMorrisList(root *Node[int]) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	cur := root

	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
		} else {
			pred := cur.Left
			for pred.Right != nil && pred.Right != cur {
				pred = pred.Right
			}

			if pred.Right == nil {
				pred.Right = cur
				cur = cur.Left
			} else {
				pred.Right = nil
				res = append(res, cur.Val)
				cur = cur.Right
			}
		}
	}

	return res
}

func rightMMost(current *Node[int]) *Node[int] {
	for current.Right != nil {
		current = current.Right
	}

	return current
}

func leftMost(current *Node[int]) *Node[int] {
	for current.Left != nil {
		current = current.Left
	}

	return current
}

// =================================================================================================
// =====================================IN ORDER SUCCESSOR==========================================
// =================================================================================================

type InOrderSuccessor struct {
	Root *Node[int]
}

// Morris is for binary tree not BST
// Morris gives you O(n) Time and O(1) Space
func (s InOrderSuccessor) Morris(target int) *Node[int] {
	if s.Root == nil {
		return nil
	}

	if s.Root.Val == target && s.Root.Right != nil {
		return leftMost(s.Root.Right)
	}

	var (
		successor   *Node[int]
		targetFound bool
	)

	cur := s.Root
	for cur != nil {
		if cur.Left == nil {
			if targetFound {
				successor = cur
				targetFound = false
			}

			if cur.Val == target {
				targetFound = true
			}

			cur = cur.Right
		} else {
			pred := cur.Left
			for pred.Right != nil && pred.Right != cur {
				pred = pred.Right
			}

			if pred.Right == nil {
				pred.Right = cur
				cur = cur.Left
			} else {
				pred.Right = nil

				if targetFound {
					successor = cur
					targetFound = false
				}

				if cur.Val == target {
					targetFound = true
				}

				cur = cur.Right
			}
		}
	}

	return successor
}

// Similar to Morris this method is for binary tree not BST
// Recursive approach gives you O(n) Time and O(h) Space
func (s InOrderSuccessor) Recursive(target int) *Node[int] {
	var targetFound bool

	return recursiveSuccessor(s.Root, target, &targetFound)
}

func recursiveSuccessor(root *Node[int], target int, targetFound *bool) *Node[int] {
	if root == nil {
		return nil
	}

	suc := recursiveSuccessor(root.Left, target, targetFound)
	if suc != nil {
		return suc
	}

	if *targetFound {
		return root
	}

	if root.Val == target {
		*targetFound = true
	}

	return recursiveSuccessor(root.Right, target, targetFound)
}

// Similar to Recursive but uses a Reverse InOrder traverse
func (s InOrderSuccessor) ReverseRecursive(target int) *Node[int] {
	last := []*Node[int]{nil}

	return reverseRecursiveSuccessor(s.Root, target, last)
}

func reverseRecursiveSuccessor(root *Node[int], target int, last []*Node[int]) *Node[int] {
	if root == nil {
		return root
	}

	suc := reverseRecursiveSuccessor(root.Right, target, last)
	if suc != nil {
		return suc
	}

	if root.Val == target {
		return last[0]
	}

	last[0] = root

	return reverseRecursiveSuccessor(root.Left, target, last)
}

func (s InOrderSuccessor) BST(target int) *Node[int] {
	if s.Root == nil {
		return nil
	}

	if s.Root.Val == target && s.Root.Right != nil {
		return leftMost(s.Root.Right)
	}

	var (
		successor   *Node[int]
		targetFound bool
	)

	cur := s.Root
	for cur != nil {
		if cur.Val == target {
			targetFound = true
		}

		if cur.Val > target {
			successor = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if targetFound {
		return successor
	}

	return nil
}

// *************************************************************************************************

// =================================================================================================
// ====================================IN ORDER PREDECESSOR=========================================
// =================================================================================================

type InOrderPredecessor struct{ Root *Node[int] }

func (s InOrderPredecessor) ReverseMorris(target int) *Node[int] {
	if s.Root == nil {
		return nil
	}

	if s.Root.Val == target && s.Root.Left != nil {
		return rightMMost(s.Root.Left)
	}

	var (
		predecessor *Node[int]
		targetFound bool
	)

	cur := s.Root
	for cur != nil {
		if cur.Right == nil {
			if targetFound {
				predecessor = cur
				targetFound = false
			}

			if cur.Val == target {
				targetFound = true
			}

			cur = cur.Left
		} else {
			suc := cur.Right
			for suc.Left != nil && suc.Left != cur {
				suc = suc.Left
			}

			if suc.Left == nil {
				suc.Left = cur
				cur = cur.Right
			} else {
				if targetFound {
					predecessor = cur
					targetFound = false
				}

				if cur.Val == target {
					targetFound = true
				}

				suc.Left = nil
				cur = cur.Left
			}
		}
	}

	return predecessor
}

func (s InOrderPredecessor) ReverseRecursive(target int) *Node[int] {
	var targetFound bool

	return reverseRecursivePredecessor(s.Root, target, &targetFound)
}

func reverseRecursivePredecessor(root *Node[int], target int, targetFound *bool) *Node[int] {
	if root == nil {
		return nil
	}

	pred := reverseRecursivePredecessor(root.Right, target, targetFound)
	if pred != nil {
		return pred
	}

	if *targetFound {
		return root
	}

	if root.Val == target {
		*targetFound = true
	}

	return reverseRecursivePredecessor(root.Left, target, targetFound)
}

func (s InOrderPredecessor) BST(target int) *Node[int] {
	if s.Root == nil {
		return nil
	}

	if s.Root.Val == target && s.Root.Left != nil {
		return rightMMost(s.Root.Left)
	}

	var (
		pred        *Node[int]
		targetFound bool
	)

	cur := s.Root
	for cur != nil {
		if cur.Val == target {
			targetFound = true
		}

		if cur.Val < target {
			pred = cur
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	if targetFound {
		return pred
	}

	return nil
}

// *************************************************************************************************

// =================================================================================================
// ======================================SUM Kth Smallest===========================================
// ==========================================================1=======================================

type SumKthSmallest struct {
	Root *Node[int]
}

func (s SumKthSmallest) Recursive(kth int) int {
	if s.Root == nil {
		return 0
	}

	return sumKthSmallest(s.Root, &kth, 0)
}

func sumKthSmallest(root *Node[int], kth *int, sum int) int {
	if root == nil || *kth <= 0 {
		return sum
	}

	sum = sumKthSmallest(root.Left, kth, sum)

	if *kth > 0 {
		sum += root.Val
		*kth--
	}

	return sumKthSmallest(root.Right, kth, sum)
}

func (s SumKthSmallest) Morris(kth int) int {
	if s.Root == nil || kth <= 0 {
		return 0
	}

	var sum int

	cur := s.Root
	for cur != nil {
		if cur.Left == nil {
			if kth > 0 {
				sum += cur.Val
				kth--
			}

			cur = cur.Right
		} else {
			pred := cur.Left
			for pred.Right != nil && pred.Right != cur {
				pred = pred.Right
			}

			if pred.Right == nil {
				pred.Right = cur
				cur = cur.Left
			} else {
				if kth > 0 {
					sum += cur.Val
					kth--
				}

				pred.Right = nil
				cur = cur.Right
			}
		}
	}

	return sum
}

// *************************************************************************************************

// =================================================================================================
// ======================================Second Largest=============================================
// =================================================================================================

type KthLargest struct {
	Root *Node[int]
}

func (s KthLargest) Recursive(kth int) int {

	return kthLargest(s.Root, &kth, 0)
}

func kthLargest(root *Node[int], kth *int, res int) int {
	if root == nil || *kth <= 0 {
		return res
	}

	res = kthLargest(root.Right, kth, res)

	*kth--
	if *kth == 0 {
		return root.Val
	}

	return kthLargest(root.Left, kth, res)
}

func (s KthLargest) Morris(kth int) int {
	if s.Root == nil {
		return 0
	}

	res := 0

	cur := s.Root
	for cur != nil {
		if cur.Right == nil {
			kth--

			if kth == 0 {
				res = cur.Val
			}

			cur = cur.Left
		} else {
			suc := cur.Right
			for suc.Left != nil && suc.Left != cur {
				suc = suc.Left
			}

			if suc.Left == nil {
				suc.Left = cur
				cur = cur.Right
			} else {
				suc.Left = nil

				kth--

				if kth == 0 {
					res = cur.Val
				}

				cur = cur.Left
			}
		}
	}

	return res
}

// *************************************************************************************************

func SumSmallest(node *Node[int], target int) int {
	if node == nil {
		return 0
	}

	if node.Val <= target {
		return node.Val + SumSmallest(node.Left, target) + SumSmallest(node.Right, target)
	}

	return SumSmallest(node.Left, target)
}

func KeysInRangeRecursion(node *Node[int], n1, n2 int) []int {
	res := make([]int, 0)

	keysInRangeInOrder(node, n1, n2, &res)

	return res
}

func keysInRangeInOrder(node *Node[int], n1, n2 int, res *[]int) {
	if node == nil {
		return
	}

	if n1 <= node.Val && node.Val <= n2 {
		keysInRangeInOrder(node.Left, n1, n2, res)
		*res = append(*res, node.Val)
		keysInRangeInOrder(node.Right, n1, n2, res)
	} else if n1 > node.Val {
		keysInRangeInOrder(node.Right, n1, n2, res)
	} else {
		keysInRangeInOrder(node.Left, n1, n2, res)
	}
}

func KeysInRangeMorris(node *Node[int], low, high int) []int {
	res := make([]int, 0)

	current := node

	for current != nil {
		if current.Left == nil {
			if low <= current.Val && current.Val <= high {
				res = append(res, current.Val)
			}

			current = current.Right
		} else {
			pre := current.Left
			for pre.Right != nil && pre.Right != current {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = current
				current = current.Left
			} else {
				pre.Right = nil
				if low <= current.Val && current.Val <= high {
					res = append(res, current.Val)
				}
				current = current.Right
			}

		}
	}

	return res
}

func BalanceBST(root *Node[int]) *Node[int] {
	if root == nil {
		return nil
	}

	arr := make([]int, 0)
	bstToSortedArray(root, &arr)

	var (
		balanced *Node[int]
	)

	balanced = sortedArrayToBalancedBST(arr, 0, len(arr)-1)

	return balanced
}

func bstToSortedArray(root *Node[int], res *[]int) {
	if root == nil {
		return
	}

	bstToSortedArray(root.Left, res)
	*res = append(*res, root.Val)
	bstToSortedArray(root.Right, res)
}

func sortedArrayToBalancedBST(arr []int, start, end int) *Node[int] {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2

	root := newNode(arr[mid])
	root.Left = sortedArrayToBalancedBST(arr, start, mid-1)
	root.Right = sortedArrayToBalancedBST(arr, mid+1, end)

	return root
}

// func sortedArrayToBalancedBST(arr []int) *Node[int] {
// 	if len(arr) == 0 {
// 		return nil
// 	}

// 	mid := len(arr) / 2

// 	root := newNode(arr[mid])
// 	root.Left = sortedArrayToBalancedBST(arr[:mid])
// 	root.Right = sortedArrayToBalancedBST(arr[(mid + 1):])

// 	return root
// }

type IsBST struct{ Root *Node[int] }

func (s IsBST) Recursive() bool {
	if s.Root == nil {
		return true
	}

	return false
}

func isBST(root *Node[int]) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftVal, leftIsBST := isBST(root.Left)
	if !leftIsBST {
		return 0, false
	}

	if leftVal > root.Val {
		return 0, false
	}

	return isBST(root.Right)
}

func BinaryTreeToBST(root *Node[int]) *Node[int] {
	if root == nil {
		return nil
	}

	arr := make([]int, 0)
	bstToSortedArray(root, &arr)

	sort.Ints(arr)

	return sortedArrayToBST(root, arr)
}

// unlike sortedArrayToBalancedBST sortedArrayToBST does not guarantee a balanced BST
func sortedArrayToBST(root *Node[int], arr []int) *Node[int] {
	if root == nil {
		return nil
	}

	var index int

	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur.Val = arr[index]
			index++

			cur = cur.Right
		} else {
			pred := cur.Left
			for pred.Right != nil && pred.Right != cur {
				pred = pred.Right
			}

			if pred.Right == nil {
				pred.Right = cur
				cur = cur.Left
			} else {
				cur.Val = arr[index]
				index++

				pred.Right = nil
				cur = cur.Right
			}
		}
	}

	return root
}

func AddGreater(root *Node[int], sum int) int {
	if root == nil {
		return sum
	}

	sum = AddGreater(root.Right, sum)

	sum += root.Val
	root.Val = sum

	return AddGreater(root.Left, sum)
}

func ContainsSameElements(root1, root2 *Node[int]) bool {
	set := make(map[int]struct{})

	inOrderSetMap(root1, set)
	inOrderCheckMap(root2, set)

	return len(set) == 0
}

func inOrderSetMap(root *Node[int], set map[int]struct{}) {
	if root == nil {
		return
	}

	inOrderSetMap(root.Left, set)
	set[root.Val] = struct{}{}
	inOrderSetMap(root.Right, set)
}

func inOrderCheckMap(root *Node[int], set map[int]struct{}) {
	if root == nil {
		return
	}

	inOrderCheckMap(root.Left, set)
	delete(set, root.Val)
	inOrderCheckMap(root.Right, set)
}

func ConstructBSTFromPreorder(p []int) *Node[int] {
	if len(p) == 0 {
		return nil
	}

	index := 0

	return constructBSTWithRange(p, &index, math.MinInt, math.MaxInt)
}

func constructBSTWithRange(pre []int, index *int, min, max int) *Node[int] {
	if *index >= len(pre) {
		return nil
	}

	val := pre[*index]
	if val <= min || val >= max {
		return nil
	}

	root := newNode(val)
	*index++

	root.Left = constructBSTWithRange(pre, index, min, val)
	root.Right = constructBSTWithRange(pre, index, val, max)

	return root
}

type List struct {
	Val  int
	Next *List
}

func SortedLinkedListToBalancedBST(list *List) *Node[int] {
	n := countListNodes(list)

	root, _ := sortedListToBST(list, n)

	return root
}

func sortedListToBST(list *List, n int) (*Node[int], *List) {
	if n <= 0 {
		return nil, list
	}

	left, list := sortedListToBST(list, n/2)

	root := newNode(list.Val)
	list = list.Next

	root.Left = left

	n = n - (n / 2) - 1

	root.Right, list = sortedListToBST(list, n)

	return root, list
}

func countListNodes(l *List) int {
	var count int

	cur := l
	for cur != nil {
		cur = cur.Next
		count++
	}

	return count
}
