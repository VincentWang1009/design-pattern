package bridge

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

func (m *Mac) Print() {
	m.printer.Print()
}

type Windows struct {
	printer Printer
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

func (w *Windows) Print() {
	w.printer.Print()
}

type Printer interface {
	Print()
}

type Epson struct{}

func (e *Epson) Print() {
	fmt.Println("epson printer print")
}

type Hp struct{}

func (h *Hp) Print() {
	fmt.Println("hp printer print")
}
