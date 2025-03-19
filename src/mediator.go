package main

type Mediator interface {
	Notify(sender interface{}, event string)
}

type ChatRoom struct{}

func (c *ChatRoom) Notify(sender interface{}, event string) {
	println("Message received:", event)
}

type User struct {
	mediator Mediator
}

func (u *User) Send(message string) {
	u.mediator.Notify(u, message)
}
