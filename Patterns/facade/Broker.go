package main

type Broker struct {
}

func (b *Broker) SendToChan(id string, data []byte) bool {
	return true
}

func (b *Broker) GetFromChan(id string) []byte {
	return []byte{}
}
