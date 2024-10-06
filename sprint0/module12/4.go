package main

import (
	"errors"
)

func Factorial(n int) (int, error){
	if n > 0{
		b := 1
		for i := 1; i <= n; i++ {
			b *= i
		}
		return b, nil
	}else if n == 0{
		return 1, nil
	}
	return 0, errors.New("factorial is not defined for negative numbers")
}