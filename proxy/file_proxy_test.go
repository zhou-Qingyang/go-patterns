package proxy

import "testing"

func TestFileProxy(t *testing.T) {
	fileWriter := NewFileWriterProxy(1)

	err := fileWriter.Write([]byte("Hello, world!"))
	if err != nil {
		t.Error("Failed to write to file", err.Error())
	}

	fileWriter2 := NewFileWriterProxy(2)
	err = fileWriter2.Write([]byte("Hello, world2!"))
	if err != nil {
		t.Error("Failed to write to file", err.Error())
	}

}
