package mergeTwoSortedArrays

import (
	"reflect"
	"testing"
)

func Test_mergeTwoSortedArrays(t *testing.T) {

	type args struct {
		firstArray  []int
		secondArray []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{firstArray: []int{1, 2, 3, 4, 6}, secondArray: []int{0}},
			want: []int{0, 1, 2, 3, 4, 6},
		},
		{
			name: "success",
			args: args{firstArray: []int{100, 200}, secondArray: []int{-1, -1, -1}},
			want: []int{-1, -1, -1, 100, 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoSortedArrays(tt.args.firstArray, tt.args.secondArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
