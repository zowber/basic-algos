package main

import (
	"math"
)

// checks if an array contains the target

func Binary_Search(arr []int, target int) bool {

	low := 0
	high := len(arr)

	for low < high {

		lowF := float64(low)
		highF := float64(high)

		mid := int(math.Floor(lowF + (highF-lowF)/2))
		val := arr[mid]

		if val == target {
			return true
		} else if val > target {
			// the target is lower than mid point
			// make the mid point the new high bound
			high = mid
		} else if val < target {
			// the target is higher than the mid point
			// make mid + 1 the new low bound
			low = mid + 1
		}
	}

	return false
}
