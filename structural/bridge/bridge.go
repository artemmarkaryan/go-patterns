package bridge

import "log"

// Абстракция
type computer interface {
	print()
	setPrinter(printer)
}

// Расширенная абстракция
type mac struct {
	data string
	p    printer // мост между абстрацией и реализацией
}

func (m mac) print() {
	m.p.print(m.data)
}

func (m mac) setPrinter(p printer) {
	m.p = p
}

// Реализация
type printer interface {
	print(s string)
}

// Конкретная реализация
type hp struct{}

func (h hp) print(s string) {
	log.Print("Printed by HP: " + s)
}

func main() {
	m := mac{
		data: "i am mac computer",
	}
	m.setPrinter(hp{})
	m.print()
}
