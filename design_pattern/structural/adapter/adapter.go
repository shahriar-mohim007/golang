package main

import "fmt"

// Old Interface
type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintOldFormat(message string) {
	fmt.Println("Printing in old format:", message)
}

// New Interface
type Printer interface {
	Print(message string)
}

// Adapter
type PrinterAdapter struct {
	legacyPrinter *LegacyPrinter
}

func (pa *PrinterAdapter) Print(message string) {
	pa.legacyPrinter.PrintOldFormat(message)
}

func main() {
	legacyPrinter := &LegacyPrinter{}
	printer := &PrinterAdapter{legacyPrinter: legacyPrinter}

	printer.Print("Hello, World!")
}
