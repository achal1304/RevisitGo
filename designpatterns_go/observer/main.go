package main

import "fmt"

type Observer interface {
	Update(string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type EmailClient struct{}

func (e EmailClient) Update(data string) {
	fmt.Println("Email received:", data)
}

type PNSClient struct{}

func (e PNSClient) Update(data string) {
	fmt.Println("PNS received:", data)
}

func main() {
	subject := Subject{}
	emailClient := EmailClient{}
	pnsClient := PNSClient{}
	subject.Register(emailClient)
	subject.Register(pnsClient)
	subject.Notify("New Update Available!")
}
