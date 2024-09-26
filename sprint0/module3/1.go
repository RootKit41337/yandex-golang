package main

import "fmt"

func main() {
	var num int
	// fmt.Print("Введите число: ")  
	fmt.Scanln(&num)
	if num > 0{
		fmt.Println("Число положительное")
		return
	}
	if num < 0{
		fmt.Println("Число отрицательное")
        return
	}
	if num == 0 {
		fmt.Println("Введен ноль")
        return
	}
}