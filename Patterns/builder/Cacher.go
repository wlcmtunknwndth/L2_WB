package main

import "fmt"

type Cacher struct {
	cache map[string]string
}

func (c *Cacher) Cache(order *Order) {
	fmt.Println("cached")
}

func (c *Cacher) GetCached(name string) *Order {
	return &Order{Name: "cached", Price: 111}
}
