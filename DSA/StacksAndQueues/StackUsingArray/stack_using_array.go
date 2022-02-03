package StackUsingArray

import (
	"fmt"
	"sync"
)
const maxStackSize = 100
type Value interface{}	//can be of any date type
type Stack struct {
	array[] Value
	Mutex sync.Mutex //for consistency and race condition
}
func (s *Stack) HandleMutex () {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
}

func (s *Stack) Push(value Value) {
	s.HandleMutex()
	if s.IsFull() {
		fmt.Println("Stack is full")
	}
	s.array = append(s.array, value)
}
func (s *Stack) Pop() (Value, bool) {
	s.HandleMutex()
	if s.IsEmpty(){
		fmt.Println("Stack is Empty")
		return -1, false
	}

	topElement :=  s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return topElement, true
}
func (s *Stack) Top() Value{
	s.HandleMutex()
	if len(s.array) == 0 {
		return nil
	}
	return s.array[len(s.array)-1]
}
func (s *Stack) IsEmpty() bool {
	s.HandleMutex()
	if len(s.array) <= 0{
		return true
	} else {
		return false
	}
}
func (s *Stack) IsFull() bool{
	s.HandleMutex()
	if len(s.array) == maxStackSize{
		return true
	} else{
		return false
	}
}
func (s *Stack) PrintStack() {
	s.HandleMutex()
	for index, value := range s.array {
		fmt.Printf("index: %v and value: %v\n", index, value)
	}
}
