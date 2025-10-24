package proxy

import (
	"errors"
	"fmt"
)

type IFileWriter interface {
	Write(message []byte) error
}

type FileWriter struct {
}

func (f *FileWriter) Write(message []byte) error {
	fmt.Println("File write:", string(message))
	return nil
}

type FileWriterProxy struct {
	IFileWriter
	UserId int
}

func (fp *FileWriterProxy) Write(message []byte) error {
	if fp.UserId == 1 {
		return fp.IFileWriter.Write(message)
	}
	return errors.New("User not authorized to write file")
}

func NewFileWriterProxy(userId int) *FileWriterProxy {
	return &FileWriterProxy{
		IFileWriter: &FileWriter{},
		UserId:      userId,
	}
}
