package main

import (
	"chatroom/db"
	"chatroom/msg"
	"chatroom/types"
	"net"
)

func handleConnection(conn net.Conn) {

	defer func() {
		conn.Write(msg.Goodbye)
		defer conn.Close()
	}()

	conn.Write(msg.Welcome)

	app := &types.App{
		Conn:   conn,
		Packet: make(types.Packet, 1024),
		State:  &types.State{Position: "auth"},
		User: &db.User{
			Conn: conn,
		},
		Room: &db.Room{},
	}

	if !auth(app) {
		return
	}

	mainmenu(app)
}
