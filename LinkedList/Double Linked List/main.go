package main

import (
	"fmt"
)

// Double LinkedList struct
type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
	Size int
}

// DNode - Node structure in the Doubly LinkedList data structure
type DNode struct {
	Next  *DNode
	Prev  *DNode
	Value string
}

func main() {
	// Test dLL, Insert, List
	var ll DoublyLinkedList
	ll.Insert("one")
	ll.Insert("two")
	ll.Insert("three")
	ll.List()

	// Test Delete, Insert, List
	ll.Delete("three")
	ll.List()

	// Test Delete First
	ll.Insert("three")
	ll.Insert("four")
	ll.Insert("five")
	ll.List()

	ll.DeleteFirst()
	ll.List()
}

func (l *DoublyLinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Printf("%+v\n", current)
		current = current.Next
	}
	fmt.Printf("\n++++++++++++++++++++++++++++++++++++++\n")
}

func (l *DoublyLinkedList) Insert(element string) {
	current := l.Head
	n := &DNode{Next: nil,
		Prev:  nil,
		Value: element,
	}
	if current == nil {
		l.Head = n
		l.Tail = n
	} else {
		for current.Next != nil {
			current = current.Next
		}
		n.Prev = current
		current.Next = n
		l.Tail = n
	}
	// } else {
	// 	tmp := l.Head
	// 	for tmp.Next != nil {
	// 		tmp = tmp.Next
	// 	}
	// 	n.Prev = tmp
	// 	tmp.Next = n
	// 	l.Tail = n
	// }

	l.Size++
}

func (l *DoublyLinkedList) Delete(element string) {
	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Value == element {
			previous.Next = current.Next
			fmt.Printf("the element you are looking for '%+v' is deleted from the current doubly linked list\n", element)
		}
		previous = current
		current = current.Next
		l.Size--
	}
}

func (l *DoublyLinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}
