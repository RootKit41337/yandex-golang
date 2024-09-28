package main 

func SwapKeysAndValues(m map[string]string) map[string]string{
	m1 := make(map[string]string)
	for key, val := range m{
		m1[val] = key
	}
	return m1
}