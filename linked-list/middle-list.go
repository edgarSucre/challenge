package linklist

func MiddleNode[T any](h *ListNode[T]) (middle *ListNode[T]) {
	middle = h
	end := h

	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}

	return middle
}
