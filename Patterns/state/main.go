package main

import "fmt"

type Action interface {
	changeTemp()
}

type Fridge struct {
	temp         int8
	currentState Action
	tooCold      Action
	tooWarm      Action
}

func NewFridge(temp int8) *Fridge {
	fridge := &Fridge{
		temp: temp,
	}

	tooCold := TooCold{fridge: fridge}
	fridge.tooCold = &tooCold

	tooWarm := TooWarm{fridge: fridge}
	fridge.tooWarm = &tooWarm

	return fridge
}

func (f *Fridge) setState(state Action) {
	f.currentState = state
}

func (f *Fridge) Handle() {
	for i := 0; i < 20; i++ {
		fmt.Println(f.temp)
		f.contextAction()
		f.currentState.changeTemp()
	}
}

func (f *Fridge) contextAction() {
	if f.temp > -20 {
		f.currentState = f.tooCold
	} else {
		f.currentState = f.tooWarm
	}
}

type TooCold struct {
	fridge *Fridge
}

func (t *TooCold) changeTemp() {
	t.fridge.temp--
}

type TooWarm struct {
	fridge *Fridge
}

func (t *TooWarm) changeTemp() {
	t.fridge.temp++
}

func main() {
	fridge := NewFridge(-32)

	fridge.Handle()
}
