package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil || l2 == nil || l2.Head == nil {
		return
	}

	// If l1 is empty, l2 becomes l1
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// Connect l1 tail to l2 head
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
