package facade

import (
	"image/color"
	"time"
)

type House struct {
	Materials string
	BuiltAt   time.Time
	Owner     string
	Credited  bool

	Color         color.Color
	WindowsAmount int
	HasRoof       bool
}

type HouseExteriorFacade struct {
	Color         color.Color
	WindowsAmount int
	HasRoof       bool
}
