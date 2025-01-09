package main

import "fmt"

// Scanner interface for scanning documents
type Scanner interface {
	Scan()
}

// Copier interface for copying documents
type Copier interface {
	Copy()
}

// MultiFunctionPrinter implements both Scanner and Copier interfaces
type MultiFunctionPrinter struct{}

func (mfp MultiFunctionPrinter) Scan() {
	fmt.Println("Scanning document...")
}

func (mfp MultiFunctionPrinter) Copy() {
	fmt.Println("Copying document...")
}

func main() {
	mfp := MultiFunctionPrinter{}
	mfp.Scan()
	mfp.Copy()
}
