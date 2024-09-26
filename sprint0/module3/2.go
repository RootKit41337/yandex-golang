package main

import "fmt"

func main() {
	var one, two int
	fmt.Scanln(&one, &two)
	if one > two{
		fmt.Println("Первое число больше второго")
		return
	}
	if one < two{
		fmt.Println("Второе число больше первого")
		return
	}
	if one == two{
		fmt.Println("Числа равны")
		return
	}
}