package main

import "fmt"

// Mediator interface defines a method to send messages
type Mediator interface {
	SendMessage(message string, user *User)
}

// ChatRoom is the concrete mediator that facilitates communication between users
type ChatRoom struct {
	users []*User
}

// SendMessage sends a message from one user to all other users
func (c *ChatRoom) SendMessage(message string, user *User) {
	for _, u := range c.users {
		if u != user { // don't send the message back to the sender
			u.ReceiveMessage(message)
		}
	}
}

// AddUser adds a user to the chatroom
func (c *ChatRoom) AddUser(user *User) {
	c.users = append(c.users, user)
}

// User represents a colleague that communicates via the mediator
type User struct {
	name     string
	chatRoom Mediator
}

// SendMessage allows the user to send a message through the mediator
func (u *User) SendMessage(message string) {
	fmt.Printf("%s sends: %s\n", u.name, message)
	u.chatRoom.SendMessage(message, u)
}

// ReceiveMessage allows the user to receive a message
func (u *User) ReceiveMessage(message string) {
	fmt.Printf("%s receives: %s\n", u.name, message)
}

func main() {
	// Create the chatroom (mediator)
	chatRoom := &ChatRoom{}

	// Create users (colleagues)
	user1 := &User{name: "Alice", chatRoom: chatRoom}
	user2 := &User{name: "Bob", chatRoom: chatRoom}
	user3 := &User{name: "Charlie", chatRoom: chatRoom}

	// Add users to the chatroom
	chatRoom.AddUser(user1)
	chatRoom.AddUser(user2)
	chatRoom.AddUser(user3)

	// Users communicate via the chatroom
	user1.SendMessage("Hello everyone!")
	user2.SendMessage("Hey Alice!")
	user3.SendMessage("Hi all!")
}
