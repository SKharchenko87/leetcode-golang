package p3100

func maxBottlesDrunk(numBottles int, numExchange int) int {
	fullBottles := 0
	emptyBottles := numBottles
	bottlesDrunk := numBottles
	for emptyBottles >= numExchange {
		for emptyBottles >= numExchange {
			emptyBottles -= numExchange
			fullBottles++
			numExchange++
		}
		bottlesDrunk += fullBottles
		emptyBottles += fullBottles
		fullBottles = 0
	}
	return bottlesDrunk
}
