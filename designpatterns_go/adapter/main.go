package main

import "fmt"

type OldPrinter interface {
	PrintOldMessage() string
}

type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintOldMessage() string {
	return "Legacy Printer: Old message"
}

type NewPrinterAdapter struct {
	oldPrinter *LegacyPrinter
}

func (npa *NewPrinterAdapter) PrintMessage() string {
	return npa.oldPrinter.PrintOldMessage() + " - adapted"
}

func main() {
	adapter := NewPrinterAdapter{&LegacyPrinter{}}
	fmt.Println(adapter.PrintMessage())
}
