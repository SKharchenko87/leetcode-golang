package p1518

func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}

func numWaterBottles0(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles/numExchange > 0 {
		res += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return res
}
