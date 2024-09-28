package main

func SumOfValuesInMap(m map[int]int) int{
	summa := 0
	for _, v := range m {
        summa += v
    }
	return summa 
}