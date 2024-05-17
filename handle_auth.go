package main

import (
	"chatroom/handler"
	"chatroom/msg"
	"chatroom/types"
)

func auth(app *types.App) bool {

	app.Conn.Write([]byte("Pilih Opsi Berikut:\n1 - Masuk\n2 - Daftar\n"))
	app.Conn.Write(msg.CmdInput)

loop:
	for {

		input, err := app.Packet.Read(app.Conn)
		if err != nil {
			return false
		} else if input == "" {
			continue loop
		} else if input == "\\q" || input == "\\quit" {
			return false
		}

		if app.State.Position == "auth" {
			switch input {
			case "1":
				app.State.Position = "login"
				handler.AuthLogin(app, "")
			case "2":
				app.State.Position = "register"
				handler.AuthRegister(app, "")
			default:
				app.Conn.Write([]byte("Pilih Opsi Berikut:\n1 - Masuk\n2 - Daftar\n"))
				app.Conn.Write(msg.CmdInput)
				continue loop
			}
		} else if app.State.Position == "login" {
			if handler.AuthLogin(app, input) {
				app.State.Position = "main_menu"
				return true
			}
		} else if app.State.Position == "register" {
			if handler.AuthRegister(app, input) {
				app.State.Position = "main_menu"
				return true
			}
		}
	}
}
