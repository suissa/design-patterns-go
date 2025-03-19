package main

type Observer interface {
	Update(string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) NotifyAll(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type EmailAlert struct{}

func (e *EmailAlert) Update(msg string) {
	println("Email received:", msg)
}
