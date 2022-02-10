package containsElement

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsElement(t *testing.T) {
	initialArray := []int{1, 2, 3, 3, 3, 4, 4, 5}
	num := 2

	output := ContainsElement(initialArray, num)

	if !reflect.DeepEqual(output, true) {
		t.Errorf("Failed ! got %v want %v", output, true)
	} else {
		t.Logf("Success!!!")
	}
}
func TestMultipleContainsElement(t *testing.T) {
	tests := []struct {
		initialArray []int
		num          int
	}{
		{[]int{1, 2, 3, 3, 3, 4, 4, 5}, 2},
		{[]int{1, 2, 3, 3, 3, 4, 4, 5}, 6},
	}

	for index, value := range tests {
		t.Run(fmt.Sprintf("index=%d", index), func(t *testing.T) {
			output := ContainsElement(value.initialArray, value.num)
			if !reflect.DeepEqual(output, true) {
				t.Errorf("Failed ! got %v want %v", output, true)
			} else {
				t.Logf("Success!!!")
			}
		})
	}
}
