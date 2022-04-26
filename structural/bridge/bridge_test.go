package bridge

import "testing"

func TestBridge(t *testing.T) {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	winComputer := &windows{}

	// mac
	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	// windows
	winComputer.setPrinter(hpPrinter)
	winComputer.print()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
}
