package evenNoOfDigits

import (
	"reflect"
	"testing"
)

func Test_evenNoOfDigits(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "success",
			input: []int{42, 564, 5775, 34, 123, 454, 1, 5, 45, 3556, 23442},
			want:  []int{42, 5775, 34, 45, 3556},
		},
		{
			name:  "success",
			input: []int{1, 0},
			want:  nil,
		},
		{
			name:  "success",
			input: []int{0, -20},
			want:  []int{-20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenNoOfDigits(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("identicalTwins() got = %v, want %v\", got, tt.want)", got, tt.want)
			}
		})
	}
}
