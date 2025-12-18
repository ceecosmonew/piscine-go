package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListClear(l *List) {
	if l == nil {
		return
	}
	l.Head = nil
	l.Tail = nil
}
