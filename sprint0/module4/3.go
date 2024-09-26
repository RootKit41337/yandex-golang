package main

import "fmt"

func main() {
	var num, count int
	fmt.Scanln(&num)
	for i := 1; i<=num; i++{
		count += i
	}
	fmt.Println(count)
}