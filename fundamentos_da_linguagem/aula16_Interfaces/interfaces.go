package main

import (
	"fmt"
	"math"
)

type forma interface{
	area() float64
}

type retangulo struct{
	altura float64
	largura float64
}

type circulo struct{
	raio float64
}

func (r retangulo) area() float64  {
	return r.altura * r.largura
}

func escreverArea(f forma)  {
	fmt.Println("A area da forma é", f.area())
}

func (c circulo) area() float64  {
	return math.Pi * math.Pow(c.raio, 2)
}


func main()  {
	retangulo := retangulo{10,15}
	escreverArea(retangulo)

	circulo := circulo{10}
	escreverArea(circulo)

}