package factory_method

import (
	"errors"
	"fmt"
)

type LoggerType string

const (
	DatabaseLoggerTYPE LoggerType = "DATABASE"
	FileLoggerTYPE     LoggerType = "FILE"
	ConsoleLogTYPE     LoggerType = "CONSOLE"
)

type Logger interface {
	Log(message string)
	Write(message string)
}

type DatabaseLogger struct {
}

func (d *DatabaseLogger) Log(message string) {
	fmt.Println("Database Logger: ", message)
}
func (d *DatabaseLogger) Write(message string) {
	fmt.Println("Database Logger Write:  ", message)
}

type FileLogger struct {
}

func (f *FileLogger) Log(message string) {
	fmt.Println("FileLogger :", message)
}
func (f *FileLogger) Write(message string) {
	fmt.Println("FileLogger write", message)
}

type RegisterLoggerFunc func() Logger

var LoggerFactory = make(map[LoggerType]Logger)

func RegisterLogger(loggerType LoggerType, initFunc RegisterLoggerFunc) {
	if initFunc == nil {
		panic("init func cannot be nil")
	}
	if _, ok := LoggerFactory[loggerType]; ok {
		return
	}
	LoggerFactory[loggerType] = initFunc()
	return
}

func NewLogger(loggerType LoggerType) (Logger, error) {
	loggerCreator, ok := LoggerFactory[loggerType]
	if !ok {
		return nil, errors.New("logger func not found")
	}
	return loggerCreator, nil
}

func init() {
	RegisterLogger(DatabaseLoggerTYPE, func() Logger {
		return &DatabaseLogger{}
	})
	RegisterLogger(FileLoggerTYPE, func() Logger {
		return &FileLogger{}
	})
}
