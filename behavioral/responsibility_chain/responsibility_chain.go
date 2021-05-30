package responsibility_chain

import "log"

type Seed interface {
	getSize() int
}

type Rice struct {
	size int
}

func (r Rice) getSize() int {
	return r.size
}

type Handler interface {
	handle(s Seed)
	process(s Seed)
	registerNext(h Handler)
}

type baseRiceHandler struct {
	next Handler
}

func (h baseRiceHandler) handle(s Seed) {
	h.process(s)
	if h.next != nil {
		h.next.handle(s)
	}
}

func (h baseRiceHandler) process(s Seed) {}

func (h baseRiceHandler) registerNext(nextH Handler) {
	h.next = nextH
}

type smallRiceHandler struct {
	baseRiceHandler
}

func (h smallRiceHandler) process(s Seed) {
	if s.getSize() < 10 {
		log.Print("Small seed")
	}
}

type bigRiceHandler struct {
	baseRiceHandler
}

func (h bigRiceHandler) process(s Seed) {
	log.Print("Big seed")
}


