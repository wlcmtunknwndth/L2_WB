package main

type Storage struct {
}

type Order struct {
}

func (s *Storage) Get(name string) *Order {
	return &Order{}
}

func (s *Storage) Save(order *Order) bool {
	return true
}
