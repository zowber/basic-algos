package bubble_sort

import (
	"fmt"
	"math/rand"
)

func Test() {
	// create an array of 100 random integers
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = rand.Intn(1000000)
	}
	// sort the array
	res := Bubble_Sort(arr)

	// check if the array is sorted
	for j := 0; j < len(res)-1; j++ {
		if res[j] > res[j+1] {
			fmt.Println("Not sorted!")
			return
		}
	}
	fmt.Println("Sorted!")
}
