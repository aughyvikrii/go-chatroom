package handler

import (
	"chatroom/db"
	"chatroom/msg"
	"chatroom/types"
	"chatroom/util"
)

func RoomCreate(app *types.App, input string) (done bool) {
	if input == "" && app.Room.Title == "" {
		app.Conn.Write(msg.CmdInputRoomName)
	} else if input != "" && app.Room.Title == "" {
		app.Room.Title = input
		return RoomCreate(app, "")
	} else if input == "" && app.Room.Password == "" {
		app.Conn.Write(msg.CmdInputRoomPass)
	} else if input != "" && app.Room.Password == "" {
		app.Room.Password = input
		return RoomCreate(app, "")
	} else {
	code_loop:
		for {
			code := db.RoomCode(util.RandStringBytes(6))
			if _, exists := db.Rooms[code]; !exists {
				app.Room.Code = code
				break code_loop
			}
		}

		app.Room.Owner = app.User
		app.Room.Participant = make(map[db.Username]*db.User)
		app.Room.Participant[app.User.Username] = app.User
		app.Room.Conversation = make([]string, 0)

		db.Rooms[app.Room.Code] = app.Room

		info := "======== ROOM CREATED ========\n"
		info += "code: " + string(app.Room.Code) + "\n"
		info += "title: " + app.Room.Title + "\n"
		info += "==============================\n"

		app.Conn.Write([]byte(info))
		app.Conn.Write(msg.ConnectedToRoom)

		return true
	}

	return
}
