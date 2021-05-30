package adapter

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
