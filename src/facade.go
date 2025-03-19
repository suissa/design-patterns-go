package main

type CPU struct{}
func (c *CPU) Execute() { println("Executing") }

type Memory struct{}
func (m *Memory) Load() { println("Loading memory") }

type ComputerFacade struct {
    cpu    *CPU
    memory *Memory
}

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        cpu:    &CPU{},
        memory: &Memory{},
    }
}

func (c *ComputerFacade) Start() {
    c.memory.Load()
    c.cpu.Execute()
}