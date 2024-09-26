package main

import "fmt"

func PrettyArrayOutput(array [9]string) string{
	for i:=0; i < len(array); i++{
		if i < 7{
			fmt.Println(i+1, "я уже сделал:", array[i])
		}else{
			fmt.Println(i+1, "не успел сделать:", array[i])
		}
	}
	return ""
}