package db

type RoomCode string

type Room struct {
	Title        string
	Code         RoomCode
	Password     string
	Participant  map[Username]*User
	Owner        *User
	Conversation []string
}
