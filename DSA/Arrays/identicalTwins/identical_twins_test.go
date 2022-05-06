package identicalTwins

import (
	"reflect"
	"testing"
)

func Test_identicalTwins(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "success",
			input: []int{1, 2, 3, 2, 1},
			want:  2,
		},
		{
			name:  "success",
			input: []int{1, 2, 2, 3, 2, 1},
			want:  4,
		},
		{
			name:  "success",
			input: []int{1, 1, 1, 1},
			want:  6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := identicalTwins(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("identicalTwins() got = %v, want %v\", got, tt.want)", got, tt.want)
			}
		})
	}
}
