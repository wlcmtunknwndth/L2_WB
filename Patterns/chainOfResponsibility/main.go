package main

import "fmt"

type Data struct {
	data     any
	Received bool
	Cached   bool
	Saved    bool
}

type Saver interface {
	receive(*Data)
	setNext(Saver)
}

type Receiver struct {
	next Saver
}

func (r *Receiver) receive(data *Data) {
	if data.Received {
		fmt.Println("Data already received")
		r.next.receive(data)
	}
	fmt.Println("Received data from outside")
	data.Received = true
	r.next.receive(data)
}

func (r *Receiver) setNext(next Saver) {
	r.next = next
}

type Cacher struct {
	next Saver
}

func (c *Cacher) receive(data *Data) {
	if data.Cached {
		fmt.Println("Data already cached")
		c.next.receive(data)
	}
	fmt.Println("Cached data")
	data.Cached = true
	c.next.receive(data)
}

func (c *Cacher) setNext(next Saver) {
	c.next = next
}

type Storage struct {
	next Saver
}

func (s *Storage) receive(data *Data) {
	if data.Saved {
		fmt.Println("Data already saved in storage")
		s.next.receive(data)
	}
	fmt.Println("Saved data to storage")
	data.Saved = true
}

func (s *Storage) setNext(next Saver) {
	s.next = next
}

func main() {
	storage := &Storage{}

	cacher := &Cacher{}
	cacher.setNext(storage)

	receiver := &Receiver{}
	receiver.setNext(cacher)

	data := &Data{}
	receiver.receive(data)
}
