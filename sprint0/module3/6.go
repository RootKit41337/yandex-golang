package main

import "fmt"

func main() {
	var a int
	fmt.Scanln(&a)
	if a >= 0{
		if a < 10{
			fmt.Println("Число меньше 10")
			return
		}else if a >= 10 && a < 100{
			fmt.Println("Число меньше 100")
			return
		}else if a >= 100 && a < 1000{
			fmt.Println("Число меньше 1000")
        	return
		}else if a >= 1000{
			fmt.Println("Число больше или равно 1000")
            return
		}
	}else{
		fmt.Println("Некорректный ввод")
	}
}