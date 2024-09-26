package main

func FindMinMaxInArray(array [10]int) (int, int){
	minim := 30000
	maxim := -30000
	for i:=0; i < len(array); i++{
		if array[i] < minim {
            minim = array[i]
        }
		if array[i] > maxim{
			maxim = array[i]
        }
		}
		return maxim, minim
}