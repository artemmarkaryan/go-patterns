package composite

type osObj interface {
	getSize() int
}

type file struct {
	size int
}

func (f file) getSize() int {
	return f.size
}

type directory struct {
	inners []osObj
}

func (d directory) getSize() (r int) {
	for _, i := range d.inners {
		r += i.getSize()
	}
	return
}



