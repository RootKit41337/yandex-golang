package main 

func SumOfArray(array [6]int) int{
	sum := 0
	for i:=0; i < len(array); i++{
		sum += int(array[i])
	}
	return sum
}