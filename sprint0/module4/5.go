package main

import "fmt"

func main() {
	var num, count int
	fmt.Scanln(&num)
	if num >= 0{
		for i:=1; i <= num; i++{
			if i % 2 != 0{
				count+= i
			}
		}
		fmt.Println(count)
	}else{
		fmt.Println("Некорректный ввод")
	}
}