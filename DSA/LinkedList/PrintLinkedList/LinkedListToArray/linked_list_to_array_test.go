package LinkedListToArray

import (
	"reflect"
	"sync"
	"testing"
)

func TestLinkedList_linkedListToArray(t *testing.T) {
	type fields struct {
		head *Node
		size int
		lock sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []Item
	}{
		{
			name: "Success",
			fields: fields{head: &Node{
				data: 1,
				next: &Node{
					data: 2,
					next: nil,
				},
			}, size: 2},
			want: []Item{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := ll.linkedListToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("linkedListToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
