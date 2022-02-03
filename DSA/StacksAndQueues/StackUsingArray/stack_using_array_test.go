package StackUsingArray

import (
	"reflect"
	"testing"
)

func CreateStack() *Stack {
	return &Stack{
		array: nil,
	}
}
func NewStack(arr []Value) *Stack {
	return &Stack{
		array: arr,
	}
}
func TestPush(t *testing.T) {
	stack := CreateStack()
	t.Run("Pushing new Item", func(t *testing.T) {
		result := NewStack([]Value{
			1,
			2,
			3,
		})
		t.Run("matching result with output stack", func(t *testing.T) {
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
		})
		if !reflect.DeepEqual(stack, result) {
			t.Fatalf("got %v; want %v", stack, result)
		} else {
			t.Logf("Success !")
		}
	})
}

//TODO All test for other functions
func TestPop(t *testing.T) {

}
func TestTop(t *testing.T) {

}
func TestIsEmpty(t *testing.T) {

}
func TestIsFull(t *testing.T) {

}
func TestPrintStack(t *testing.T) {

}