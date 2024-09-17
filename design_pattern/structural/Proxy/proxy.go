package main

import (
	"fmt"
)

// Subject interface
type Document interface {
	Read() string
}

// RealSubject
type SensitiveDocument struct {
	content string
}

func (d *SensitiveDocument) Read() string {
	return d.content
}

// Proxy
type DocumentProxy struct {
	realDocument *SensitiveDocument
	userRole     string
}

func (p *DocumentProxy) Read() string {
	if p.userRole == "admin" {
		if p.realDocument == nil {
			// Load the document if it has not been loaded yet
			p.realDocument = &SensitiveDocument{content: "This is a sensitive document."}
		}
		return p.realDocument.Read()
	}
	return "Access Denied: You do not have permission to read this document."
}

// Client
func main() {
	// User with admin role
	adminProxy := &DocumentProxy{userRole: "admin"}
	fmt.Println(adminProxy.Read()) // Access granted, displays content

	// User without admin role
	guestProxy := &DocumentProxy{userRole: "guest"}
	fmt.Println(guestProxy.Read()) // Access denied
}
