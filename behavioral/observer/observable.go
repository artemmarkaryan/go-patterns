package main

import (
	"fmt"
)

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	observerListLength := len(i.observerList)
	for index, v := range i.observerList {
		if o.getID() == v.getID() {
			i.observerList[observerListLength-1], i.observerList[index] = i.observerList[index], i.observerList[observerListLength-1]
			i.observerList = i.observerList[:observerListLength-1]
		}
	}
}

func (i *item) notifyAll() {
	for _, obs := range i.observerList {
		fmt.Printf("Sending update to %q\n", obs.getID())
		obs.update(i.name)
	}
}
