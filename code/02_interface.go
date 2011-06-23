package main

import (
	"fmt"
)

type Client struct {
	id      int
	name    string
}

type Manager struct {
	name    string
	level   int
}

type Printable interface {
	Print()
}

func (cl Client) Print() {
	fmt.Printf("ID: %d, Name: %s\n", cl.id, cl.name)
}

func (mg Manager) Print() {
	fmt.Printf("Level: %d, Name: %s\n", mg.level, mg.name)
}

func main() {
    var c *Client
    var m *Manager
    var p Printable

    c = new(Client); c.id = 5; c.name = "Kosh"
    m = new(Manager); m.name = "Krabutnica"; m.level = 10

    p = c; p.Print()
    p = m; p.Print()
}
