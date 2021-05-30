package prototype

type Cloneable interface {
	Clone() Cloneable
}

type House struct {
	address string
}

func (h House) Clone() Cloneable {
	return House{address: h.address}
}
