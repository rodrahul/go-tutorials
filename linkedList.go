package main

import (
	"fmt"
	"time"
)

type Node struct {
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	size int
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (list *List) empty() bool {
	if list.head == nil {
		return true
	} else {
		return false
	}
}

func (list *List) Size() int {
	return list.size
}

func (list *List) returnLastNode() *Node {
	tempNode := list.head
	for tempNode.next != nil {
		tempNode = tempNode.next
	}
	return tempNode
}

/*
Returns the value at an index
*/
func (list *List) ValueAt(index int) (error, interface{}) {
	if index > list.size-1 {
		fmt.Println("Return error error")
		return &MyError{
			time.Now(),
			"Index out of bounds",
		}, nil
	} else {
		// fmt.Println("index and size is", index, list.size)
		tempNode := list.head
		for i := 0; i < index; i++ {
			tempNode = tempNode.next
		}
		return nil, tempNode.key
	}
}

/*
ValueFrom tail
*/
func (list *List) ValueFromTail(index int) interface{} {
	if index > list.size {
		fmt.Println("ValueFromTail: Index out of bounds")
		return nil
	} else {
		index = list.size - index - 1
		_, returnValue := list.ValueAt(index)
		return returnValue
	}
}

/*
To add a new head node
*/
func (list *List) PushFront(key interface{}) {
	newNode := &Node{nil, key}
	if list.empty() {
		list.head = newNode
		list.size++
	} else {
		newNode.next = list.head
		list.head = newNode
		list.size++
	}

}

/*
Remove head and return the value
*/
func (list *List) PopFront() (error, interface{}) {
	if list.empty() {
		return &MyError{
			time.Now(),
			"Cannot Pop, list is empty",
		}, nil
	} else {
		returnValue := list.head.key
		list.head = list.head.next
		list.size--
		return nil, returnValue
	}
}

/*
PopBack, to remove the last node, and return the value
*/
func (list *List) PopBack() interface{} {
	if list.empty() {
		fmt.Println("List is empty cannot pop back")
		return nil
	} else {
		var previousNode *Node
		tempNode := list.head
		for tempNode.next != nil {
			previousNode = tempNode
			tempNode = tempNode.next
		}
		returnValue := tempNode.key
		previousNode.next = nil
		list.size--
		return returnValue
	}
}

/*
Append to add at the end of the list
*/
func (list *List) Append(key interface{}) {
	newNode := &Node{nil, key}
	if list.empty() {
		list.head = newNode
		list.size++
	} else {
		lastNode := list.returnLastNode()
		lastNode.next = newNode
		list.size++
	}
}

/*
InsertAfter is to insert of a specified node
*/

func (list *List) InsertAfter(index int, key interface{}) {
	newNode := &Node{nil, key}
	if list.empty() {
		list.head = newNode
		list.size++
	} else {
		tempNode := list.head
		for i := 0; i < index; i++ {
			tempNode = tempNode.next
		}
		newNode.next = tempNode.next
		tempNode.next = newNode
		list.size++
	}
}

/*
Delete the node at specified index
*/
func (list *List) DeleteIndex(index int) {
	if list.empty() {
		fmt.Println("List is empty cannot del")
	}
}

/*
Reverses the list
*/
func (list *List) ReverseList() {
	previousNode := &Node{nil, nil}
	currentNode := list.head
	nextNode := list.head.next
	for i := 0; i < list.size-1; i++ {

		fmt.Println("P:", previousNode.key, "C:", currentNode.key, "N:", nextNode.key)
		fmt.Println("I is: ", i)
		previousNode = currentNode
		if i == 0 {
			previousNode.next = nil
		} else if i == list.size-2 {
			list.head = nextNode
		}
		currentNode = nextNode
		nextNode = currentNode.next
		currentNode.next = previousNode
	}
}

func (list *List) Display() {
	tempNode := list.head
	for tempNode != nil {
		fmt.Printf("%+v -> ", tempNode.key)
		tempNode = tempNode.next
	}
	fmt.Println()
	fmt.Println("Linked List size is:", list.size)
}

func main() {
	ll := List{}
	fmt.Println("Is list empty?", ll.empty())
	ll.Display()

	ll.Append(5)
	fmt.Println("Is list empty?", ll.empty())
	ll.Display()

	ll.Append(6)
	ll.Display()
	ll.Append(7)
	ll.Append(8)
	ll.Append(9)
	ll.Append("Rahul")
	ll.Display()

	ll.InsertAfter(2, "new2")
	ll.Display()

	if err, value := ll.ValueAt(6); err == nil {
		fmt.Println("Value at index 6 is", value)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Push at front 4")
	ll.PushFront(4)

	ll.Display()
	fmt.Println("Pushing new head")
	ll.PushFront("New Head")
	ll.Display()

	if err, value := ll.PopFront(); err == nil {
		fmt.Println("popping head", value)
	} else {
		fmt.Println(err)
	}
	ll.Display()

	fmt.Println("Popping tail", ll.PopBack())
	ll.Display()

	fmt.Println("Value from tail ", ll.ValueFromTail(3))

	ll.ReverseList()
	ll.Display()

	// fmt.Println("List size is ", ll.Size())
}
