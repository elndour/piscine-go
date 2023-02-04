package piscine

func ListPushFront(l *List, data interface{}) {
	b := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = b
		l.Tail = b
	} else {
		b.Next = l.Head
		l.Head = b
	}
}
