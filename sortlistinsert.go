package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	// Case 1: empty list OR insert before head
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	current := l

	// Find insertion point
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert node
	newNode.Next = current.Next
	current.Next = newNode

	return l
}
