package handler

import (
	"chatroom/db"
	"chatroom/msg"
	"chatroom/types"
)

func AuthLogin(app *types.App, input string) (done bool) {
	if input == "" && app.User.Username == "" { // if input empty and username empty, tell user to input username
		app.Conn.Write(msg.CmdInputUsername)
	} else if input != "" && app.User.Username == "" { // if input not empty and username empty, then add input as username
		username := db.Username(input)

		if _, ok := db.Users[username]; !ok { // if username already exists then stop it
			app.Conn.Write(msg.UserNotFound)
		} else {
			app.User.Username = username
		}

		return AuthLogin(app, "") // redirect to input password
	} else if input == "" && app.User.Password == "" { // if input empty and password empty and (username already exists) then tell user to input password
		app.Conn.Write(msg.CmdInputPassword)
	} else if input != "" && app.User.Password == "" {
		userInDB := db.Users[app.User.Username]

		if userInDB.Password != input {
			app.Conn.Write(msg.UserPassNotMatch)
			return AuthLogin(app, "")
		}

		app.Conn.Write(msg.LoginSuccess)
		return true

	} else {
		app.Conn.Write(msg.UnknownProcess)
	}

	return
}
