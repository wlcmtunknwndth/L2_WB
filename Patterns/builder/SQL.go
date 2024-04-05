package main

import "fmt"

type SQL struct {
	Name string
	Pass string
}

func NewSQL() *SQL {
	return &SQL{}
}

func (db *SQL) SaveOrder(order *Order) {
	fmt.Println("saved, but SQL")
}

func (db *SQL) GetOrder(name string) *Order {
	return &Order{Name: "idk, but SQL", Price: 321}
}
