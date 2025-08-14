package p2594

import (
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	low := int64(1)
	high := int64(1e18)
	for low < high {
		mid := low + (high-low)/2
		var carsRepaired int64 = 0
		for _, r := range ranks {
			carsRepaired += int64(math.Floor(math.Sqrt(float64(mid) / float64(r))))
		}
		if carsRepaired >= int64(cars) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
