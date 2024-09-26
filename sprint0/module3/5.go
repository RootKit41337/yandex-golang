package main

import "fmt"

func main() {
	var a, b ,c int
	fmt.Scanln(&a, &b, &c)
	if a != 0 && b != 0 && c != 0{
		if a == b && b == c && a == c{
			fmt.Println("Все числа равны")
			return
		} else if a == b || b == c || c == a{
			fmt.Println("Два числа равны")
			return
		}else if a != b && a != c && b != c{
			fmt.Println("Все числа разные")
			return
		}else{
			fmt.Println("Некорректный ввод")
		}
	}else{
		fmt.Println("Некорректный ввод")
	}	
}