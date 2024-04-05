package main

type Order struct {
	Name  string
	Price uint
}

type Service interface {
	SaveOrder(storage *Order)
	GetOrder(name string) *Order
	//Cache(storage *Order)
	//GetCached(name string) *Order
}

type Storage struct {
	service Service
}

func NewStorage(srvc Service) *Storage {
	return &Storage{
		service: srvc,
	}
}

func main() {
	//cache := Cacher{make(map[string]string)}
	postgres := NewSQL()

	mongo := NewNoSQL()

	storage1 := NewStorage(postgres)
	storage1.service.SaveOrder(&Order{})
	storage1.service.GetOrder("idk")

	storage2 := NewStorage(mongo)
	storage2.service.SaveOrder(&Order{})
	storage2.service.GetOrder("idk")

}
