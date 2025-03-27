package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("sending email noti ", message)
}

type NotifierDecorator struct {
	Wrappee Notifier
}

func (d *NotifierDecorator) Send(message string) {
	d.Wrappee.Send(message)
}

type LoggingDecorator struct {
	Wrappee Notifier
}

func (l *LoggingDecorator) Send(message string) {
	fmt.Println("[Log] Message about to be sent:", message)
	l.Wrappee.Send(message)
}

type SlackDecorator struct {
	Wrappee Notifier
}

func (s *SlackDecorator) Send(message string) {
	s.Wrappee.Send(message)
	fmt.Println("Sending SLACK notification:", message)
}

func main() {
	email := &EmailNotifier{}

	logged := &LoggingDecorator{Wrappee: email}
	slackAndLogged := &SlackDecorator{Wrappee: logged}
	slackAndLogged.Send("Server down!")

}
