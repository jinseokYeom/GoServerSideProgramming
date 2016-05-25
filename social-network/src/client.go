package client

import ()

// define a client
type Client struct {
	u    *User       // user info
	conn net.Conn    // connection
	ch   chan string // channel
}

// define a user
type User struct {
	Id       string
	Password []byte
	UserName string
	Email    string
}
