package builder

import (
	"log"
)

// House — нужный нам продукт
type House interface {
	AddHabitats()
	RemoveHabitats()
}
// HouseBuilderDirector — директор
type HouseBuilderDirector struct {
	builder HouseBuilder
}
// Build — метод директора, который отдаёт нам дом, того типа, который заложен в билдере
func (d HouseBuilderDirector) Build() House {
	return d.builder.Build()
}

// HouseBuilder — интерфейс билдера, который должен иметь метод Build
type HouseBuilder interface {
	Build() House
}

// ModernHouse — реализация интерфейса House
type ModernHouse struct{}

func (m ModernHouse) AddHabitats() {}

func (m ModernHouse) RemoveHabitats() {}

// ModernHouseBuilder — билдер для ModernHouse
type ModernHouseBuilder struct{}

func (m ModernHouseBuilder) Build() House {
	return ModernHouse{}
}

func main() {
	director := HouseBuilderDirector{builder: ModernHouseBuilder{}}
	house := director.Build()
	log.Print(house)
}