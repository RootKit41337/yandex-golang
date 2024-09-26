package main

import "fmt"

func main() {
	var num, count int
	fmt.Scanln(&num)
	for i:=1; i <= num; i++{
		if (i % 3 != 0) && (i % 5 != 0){
			count+= i
		}
	}
	fmt.Println(count)
}