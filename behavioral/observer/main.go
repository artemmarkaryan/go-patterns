package main

func main() {
	shirtItem := newItem("Nike Shirt")

	observer1 := &customer{id: "abc@gmail.com"}
	observer2 := &customer{id: "xyz@gmail.com"}

	shirtItem.register(observer1)
	shirtItem.register(observer2)

	shirtItem.updateAvailability()
}