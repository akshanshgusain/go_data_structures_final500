package LinkedList

import "fmt"

type Node struct {
	key  interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

// Add node to the front
func (l *LinkedList) prepend(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++
}

// Delete node from the list
func (l *LinkedList) deleteWithValue(value int) {
	// If the list is empty
	if l.length == 0 {
		return
	}
	// If the node is the first Node
	if l.head.key == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.key != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

// Print list data
func (l LinkedList) printList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d, ", toPrint.key)
		toPrint = toPrint.next
		l.length--
	}
}

func TestRun() {
	list := LinkedList{}
	node1 := &Node{key: 45}
	node2 := &Node{key: 4}
	node3 := &Node{key: 5}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)

	list.printList()
}
