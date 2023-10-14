package main

import (
	"math"
)

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
			high = mid
		} else if val < target {
			low = mid + 1
		}
	}

	return false
}
