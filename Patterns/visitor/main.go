package main

import "fmt"

type Animal interface {
	getType() string
	accept(Visitor)
}

type Visitor interface {
	visitForDog(*Dog)
	visitForCat(*Cat)
}

type Cat struct {
	name string
}

func (c *Cat) accept(v Visitor) {
	v.visitForCat(c)
}

func (c *Cat) getType() string {
	return "Cat"
}

type Dog struct {
	name string
}

func (d *Dog) accept(v Visitor) {
	v.visitForDog(d)
}

func (d *Dog) getType() string {
	return "Dog"
}

type Voice struct {
	voice string
}

func (v *Voice) visitForDog(d *Dog) {
	fmt.Println(d.name, "said:", "bark")
}

func (v *Voice) visitForCat(cat *Cat) {
	fmt.Println(cat.name, "said:", "meow")
}

type Call struct {
	byName string
}

func (c *Call) visitForDog(d *Dog) {
	fmt.Println("Come here,", d.name)
}

func (c *Call) visitForCat(cat *Cat) {
	fmt.Println("Come here,", cat.name)
}

func main() {
	cat := &Cat{name: "Senya"}
	dog := &Dog{name: "Beta"}

	caller := &Call{}
	cat.accept(caller)
	dog.accept(caller)

	voicePls := &Voice{}
	cat.accept(voicePls)
	dog.accept(voicePls)
}
