package main

import (
	"chatroom/handler"
	"chatroom/msg"
	"chatroom/types"
)

func mainmenu(app *types.App) {

	app.Conn.Write(msg.MainMenu)
	app.Conn.Write(msg.CmdInput)

loop:
	for {

		input, err := app.Packet.Read(app.Conn)
		if err != nil {
			return
		} else if input == "" {
			continue loop
		} else if input == "\\q" || input == "\\quit" {
			handler.RoomLeave(app)
			return
		}

		if app.State.Position == "main_menu" {
			switch input {
			case "1":
				app.State.Position = "create_room"
				handler.RoomCreate(app, "")
			case "2":
				app.State.Position = "join_room"
				handler.RoomJoin(app, "")
			default:
				app.Conn.Write(msg.InvalidCmd)
				app.Conn.Write(msg.MainMenu)
				app.Conn.Write(msg.CmdInput)
				continue loop
			}
		} else if app.State.Position == "create_room" {
			if handler.RoomCreate(app, input) {
				app.State.Position = "chit_chat"
			}
		} else if app.State.Position == "join_room" {
			if handler.RoomJoin(app, input) {
				app.State.Position = "chit_chat"
			}
		} else if app.State.Position == "chit_chat" {
			participants := app.Room.Participant
			isLeave := input == "\\leave" || input == "\\l"

			if isLeave {
				handler.RoomLeave(app)
				input = "keluar dari percakapan."
			}

			for _, participant := range participants {
				if participant.Username != app.User.Username {
					go participant.Conn.Write([]byte("@" + string(app.User.Username) + "> " + input + "\n"))
				}
			}
		}

	}
}
