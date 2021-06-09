package main

func main() {
	// Посредник
	manager := newStationManger()

	// При создании компонента указываем его посредника
	pTrain := &passengerTrain{
		mediator: manager,
	}

	pTrain.arrive()
	pTrain.arrive()
	pTrain.arrive()
	pTrain.depart()
	pTrain.depart()
	pTrain.depart()
}
