package sort

import (
	"reflect"
	"testing"
)

func TestBubblesort(t *testing.T) {

	t.Run("Test empty Bubblesort", func(t *testing.T) {

		initial := []int{}
		actual := Bubblesort[int](initial)
		expected := []int{}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Bubblesort", func(t *testing.T) {

		initial := []int{1, 3, 2, 5, 4}
		actual := Bubblesort[int](initial)
		expected := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
