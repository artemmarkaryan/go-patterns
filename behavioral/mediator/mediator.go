package main

type mediator interface {
	// Условные действия медиатора
	canArrive(train) bool
	notifyAboutDeparture()
}

type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

// Конструктор
func newStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool {
	// Обработчик события "поезд просится на вокзал"
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	} else {
		s.trainQueue = append(s.trainQueue, t)
		return false
	}
}

func (s *stationManager) notifyAboutDeparture() {
	// Обработчик события "поезд уехал"
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
