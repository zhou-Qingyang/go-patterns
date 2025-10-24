package bridge

import "fmt"

type Sender interface {
	Send(message string)
}

type EmailSender struct {
}

func (e *EmailSender) Send(message string) {
	fmt.Println("Email sent: ", message)
}

type SMSSender struct {
}

func (s *SMSSender) Send(message string) {
	fmt.Println("SMS sent: ", message)
}

type Message struct {
	sender Sender // 持有Sender接口，与实现维度桥接（核心）
}

func (m *Message) SetSender(sender Sender) {
	m.sender = sender
}

type MessageSender interface {
	SendMessage(message string)
	SetSender(sender Sender)
}

type NormalMessage struct {
	Message
}

func (m *NormalMessage) SendMessage(message string) {
	fmt.Println("Normal message sent: ", message)
	m.sender.Send(message)
}

func (m *NormalMessage) SetSender(sender Sender) {
	m.Message.SetSender(sender)
}

type SystemMessage struct {
	Message
}

func (m *SystemMessage) SendMessage(message string) {
	fmt.Println("System message sent: ", message)
	m.sender.Send(message)
}

func (m *SystemMessage) SetSender(sender Sender) {
	m.Message.SetSender(sender)
}
