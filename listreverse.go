package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	if l == nil {
		return
	}

	// Empty list
	if l.Head == nil {
		l.Tail = nil
		return
	}

	var prev *NodeL
	current := l.Head

	// Old head becomes new tail
	l.Tail = l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// New head
	l.Head = prev
}
