package handler

import (
	"chatroom/db"
	"chatroom/msg"
	"chatroom/types"
)

func AuthRegister(app *types.App, input string) (done bool) {
	if input == "" && app.User.Username == "" { // if input empty and username empty, tell user to input username
		app.Conn.Write(msg.CmdInputUsername)
	} else if input != "" && app.User.Username == "" { // if input not empty and username empty, then add input as username
		username := db.Username(input)

		if _, ok := db.Users[username]; ok { // if username already exists then stop it
			app.Conn.Write(msg.UserExists)
		} else {
			app.User.Username = username
			db.Users[username] = app.User
		}

		return AuthRegister(app, "") // redirect to input password
	} else if input == "" && app.User.Password == "" { // if input empty and password empty and (username already exists) then tell user to input password
		app.Conn.Write(msg.CmdInputPassword)
	} else if input != "" && app.User.Password == "" {

		app.User.Password = input
		db.Users[app.User.Username] = app.User

		app.Conn.Write(msg.RegisterSuccess)
		return true

	} else {
		app.Conn.Write(msg.UnknownProcess)
	}

	return
}
