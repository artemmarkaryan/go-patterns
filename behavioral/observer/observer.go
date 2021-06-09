package main

import "fmt"

type observer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Observer %q recieved email for item %q\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
