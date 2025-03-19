package main

type Command interface {
	Execute()
}

type Light struct{}

func (l *Light) On()  { println("Light on") }
func (l *Light) Off() { println("Light off") }

type LightOnCommand struct {
	light *Light
}

func (l *LightOnCommand) Execute() {
	l.light.On()
}
