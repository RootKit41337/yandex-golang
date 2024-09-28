package main 

func FindMaxKey(m map[int]int) int{
	maxkey := -300000
	for key, _ := range m{
		if key > maxkey{
			maxkey = key
		}
    }
	return maxkey
}