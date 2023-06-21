package linklist

func ListFromslice[T any](s []T) *ListNode[T] {
	var h *ListNode[T]
	var n *ListNode[T]

	for _, v := range s {
		if n != nil {
			n.Next = &ListNode[T]{Val: v}
			n = n.Next
			continue
		}

		h = &ListNode[T]{Val: v}
		n = h
	}

	return h
}
