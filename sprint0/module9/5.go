package main 

func DeleteLongKeys(m map[string]int) map[string]int{
	m1 := make(map[string]int)
	for key, value := range m{
		if len(key) >= 6 {
			m1[key] = value
        }
	}
	return m1
}