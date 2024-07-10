package p1701

func averageWaitingTime(customers [][]int) float64 {
	chefTime := 0
	sumTime := 0
	for _, customer := range customers {
		arrival := customer[0]
		time := customer[1]
		if chefTime <= arrival {
			sumTime += time
			chefTime += time + (arrival - chefTime)
		} else {
			sumTime += time + (chefTime - arrival)
			chefTime += time
		}
	}
	return float64(sumTime) / float64(len(customers))
}
