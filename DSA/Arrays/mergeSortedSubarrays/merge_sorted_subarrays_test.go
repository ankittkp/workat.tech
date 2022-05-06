package mergeSortedSubarrays

import (
	"reflect"
	"testing"
)

func Test_mergeSortedSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		endIndex int
		want     []int
	}{
		{
			name:     "success",
			input:    []int{1, 3, 5, 7, 9, 11, 0, 4, 6, 8},
			endIndex: 5,
			want:     []int{0, 1, 3, 4, 5, 6, 7, 8, 9, 11},
		},
		{
			name:     "success",
			input:    []int{3, 3, 9, 11, 1, 3, 3, 4},
			endIndex: 3,
			want:     []int{1, 3, 3, 3, 3, 4, 9, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortedSubarrays(tt.input, tt.endIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortedSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
