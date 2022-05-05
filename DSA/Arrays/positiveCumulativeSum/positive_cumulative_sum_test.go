package positiveCumulativeSum

import (
	"reflect"
	"testing"
)

func Test_positiveCumSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{arr: []int{1, -2, 3, 4, -6}},
			want: []int{1, 2, 6},
		},
		{
			name: "success",
			args: args{arr: []int{1, -1, -1, -1, 1}},
			want: []int{1},
		},
		{
			name: "success",
			args: args{arr: []int{1, 3, 5, 7}},
			want: []int{1, 4, 9, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := positiveCumSum(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("positiveCumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
