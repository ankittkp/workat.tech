package insertionSort

import (
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "success",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "success",
			input: []int{3, 11, 4, 200},
			want:  []int{3, 4, 11, 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSort(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
