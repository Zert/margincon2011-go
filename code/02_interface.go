package main

import (
	"fmt"
)

type Client struct {
	id      int
	name    string
}

type Manager struct {
	name     string
	clients  []int
}

type Printable interface {
	Print()
}

func (cl Client) Print() {
	fmt.Printf("ID: %d, Name: %s\n", cl.id, cl.name)
}

func (mg Manager) Print() {
	fmt.Printf("Name: %s, Clients: %d\n",
               mg.name, len(mg.clients))
}

func main() {
    var c *Client
    var m *Manager
    var p Printable

    c = new(Client); c.id = 5; c.name = "Kosh"
    m = new(Manager); m.name = "Krabutnica"; m.clients = []int{1,2,3}

    p = c; p.Print()
    p = m; p.Print()
}
