package main

import (
	"math"
)


type  Circle struct{
	radius float64 
}

type Rectangle struct{
	width float64
	height float64
}

type Shape interface{
	Area()
}

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func (c Circle) Area() float64{
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() float64{
    return  roundFloat(r.height * r.width, 2)
}