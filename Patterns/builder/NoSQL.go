package main

import "fmt"

type NoSQL struct {
	Name string
	Pass string
}

func NewNoSQL() *NoSQL {
	return &NoSQL{}
}

func (db *NoSQL) SaveOrder(order *Order) {
	fmt.Println("saved")
}

func (db *NoSQL) GetOrder(name string) *Order {
	return &Order{Name: "rnd", Price: 123}
}
