package adapter

import "fmt"

type NewPrintler interface {
	NewPrint() string
}

type OldPrintler interface {
	Print() string
}

type Adapter struct {
}

func (a *Adapter) Print() string {
	fmt.Println("old Print")
	return ""
}

type PrintAdapter struct {
	old *Adapter
}

func (a *PrintAdapter) NewPrint() string {
	oldResult := a.old.Print()
	return "new Print" + oldResult
}

func NewPrintlerAdapter(old *Adapter) NewPrintler {
	return &PrintAdapter{old: old}
}
