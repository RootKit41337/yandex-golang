package main 


func CalculateSeriesSum(n int) float64{
	var sum float64 
	for i:= 1; i <= n; i++{
		sum = sum +  (1.0 / float64(i))
	}
	return sum 
}