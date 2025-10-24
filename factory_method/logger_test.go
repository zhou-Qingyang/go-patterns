package factory_method

import "testing"

func TestLoggerCreator(t *testing.T) {
	dbLogger, _ := NewLogger(DatabaseLoggerTYPE)
	dbLogger.Log("data")
	dbLogger.Write("writing")

	fileLogger, _ := NewLogger(FileLoggerTYPE)
	fileLogger.Log("data")
	fileLogger.Write("writing")

	consoleLogger, err := NewLogger(ConsoleLogTYPE)
	if err != nil {
		t.Errorf("NewLogger error %s", err.Error())
		return
	}
	consoleLogger.Log("data")
	consoleLogger.Write("writing")
}
