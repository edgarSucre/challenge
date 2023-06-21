package linklist

type ListNode[T any] struct {
	Next *ListNode[T]
	Val  T
}
