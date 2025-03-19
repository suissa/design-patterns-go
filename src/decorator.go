
package main

type Pizza interface {
    GetPrice() int
}

type VeggiePizza struct{}
func (p *VeggiePizza) GetPrice() int { return 15 }

type CheeseTopping struct {
    pizza Pizza
}

func (c *CheeseTopping) GetPrice() int {
    return c.pizza.GetPrice() + 5
}