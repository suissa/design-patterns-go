package main

type Button interface{ Render() }
type Label interface{ Display() }

type UIFactory interface {
    CreateButton() Button
    CreateLabel() Label
}

type WindowsFactory struct{}
func (w WindowsFactory) CreateButton() Button { return WinButton{} }
func (w WindowsFactory) CreateLabel() Label   { return WinLabel{} }

type MacFactory struct{}
func (m MacFactory) CreateButton() Button { return MacButton{} }
func (m MacFactory) CreateLabel() Label   { return MacLabel{} }

type WinButton struct{}
func (WinButton) Render() { println("Windows button") }

type MacButton struct{}
func (MacButton) Render() { println("Mac button") }

type WinLabel struct{}
func (WinLabel) Display() { println("Windows label") }

type MacLabel struct{}
func (MacLabel) Display() { println("Mac label") }