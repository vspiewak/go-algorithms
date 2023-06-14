package search

func BinaryRec(values []int, key int) (int, error) {
	return binary_recursive(values, key, 0, len(values)-1)
}

func binary_recursive(values []int, key int, start int, end int) (int, error) {

	if start > end {
		return -1, ErrNotFound
	}

	//mid := (start + end) / 2
	mid := start + ((end - start) / 2)

	if values[mid] == key {
		return mid, nil
	} else if key < values[mid] {
		return binary_recursive(values, key, start, mid-1)
	} else {
		return binary_recursive(values, key, mid+1, end)
	}

}

func BinaryIter(values []int, key int) (int, error) {

	start := 0
	end := len(values) - 1

	for start <= end {
		mid := start + ((end - start) / 2)
		if values[mid] == key {
			return mid, nil
		} else if key < values[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1, ErrNotFound

}

/*
func BinaryGen[K comparable](values []K, key K) (int, error) {
	start := 0
	end := len(values) - 1

	for start <= end {
		mid := start + ((end - start) / 2)
		if values[mid] == key {
			return mid, nil
		} else if key < values[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1, ErrNotFound
}
*/
