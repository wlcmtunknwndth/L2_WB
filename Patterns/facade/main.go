package main

import "fmt"

// Facade
type BackEnd struct {
	s *Storage
	c *Cacher
	h *Handler
	b *Broker
}

func New(s *Storage, c *Cacher, h *Handler, b *Broker) *BackEnd {
	return &BackEnd{s: s, c: c, h: h, b: b}
}

func (b *BackEnd) Start() {
	fmt.Println("Application started")
}

func (b *BackEnd) Stop() {
	fmt.Println("Application stopped")
}

// В данном примере я описал программу из первой задания на стажировке. Дело в том, что нам не требуются обилие методов для работы с приложением, ведь
// все происходит через протокол http, для чего нам нужно лишь запустить без ошибок кэширователь, хранилище, брокер и http-хэндлер и ждать
// запросов от пользователей. Минус только в том, что данный паттерн не устойчив к усложнению, ведь в случае изменения логики хотя бы одного
// из разделов, нам придется рефакторить всю структуру фасада.
func main() {
	s := Storage{}
	c := Cacher{}
	b := Broker{}
	h := Handler{&b}
	app := New(&s, &c, &h, &b)
	app.Start()

	app.Stop()
}
