package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	if num < 0 && num % 2 == 0{
		fmt.Println("Число отрицательное и четное")
		return
	}
	if num < 0 && num % 2 != 0{
		fmt.Println("Число отрицательное и нечетное")
		return
	}
	if num > 0 && num % 2 == 0{
		fmt.Println("Число положительное и четное")
		return
	}
	if num > 0 && num % 2 != 0{
		fmt.Println("Число положительное и нечетное")
		return
	}
}