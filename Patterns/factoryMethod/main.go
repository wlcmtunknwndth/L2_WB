package main

import "fmt"

//type Location interface {
//	setLocation(...any)
//	getLocation() []any
//}

type ChessLoc struct {
	xAxis byte
	yAxis byte
}

func (b *ChessLoc) setLocation(xAxis, yAxis byte) {
	b.xAxis = xAxis
	b.yAxis = yAxis
}

func (b *ChessLoc) getLocation() (byte, byte) {
	return b.xAxis, b.xAxis
}

type Figure interface {
	setLocation(byte, byte)
	getLocation() (byte, byte)
	setValue(uint8)
	getValue() uint8
}

type ChessFigure struct {
	val byte
	loc *ChessLoc
}

func (f *ChessFigure) setValue(val uint8) {
	f.val = val
}

func (f *ChessFigure) getValue() uint8 {
	return f.val
}

func (f *ChessFigure) setLocation(x, y byte) {
	f.loc = &ChessLoc{
		xAxis: x,
		yAxis: y,
	}
}

func (f *ChessFigure) getLocation() (byte, byte) {
	return f.loc.xAxis, f.loc.yAxis
}

type Rook struct {
	ChessFigure
}

func newRook(x, y byte) Figure {
	return &Rook{
		ChessFigure: ChessFigure{
			val: 1,
			loc: &ChessLoc{x, y},
		},
	}
}

type Bishop struct {
	ChessFigure
}

func newBishop(x, y byte) Figure {
	return &Bishop{
		ChessFigure: ChessFigure{
			val: 5,
			loc: &ChessLoc{x, y},
		},
	}
}

func getFigure(figName string) (Figure, error) {
	if figName == "rook" {
		return newRook(4, 4), nil
	}
	if figName == "bishop" {
		return newBishop(3, 3), nil
	}

	return nil, fmt.Errorf("no figure named %s exists", figName)
}

func main() {
	rook, err := getFigure("rook")
	if err != nil {
		fmt.Println("couldn't create fig: ", err)
	}
	fmt.Printf("%+v", rook)

	bishop, err := getFigure("bishop")
	if err != nil {
		fmt.Println("couldn't create fig: ", err)
	}
	fmt.Printf("%+v", bishop)
}
