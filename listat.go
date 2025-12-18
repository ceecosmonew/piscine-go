package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current
		}
		current = current.Next
		index++
	}

	return nil
}
