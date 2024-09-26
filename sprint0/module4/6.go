package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	if num >= 0{
		for i:=1; i <= num; i++{
			if i % 3 == 0{
				fmt.Println(i)
			}
		}
	}
}