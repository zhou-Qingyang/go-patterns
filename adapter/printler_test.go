package adapter

import (
	"testing"
)

func TestPrintler(t *testing.T) {
	old := &Adapter{}
	p := NewPrintlerAdapter(old)
	t.Log(p.NewPrint())
}
