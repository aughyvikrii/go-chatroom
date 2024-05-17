package db

var Rooms map[RoomCode]*Room
var Users map[Username]*User

func init() {
	Rooms = make(map[RoomCode]*Room)
	Users = make(map[Username]*User)
}
