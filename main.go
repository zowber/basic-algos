package main

func main() {

	// Stack

	// Queue
	// q := new(Queue)

	// q.Enqueue(5)
	// q.Enqueue(10)
	// q.Enqueue(20)

	// q.Dequeue()
	// q.Dequeue()
	// q.Dequeue()

	// Bubble sort
	// Create an array of 500 random integers
	// arr := make([]int, 100)
	// for i := range arr {
	// 	arr[i] = rand.Intn(1000000)
	// }
	// res := Bubble_Sort(arr)
	// fmt.Println(res)

	// Singly linked list

	l := new(ListSl)

	for i := 0; i < 10; i++ {
		l.Insert(i)
	}

	// l.Walk()

	l.InsertAt(5, 20)

	l.Walk()

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
