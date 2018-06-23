package algorithm

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	Sort(arr)
	i := 1
	for _, value := range arr {
		if value != i {
			t.Errorf("sort failed %d,%d,%d,%d,%d", arr[0], arr[1], arr[2], arr[3], arr[4])
		}
		i++
	}
}
