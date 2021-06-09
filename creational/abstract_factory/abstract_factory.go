package abstract_factory

type Roof interface {
	ResistRain()
}

type Walls interface {
	ResistWind()
}

type HouseFactory interface {
	MakeWalls() Walls
	MakeRoof() Roof
}

type ModernRoof struct{}

func (m ModernRoof) ResistRain() {}


type ModernWalls struct{}

func (m ModernWalls) ResistWind() {}


type ModernHouseFactory struct {}

func (m ModernHouseFactory) MakeWalls() Walls {
	return ModernWalls{}
}

func (m ModernHouseFactory) MakeRoof() Roof {
	return ModernRoof{}
}

func main() {
	// todo
}