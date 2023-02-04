package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	a := l.Head
	b := l.Head
	c := l.Head
	c = nil
	for b != nil {
		next := b.Next
		b.Next = c
		c = b
		b = next
	}
	l.Head = c
	l.Tail = a
	l.Tail.Next = nil
}
