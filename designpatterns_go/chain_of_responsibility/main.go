package main

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (bh *BaseHandler) SetNext(handler Handler) {
	bh.next = handler
}
func (bh *BaseHandler) HandleNext(request string) {
	fmt.Println("Performnig some operations in base ")
	if bh.next != nil {
		bh.next.Handle(request)
	}
}

// This is using embedding in go to satisfy interface
// Go looks for a HandleNext() method in ConcreteHandler.
// It doesn’t find it directly, but since BaseHandler is embedded,
// it automatically promotes the methods of BaseHandler to ConcreteHandler.
type ConcreteHandler struct {
	BaseHandler
}

func (c *ConcreteHandler) Handle(request string) {
	if request == "Concrete" {
		fmt.Println("this is concrete handler")
	} else {
		c.HandleNext(request)
	}
}

type TempHandler struct {
	BaseHandler
}

func (t *TempHandler) Handle(request string) {
	if request == "Temp" {
		fmt.Println("this is temp handler")
	} else {
		t.HandleNext(request)
	}
}

func main() {
	t := &TempHandler{}
	c := &ConcreteHandler{}

	t.SetNext(c)

	fmt.Println("new request")
	t.Handle("Temp")
	fmt.Println("new request")
	t.Handle("Concrete")
}
