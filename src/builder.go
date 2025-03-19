package main

type House struct {
    Walls   string
    Roof    string
    Windows int
}

type HouseBuilder interface {
    BuildWalls() HouseBuilder
    BuildRoof() HouseBuilder
    BuildWindows() HouseBuilder
    GetHouse() House
}

type ModernHouseBuilder struct{ house House }

func (b *ModernHouseBuilder) BuildWalls() HouseBuilder {
    b.house.Walls = "Glass"
    return b
}

func (b *ModernHouseBuilder) BuildRoof() HouseBuilder {
    b.house.Roof = "Flat"
    return b
}

func (b *ModernHouseBuilder) BuildWindows() HouseBuilder {
    b.house.Windows = 4
    return b
}

func (b *ModernHouseBuilder) GetHouse() House { return b.house }