package types

import (
	"chatroom/db"
	"net"
)

type App struct {
	Conn   net.Conn
	Packet Packet
	State  *State
	User   *db.User
	Room   *db.Room
}
