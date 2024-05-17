package msg

var Welcome = []byte("Welcome To ChatRoom!\n")
var Goodbye = []byte("Goodbye!:)\n")

var InvalidCmd = []byte("invalid option!\n")
var InvalidRoom = []byte("room not found!\n")
var InvalidRoomPass = []byte("room password invalid!\n")
var ConnectedToRoom = []byte("connected! welcome to chit chat~~\n")
var UserExists = []byte("username already registered!\n")
var UserPassNotMatch = []byte("invalid password!\n")
var RegisterSuccess = []byte("register success! welcome\n")
var LoginSuccess = []byte("login success! welcome\n")
var UserNotFound = []byte("user not found!\n")

var CmdInputUsername = []byte("#username> ")
var CmdInputPassword = []byte("#password> ")

var CmdInputRoomName = []byte("#room name> ")
var CmdInputRoomPass = []byte("#room password> ")
var CmdInputRoomCode = []byte("#room code> ")
var CmdInput = []byte("#> ")

var UnknownProcess = []byte("unknown process!type \"\\q\" to exit\n")

var MainMenu = []byte("Pilih menu:\n1 - Buat Room\n2 - Join Room\n")
