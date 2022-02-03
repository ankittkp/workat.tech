package cumulative_sum

import (
	"fmt"
	"reflect"
	"testing"
)


func TestCumSum(t *testing.T) {
	initialArray := []int{1, 2, 3, 4}
	cumulativeSum := []int{1,3,6,10}
	var output [] int
	output = CumSum(initialArray)
	if !reflect.DeepEqual(output, cumulativeSum) {
		t.Errorf("Failed ! got %v want %c", output, cumulativeSum)
	} else {
		t.Logf("Success !")
	}

}
func TestMultipleCumSum(t *testing.T){
	testcases := []struct {
		initialArray []int
		cumulativeSum []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15}},
		{[]int{}, []int{}},
		{[]int{-1, -1}, []int{-1, -2}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 5, 7, 9}, []int{1, 4, 9, 16, 25}},
	}
	for index, value := range testcases {
		t.Run(fmt.Sprintf("index=%d", index), func(t *testing.T) {
			var output [] int
			output = CumSum(value.initialArray)
			if !reflect.DeepEqual(output, value.cumulativeSum) {
				t.Fatalf("got %v; want %v", output, value.cumulativeSum)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
