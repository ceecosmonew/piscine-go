package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	// Create a new node
	newNode := &NodeL{
		Data: data,
		Next: nil,
	}

	// If the list is empty
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	// If the list is not empty
	newNode.Next = l.Head
	l.Head = newNode
}
