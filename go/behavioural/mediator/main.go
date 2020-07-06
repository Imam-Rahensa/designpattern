package main

func main() {
	room := newChatroom("japri_jobud")

	joko := chatuser{name: "joko", room: room}
	budi := chatuser{name: "budi", room: room}

	room.addUser(joko)
	room.addUser(budi)

	joko.send("halo budi")
	budi.send("hi joko")
}
