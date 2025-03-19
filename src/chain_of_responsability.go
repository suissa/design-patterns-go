package main

type Handler interface {
	SetNext(Handler)
	Handle(int)
}

type BasicHandler struct {
	next Handler
}

func (b *BasicHandler) SetNext(h Handler) { b.next = h }

type NegativeHandler struct{ BasicHandler }

func (n *NegativeHandler) Handle(num int) {
	if num < 0 {
		println("Negative handled")
		return
	}
	n.BasicHandler.next.Handle(num)
}

type PositiveHandler struct{ BasicHandler }

func (p *PositiveHandler) Handle(num int) {
	if num >= 0 {
		println("Positive handled")
		return
	}
	p.BasicHandler.next.Handle(num)
}
