package CountingSort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{1, 4, 1, 2, 7, 5, 2}
	res := []int{1, 1, 2, 2, 4, 5, 7}
	output := CountingSort(arr)
	if !reflect.DeepEqual(output, res) {
		t.Errorf("Failed ! got %v want %c", output, output)
	} else {
		t.Logf("Success !")
	}
}
func TestMultipleCumSum(t *testing.T) {
	testcases := []struct {
		initialArray []int
		output       []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2, 4, 3}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{-1, -2}, []int{-2, -1}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		{[]int{1, 9, 5, 1000, 100}, []int{1, 5, 9, 100, 1000}},
	}
	for index, value := range testcases {
		t.Run(fmt.Sprintf("index=%d", index), func(t *testing.T) {

			output := CountingSort(value.initialArray)
			if !reflect.DeepEqual(output, value.output) {
				t.Fatalf("got %v; want %v", output, value.output)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
