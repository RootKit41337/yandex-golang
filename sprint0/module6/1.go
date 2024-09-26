package main

func FiveSteps(array [5]int) [5]int{
	var tmp int
	for i := 0; i < len(array)/2; i++ {
        tmp = array[i]
		array[i] = array[len(array)-i-1]
		array[len(array)-i-1] = tmp
    }
	return array
}