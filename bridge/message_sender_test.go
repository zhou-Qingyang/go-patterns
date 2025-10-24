package bridge

import (
	"fmt"
	"testing"
)

func TestMessageSender(t *testing.T) {
	emailSender := &EmailSender{}
	smsSender := &SMSSender{}

	normalMessage := &NormalMessage{}
	normalMessage.SetSender(emailSender)
	normalMessage.SendMessage("sender")

	fmt.Println()

	normalMessage.SetSender(smsSender)
	normalMessage.SendMessage("sender")
}
