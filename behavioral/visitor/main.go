package main

func main() {
	s := &square{side: 2}
	c := &circle{radius: 3}

	calculator := &areaCalculator{}

	s.accept(calculator)
	c.accept(calculator)
}
