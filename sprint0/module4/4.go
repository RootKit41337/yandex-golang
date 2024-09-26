package main

import "fmt"

func main() {
	var num int
	count:=1
	fmt.Scanln(&num)
	for i := 1; i<=num; i++{
		count *= i
	}
	fmt.Println(count)
}