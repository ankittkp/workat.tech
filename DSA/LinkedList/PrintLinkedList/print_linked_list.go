package PrintLinkedList

import (
	"fmt"
	"sync"
)

type Item int
type Node struct {
	data Item
	next *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.Mutex
}

func (ll *LinkedList) Insert(value Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	newNode := Node{data: value, next: nil}
	ll.size++
	if ll.head == nil {
		ll.head = &newNode
	} else {
		newNode.next = ll.head
		ll.head = &newNode
	}
}
func (ll *LinkedList) Append(value Item) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	newNode := Node{data: value, next: nil}
	ll.size++
	if ll.head == nil {
		ll.head = &newNode
	} else {
		current := ll.head
		for nil != current.next {
			current = current.next
		}
		current.next = &newNode
	}
}
func (ll *LinkedList) InsertAt(value Item, position int) {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	newNode := Node{data: value, next: nil}
	ll.size++
	if ll.head == nil {
		ll.head = &newNode
	} else {
		current := ll.head
		for i := 1; i <= position && current.next != nil; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = &newNode
	}
}

//ToDO

func Delete() {

}

func (ll *LinkedList) IsEmpty() bool {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	return ll.head == nil
}

func (ll *LinkedList) Head() *Node {
	ll.lock.Lock()
	defer ll.lock.Unlock()
	return ll.head
}
func PrintLinkedList(values []Item) {
	ll := LinkedList{}
	for _, value := range values {
		ll.Append(value)
	}
	ll.Insert(2)
	ll.InsertAt(10, 2)

	current := ll.head
	for i := 1; i <= ll.size; i++ {
		fmt.Printf("Index: %v and Value: %v\n", i, current.data)
		current = current.next
	}

}
