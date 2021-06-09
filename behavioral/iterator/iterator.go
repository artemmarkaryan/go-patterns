package main

import "log"

type iterator interface {
	getNext() *listEl
	hasMore() bool
}

type iterable interface {
	createIterator() iterator
}

type list struct {
	head *listEl
}

type listEl struct {
	data string
	next *listEl
}

func (l list) createIterator() iterator {
	return &listIterator{l, l.head}
}

func (l *list) push(d string) {
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

func (l *list) pop() {
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
	list  list
	state *listEl
}

func (l listIterator) hasMore() bool {
	return l.state != nil
}

func (l *listIterator) getNext() *listEl {
	if l.hasMore() {
		s := l.state
		l.state = l.state.next
		return s
	}
	return nil
}

func main() {
	l := list{}
	l.push("a")
	l.push("b")
	l.push("c")

	i := l.createIterator()

	for i.hasMore() {
		log.Print(i.getNext())
	}
}
