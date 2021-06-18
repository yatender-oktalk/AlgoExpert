package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func main() {
	l := LinkedList{}
	fmt.Println("\n************* Insert *************")
	l.Insert(12)
	l.Insert(14)
	l.Insert(24)
	l.Insert(44)
	fmt.Println("\n************* Print *************")
	fmt.Println(l.Sum())
}

func (l *LinkedList) Insert(v int) {
	// Initialize a new node
	n := Node{value: v}
	// fmt.Println(l.len)

	if l.len == 0 {
		// We'll give address of the head node in list
		l.head = &n
		l.len++
		return
	}
	// We'll take pointer from the head of linked list
	ptr := l.head

	for ; ptr.next != nil; ptr = ptr.next {
		// This will take list pointer to end
		// node once it comes out of the loop
	}

	ptr.next = &n
	l.len++
}

func (l *LinkedList) Print() {
	// Initialize a new node
	ptr := l.head
	for ; ptr.next != nil; ptr = ptr.next {
		// This will take list pointer to end
		// node once it comes out of the loop
		fmt.Println(ptr.value)
	}
	fmt.Println(ptr.value)
}

func (l *LinkedList) Sum() int {
	ptr := l.head
	sum := 0

	for i := 0; i < l.len; i++ {
		sum += ptr.value
		ptr = ptr.next
	}
	return sum
}
