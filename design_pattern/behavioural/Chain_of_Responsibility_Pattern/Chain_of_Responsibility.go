package main

import "fmt"

// Handler interface defines a method to handle a request.
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler struct is the base for all concrete handlers
type BaseHandler struct {
	next Handler
}

// SetNext sets the next handler in the chain
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// Handle checks if there's a next handler in the chain
func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// ConcreteHandler1 handles "Request1"
type ConcreteHandler1 struct {
	BaseHandler
}

func (h *ConcreteHandler1) Handle(request string) {
	if request == "Request1" {
		fmt.Println("ConcreteHandler1 handled the request.")
	} else {
		fmt.Println("ConcreteHandler1 passed the request.")
		h.BaseHandler.Handle(request) // pass to the next handler
	}
}

// ConcreteHandler2 handles "Request2"
type ConcreteHandler2 struct {
	BaseHandler
}

func (h *ConcreteHandler2) Handle(request string) {
	if request == "Request2" {
		fmt.Println("ConcreteHandler2 handled the request.")
	} else {
		fmt.Println("ConcreteHandler2 passed the request.")
		h.BaseHandler.Handle(request)
	}
}

func main() {
	// Create handlers
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}

	// Set up the chain: handler1 -> handler2
	handler1.SetNext(handler2)

	// Pass requests through the chain
	fmt.Println("Sending Request1:")
	handler1.Handle("Request1")

	fmt.Println("\nSending Request2:")
	handler1.Handle("Request2")

	fmt.Println("\nSending Request3:")
	handler1.Handle("Request3") // no handler for this, passes through the chain
}
