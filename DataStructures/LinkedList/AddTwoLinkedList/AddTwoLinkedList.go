package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

var sumLinkedList []LinkedList

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	remainder := 0
	for linkedListOne.Next != nil && linkedListTwo.Next != nil {
		sum := linkedListOne.Value + linkedListTwo.Value
		remainder = sum % 10
		val := sum / 10

		sumLinkedList.Value = val
		sumLinkedList.Next = nil
	}
	return nil
}

func insert(l *LinkedList) *LinkedList {
	n := 
}