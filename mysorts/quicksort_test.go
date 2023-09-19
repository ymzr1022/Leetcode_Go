package mysorts

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{2, 3, 1, 1, 0}
	quickSort(arr, 0, 4)
	fmt.Printf("arr: %v\n", arr)
}
