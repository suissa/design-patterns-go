package main

type LegacyPrinter interface {
    Print(s string) string
}

type ModernPrinter interface {
    PrintStored() string
}

type PrinterAdapter struct {
    LegacyPrinter LegacyPrinter
}

func (p *PrinterAdapter) PrintStored() string {
    return p.LegacyPrinter.Print("Adapter works")
}