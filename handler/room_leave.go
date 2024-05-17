package handler

import (
	"chatroom/db"
	"chatroom/msg"
	"chatroom/types"
)

func RoomLeave(app *types.App) {
	delete(db.Rooms[app.Room.Code].Participant, app.User.Username)

	if db.Rooms[app.Room.Code] != nil && len(db.Rooms[app.Room.Code].Participant) == 0 {
		delete(db.Rooms, app.Room.Code)
	}

	app.Conn.Write([]byte("keluar dari percakapan.\n"))
	app.State.Position = "main_menu"
	app.Conn.Write(msg.MainMenu)
	app.Conn.Write(msg.CmdInput)
	app.Room = &db.Room{}
}
