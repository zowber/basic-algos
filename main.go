package main

import (
	"basic-algos/doublylinkedlist"
	"fmt"
)

func main() {

	list := new(doublylinkedlist.List)

	// append 100_000 vals
	for i := 0; i < 100_000; i++ {
		list.Append(i)
	}
	fmt.Println(list.GetAt(68999))
	fmt.Println(list.GetAt(69000))
	fmt.Println(list.GetAt(69001))
	list.RemoveAt(69000)
	fmt.Println(list.GetAt(68999))
	fmt.Println(list.GetAt(69000))
	fmt.Println(list.GetAt(69001))

	// Two Crystal Balls
	// var floors []bool

	// for i := 0; i < 69420; i++ {
	// 	if i < 42069 {
	// 		floors = append(floors, false)
	// 	} else {
	// 		floors = append(floors, true)
	// 	}
	// }
	// expect 256
	// Two_Crystal_Balls(floors)

	// Stack
	// s := new(stack.Stack)
	// // push 10 million nodes
	// for i := 0; i < 10_000_000; i++ {
	// 	s.Push(i + 1)
	// }
	// // pop 10 million nodes
	// for j := 0; j < 10_000_000; j++ {
	// 	val := s.Pop()
	// 	// Print val of every millionth pop
	// 	if j%1_000_000 == 0 {
	// 		fmt.Println("Popped value:", val)
	// 	}
	// }

	// Queue
	// q := new(Queue)

	// q.Enqueue(5)
	// q.Enqueue(10)
	// q.Enqueue(20)

	// q.Dequeue()
	// q.Dequeue()
	// q.Dequeue()

	// Bubble sort
	// bubble_sort.Test()

	// Singly linked list
	// l := singly_linked_list.List{}
	// for i := 0; i < 10; i++ {
	// 	l.Insert(i)
	// }
	// // l.Walk()
	// l.InsertAt(5, 20)
	// l.Walk()

	// Binary search
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	// searchFor8 := Binary_Search(arr, 8)
	// fmt.Printf("8 exists in array?: %t\n", searchFor8)

	// Linear search
	// searchMe := []int{1, 2, 3, 4, 10}
	// searchFor5 := Linear_Search(searchMe, 5)
	// searchFor10 := Linear_Search(searchMe, 10)
	// fmt.Printf("5 exists in array?: %t\n", searchFor5)
	// fmt.Printf("10 exists in array?: %t\n", searchFor10)

}
