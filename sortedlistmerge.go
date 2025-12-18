package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// If one list is empty, return the other
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	var head *NodeI

	// Choose initial head
	if n1.Data <= n2.Data {
		head = n1
		n1 = n1.Next
	} else {
		head = n2
		n2 = n2.Next
	}

	current := head

	// Merge while both lists have elements
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Attach remaining nodes
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	return head
}
