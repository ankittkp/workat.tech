package Remove_Occurences

import (
	"fmt"
	"testing"
)

func TestRemoveOccurences(t *testing.T) {
	array := [] int {1, 4, 2, 6, 2, 6, 9, 4}
	k := 4
	result := 6
	output := RemoveOccurences(array, k)
	if output!= result{
		t.Errorf("Failed ! got %v want %c", output, result)
	} else {
		t.Logf("Success !")
	}
}

func TestMultipleRemoveOccurences(t *testing.T) {
	tests := [] struct {
		array [] int
		k int
		result int
	}{
		{[]int {1, 4, 2, 6, 2, 6, 9, 4}, 4, 6},
		{[]int{1,3,3,3,4,4},3, 3},
	}
	for index, value := range tests{
		t.Run(fmt.Sprintf("index: %d", index), func(t *testing.T) {
			output := RemoveOccurences(value.array, value.k)
			if output != value.result{
				t.Errorf("Failed ! got %v want %c", output, value.result)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
