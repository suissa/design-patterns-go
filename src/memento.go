package main

type Editor struct {
	content string
}

func (e *Editor) CreateMemento() *Memento {
	return &Memento{content: e.content}
}

func (e *Editor) Restore(m *Memento) {
	e.content = m.content
}

type Memento struct {
	content string
}
