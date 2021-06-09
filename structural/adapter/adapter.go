package adapter

import "log"

type Circle struct {
	radius int
}

type Square struct {
	side int
}

type SquareToCircleAdapter struct{}

func (a SquareToCircleAdapter) Convert(s Square) Circle {
	return Circle{radius: s.side / 2}
}

func main() {
	adapter := SquareToCircleAdapter{}
	circle := adapter.Convert(Square{side: 10})
	log.Print(circle)
}