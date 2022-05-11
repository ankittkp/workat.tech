package IdenticalTwins

import "testing"

func Test_identicalTwins(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Test 1",
			arr:  []int{1, 2, 3, 2, 1},
			want: 2,
		},
		{
			name: "Test 2",
			arr:  []int{1, 2, 2, 3, 2, 1},
			want: 4,
		},
		{
			name: "Test 1",
			arr:  []int{1, 1, 1, 1},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := identicalTwins(tt.arr); got != tt.want {
				t.Errorf("identicalTwins() = %v, want %v", got, tt.want)
			}
		})
	}
}
