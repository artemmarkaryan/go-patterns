package main

import "fmt"

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}

