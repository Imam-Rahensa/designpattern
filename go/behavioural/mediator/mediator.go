package main

type mediator interface {
	sendMessage(msg string, from chatuser)
	addUser(user chatuser)
	getRoomName() string
}
