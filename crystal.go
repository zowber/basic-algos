package main

import (
	"log"
	"math"
)

// Given two crystal balls that will break if dropped from a high enough floor.
// Determine in the most optimised way the exact floor at which a ball can break.

func Two_Crystal_Balls(floors []bool) int {

	sqrtFloorsFloat := math.Floor(math.Sqrt(float64(len(floors))))
	sqrtFloors := int(sqrtFloorsFloat)
	log.Println("sqrtFloors:", sqrtFloors)

	var firstBallBreaksAt int
	for i := 0; i < len(floors); i += sqrtFloors {
		log.Println("Visiting floor:", i, floors[i])
		if floors[i] {
			firstBallBreaksAt = i
			break
		}
	}
	log.Println("First break on:", firstBallBreaksAt)
	var breaksAtFloor int
	for j := firstBallBreaksAt - sqrtFloors + 1; j < len(floors); j++ {
		log.Println("Visiting floor:", j, floors[j])
		if floors[j] {
			breaksAtFloor = j
			break
		}
	}

	log.Println("Balls will break at floor:", breaksAtFloor)
	return breaksAtFloor
}
