package main 

func CountingSort(contacts []string) map[string]int{
	m := make(map[string]int)
	for _, value := range contacts{
		m[value]++
	}
	return m
}