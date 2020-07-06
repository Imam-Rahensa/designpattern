package main

type chatroom struct {
	name  string
	users map[string]chatuser
}

func newChatroom(name string) *chatroom {
	return &chatroom{
		name:  name,
		users: make(map[string]chatuser),
	}
}

func (c *chatroom) sendMessage(msg string, from chatuser) {
	for _, u := range c.users {
		if u.name == from.name {
			continue
		}
		u.receive(msg)
	}
}

func (c *chatroom) addUser(user chatuser) {
	c.users[user.name] = user
}

func (c *chatroom) getRoomName() string {
	return c.name
}
