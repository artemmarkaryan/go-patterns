package factory_method

type House interface {
	AddHabitats()
	RemoveHabitats()
}

type HouseCreator interface {
	Create() House
}

type ModernHouse struct{}

func (m ModernHouse) AddHabitats() {}

func (m ModernHouse) RemoveHabitats() {}

type ModernHouseCreator struct{}

func (m ModernHouseCreator) Create() House {
	return ModernHouse{}
}
