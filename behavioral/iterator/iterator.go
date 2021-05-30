package iterator

type iterable interface {
	createIterator() iterator
}

type iterator interface {
	next()
	hasMore() bool
}

type listEl struct {
	data string
	next *listEl
}

type list struct {
	head *listEl
}

func (l list) createIterator() iterator {
	return listIterator{l, l.head}
}

func (l list) push(d string) {
	newEl := &listEl{d, nil}
	if l.head == nil {
		l.head = newEl
		return
	}
	el := l.head
	for el.next != nil {
		el = el.next
	}
	el.next = newEl
}

func (l list) pop() {
	if l.head == nil {
		return
	}
	el := l.head
	for el.next.next != nil {
		el = el.next
	}
	el.next = nil
}

type listIterator struct {
	list list
	state *listEl
}

func (l listIterator) hasMore() bool {
	return l.state.next != nil
}

func (l listIterator) next() {
	if l.hasMore() {
		l.state = l.state.next
	}
}

