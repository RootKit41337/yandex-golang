package main

import "fmt"

func IsPowerOfTwoRecursive(n int){
	var i int
	i = 1
	for i < n {
		i = i * 2
	}
	if i == n{
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}