package main

import ("fmt"
		"math"
)

func SqRoots(){
	var x1, x2, a, b, c float64
	fmt.Scanln(&a, &b, &c)
	disckrim := math.Pow(b, 2) - 4.0 * a * c
	if disckrim > 0{
		x1 = (-b - math.Sqrt(disckrim)) / (2*a)
		x2 = (-b + math.Sqrt(disckrim)) / (2*a)
		fmt.Println(x1,x2)
	}
	if disckrim == 0{
        x1 = -b / (2*a) 
		fmt.Println(x1)
    }
	if disckrim < 0{
		fmt.Println(x1, x2)    
    }
}