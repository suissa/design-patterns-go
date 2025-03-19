package main

type Renderer interface {
    RenderCircle(radius float32)
}

type VectorRenderer struct{}
func (v *VectorRenderer) RenderCircle(radius float32) {
    println("Drawing circle with radius", radius)
}

type Circle struct {
    renderer Renderer
    radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
    return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
    c.renderer.RenderCircle(c.radius)
}