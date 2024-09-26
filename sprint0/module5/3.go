package main


func SumDigitsRecursive(n int) int{
	if n == 0{
		return 0
	}else{
		return (n % 10 + SumDigitsRecursive(n / 10)) 
	} 
}