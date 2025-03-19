package main

type Cloneable interface {
    Clone() Cloneable
}

type Car struct {
    Model string
}

func (c *Car) Clone() Cloneable {
    return &Car{Model: c.Model}
}