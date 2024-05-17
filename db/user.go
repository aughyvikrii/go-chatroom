package db

import "net"

type Username string

type User struct {
	Username Username
	Password string
	Conn     net.Conn
}
