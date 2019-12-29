package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		// 在 Go 中不能对切片使用等号运算符, 必须迭代每个元素
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("测试Sum", func(t *testing.T) {
		numbers := []int{1, 2, 3} // slice
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("测试smallSum", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{-10, 5, -6}, []int{})
		want := []int{3, 9, -11, 0}

		checkSum(t, got, want)
	})

	t.Run("测试SumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{0, 9, 12, 15})
		want := []int{7, 36}

		checkSum(t, got, want)
	})

	t.Run("空切片传入", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5}, []int{2})
		want := []int{0, 9, 0}

		checkSum(t, got, want)
	})
}
