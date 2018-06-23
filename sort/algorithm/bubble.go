package algorithm

func BSort(values []int) (arr []int, err error) {
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return
}
