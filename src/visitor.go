package main

type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
}

type Shape interface {
	Accept(Visitor)
}

type Circle struct{}

func (c *Circle) Accept(v Visitor) { v.VisitCircle(c) }

type Square struct{}

func (s *Square) Accept(v Visitor) { v.VisitSquare(s) }

type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	println("Calculating circle area")
}
func (a *AreaCalculator) VisitSquare(s *Square) {
	println("Calculating square area")
}
