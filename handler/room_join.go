package handler

import (
	"chatroom/db"
	"chatroom/msg"
	"chatroom/types"
)

func RoomJoin(app *types.App, input string) (done bool) {
	if input == "" && app.Room.Code == "" {
		app.Conn.Write(msg.CmdInputRoomCode)
	} else if input != "" && app.Room.Code == "" {
		code := db.RoomCode(input)
		if _, exists := db.Rooms[code]; !exists {
			app.Conn.Write(msg.InvalidRoom)
		} else {
			app.Room.Code = db.RoomCode(input)
		}

		return RoomJoin(app, "")
	} else if input == "" && app.Room.Password == "" {
		app.Conn.Write(msg.CmdInputRoomPass)
	} else if input != "" && app.Room.Password == "" {
		if roomDB, exists := db.Rooms[app.Room.Code]; !exists {
			app.Room.Code = ""
			app.Conn.Write(msg.InvalidRoom)
		} else if input != roomDB.Password {
			app.Conn.Write(msg.InvalidRoomPass)
		} else {
			app.Room.Password = input
		}

		return RoomJoin(app, "")
	} else {

		db.Rooms[app.Room.Code].Participant[app.User.Username] = app.User
		app.Room = db.Rooms[app.Room.Code]
		app.Conn.Write(msg.ConnectedToRoom)

		for _, other := range db.Rooms[app.Room.Code].Participant {
			if other.Username != app.User.Username {
				other.Conn.Write([]byte("@" + app.User.Username + " joined to chat room!\n"))
			}
		}

		return true
	}

	return
}
