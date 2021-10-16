package main

import (
	"fmt"
	"time"
)

type Courier struct {
	Name string
}
type Product struct {
	Name  string
	Price int
	ID    int
}
type Parcel struct {
	Pdt           *Product
	ShiooedTime   time.Time
	DeliveredTime time.Time
}

func (c Courier) SendProduct(pdt *Product) *Parcel {
	p := &Parcel{}
	p.Pdt = pdt
	p.ShiooedTime = time.Now()
	return p
}

func (p Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {
	a := Product{"감자", 2, 10}
	b := Parcel{
		&a, time.Now(), time.Now()}
	c := b.Pdt
	d := b.ShiooedTime
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
