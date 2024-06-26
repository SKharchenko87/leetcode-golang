package p1052

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	sum := 0
	prevSum := 0
	maxSum := 0
	for i, customer := range customers {
		sum += (1 - grumpy[i]) * customer
		prevSum += grumpy[i] * customer
		if i >= minutes {
			prevSum -= grumpy[i-minutes] * customers[i-minutes]
		}
		if prevSum > maxSum {
			maxSum = prevSum
		}
	}
	return sum + maxSum
}

func maxSatisfied1(customers []int, grumpy []int, minutes int) int {
	sum := 0
	prevSum := 0
	maxSum := 0
	for i, customer := range customers {
		if grumpy[i] == 0 {
			sum += customer
		}

		if prevSum > maxSum {
			maxSum = prevSum
		}
		prevSum += grumpy[i] * customer
		if i >= minutes {
			prevSum -= grumpy[i-minutes] * customers[i-minutes]
		}
	}
	if prevSum > maxSum {
		maxSum = prevSum
	}

	return sum + maxSum
}

func maxSatisfied0(customers []int, grumpy []int, minutes int) int {
	sum := 0
	prevSum := 0
	maxSum := 0
	for i, customer := range customers {
		sum += (1 - grumpy[i]) * customer
		if prevSum > maxSum {
			maxSum = prevSum
		}
		prevSum += grumpy[i] * customer
		if i >= minutes {
			prevSum -= grumpy[i-minutes] * customers[i-minutes]
		}
	}
	if prevSum > maxSum {
		maxSum = prevSum
	}

	return sum + maxSum
}
