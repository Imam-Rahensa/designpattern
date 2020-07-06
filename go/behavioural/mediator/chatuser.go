package main

import "fmt"

type chatuser struct {
	name string
	room mediator
}

func (c *chatuser) send(msg string) {
	fmt.Println(c.name, "send message in", c.room.getRoomName(), msg)
	c.room.sendMessage(msg, *c)
}

func (c *chatuser) receive(msg string) {
	fmt.Println(c.name, "receiving message in", c.room.getRoomName(), msg)
}
