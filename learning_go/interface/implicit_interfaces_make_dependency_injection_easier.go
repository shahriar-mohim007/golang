package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Logger interface for logging messages
type Logger interface {
	Log(message string)
}

// LoggerAdapter is an adapter to convert a function to Logger interface
type LoggerAdapter func(message string)

// Log implements the Logger interface
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// LogOutput is a simple logger function
func LogOutput(message string) {
	fmt.Println(message)
}
func LogOutput2(message string) {
	fmt.Println(">>>>>>>")
	fmt.Println(message)
}

// DataStore interface for retrieving user data
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// SimpleDataStore is a basic implementation of DataStore
type SimpleDataStore struct {
	userData map[string]string
}

// UserNameForID retrieves a username based on userID
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDataStore creates a new instance of SimpleDataStore
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// Logic interface defines the business logic methods
type Logic interface {
	SayHello(userID string) (string, error)
}

// SimpleLogic implements Logic interface
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// SayHello greets a user based on userID
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

// NewSimpleLogic creates a new instance of SimpleLogic
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// Controller handles HTTP requests
type Controller struct {
	l     Logger
	logic Logic
}

// SayHello handles the /hello endpoint
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// NewController creates a new instance of Controller
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput2)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
