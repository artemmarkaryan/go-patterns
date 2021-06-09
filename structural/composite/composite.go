package main

import "log"

type osObj interface {
	// Композитный объект. Может или сам иметь размер,
	// или включать в себя другие объекты.
	getSize() int
}

type file struct {
	// Файл сам имеет размер
	size int
}

func (f file) getSize() int {
	return f.size
}

type directory struct {
	// Директория не имеет размер. 
	// Её размер = сумма размеров включенных директорий и файлов
	inners []osObj
}

func (d directory) getSize() (r int) {
	for _, i := range d.inners {
		r += i.getSize()
	}
	return
}

func main() {
	d := directory{
		inners: []osObj{
			directory{inners: []osObj{
				file{10},
				file{20},
			}},
			file{50},
		},
	}
	log.Print(d.getSize())
}
