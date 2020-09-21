package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{5, 8, -5}
	endRoids := asteroids(input)
	fmt.Printf("endRoids: %v", endRoids)
}

func asteroids(asteroids []int) []int {
	newRoids := make([]int, 0)
	for index, value := range asteroids {
		if index == 0 {
			newRoids = append(newRoids, value)
			// skip the first item, we don't have a "previous" to compare with.
			continue
		}

		previous := newRoids[len(newRoids)-1]

		// are the signs different?
		if value*previous > 0 {
			// they are the same direction, won't collide, skip.
			newRoids = append(newRoids, value)
			continue
		}

		if math.Abs(float64(value)) > math.Abs(float64(previous)) {
			// This value is bigger than the previous one. Remove the previous one.
			newRoids = newRoids[:len(newRoids)-1]
		}
	}
	return newRoids
}
