package flyweight

import "image/color"

// Большие данные
type catType struct {
	weight  float32
	race    color.Color
	texture []byte
}

// Маленькие данные
type cat struct {
	catType  catType
	position struct {
		x int
		y int
	}
}

type catFactory struct{}

func (c catFactory) GenerateCat(t catType, p struct {
	x int
	y int
}) cat {
	return cat{
		catType:  t,
		position: p,
	}
}
