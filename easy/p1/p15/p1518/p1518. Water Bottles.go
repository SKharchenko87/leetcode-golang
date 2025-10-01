package p1518

func numWaterBottles(numBottles int, numExchange int) int {
	res := 0
	for numBottles >= numExchange {
		res += numBottles / numExchange * numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	res += numBottles
	return res
}
