package LinkedListToArray

type Item int
type Node struct {
	data Item
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) insert(value Item) {
	newNode := Node{
		data: value,
		next: nil,
	}
	ll.size++
	if ll.head == nil {
		ll.head = &newNode
	} else {
		newNode.next = ll.head
		ll.head = &newNode
	}
}
func (ll *LinkedList) Append(value Item) {
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

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) Head() *Node {
	return ll.head
}
func (ll *LinkedList) Size() int {
	return ll.size
}
func (ll *LinkedList) linkedListToArray() []Item {
	current := ll.head
	var arr []Item
	for current.next != nil {
		arr = append(arr, current.data)
		current = current.next
	}
	arr = append(arr, current.data)
	return arr
}
