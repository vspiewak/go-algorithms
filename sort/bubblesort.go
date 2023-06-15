package sort

import . "github.com/vspiewak/go-algorithms/constraints"

func Bubblesort[T Ordered](values []T) []T {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(values)-1; i++ {
			if values[i] > values[i+1] {
				values[i+1], values[i] = values[i], values[i+1]
				swapped = true
			}
		}
	}
	return values
}
