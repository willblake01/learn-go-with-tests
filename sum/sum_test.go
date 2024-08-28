package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("sum all numbers in multiple slices", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{0, 9}

		got := SumAll(nums1, nums2)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}
