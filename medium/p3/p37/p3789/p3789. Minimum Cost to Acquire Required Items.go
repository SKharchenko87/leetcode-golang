package p3789

func minimumCost(cost1, cost2, costBoth, need1, need2 int) int64 {
	if need1 > need2 {
		cost1, cost2 = cost2, cost1
		need1, need2 = need2, need1
	}
	cost12 := cost1 + cost2
	if costBoth <= cost12 {
		return int64(need1*costBoth) + int64((need2-need1)*min(cost2, costBoth))
	} else {
		return int64(cost1*need1) + int64(cost2*need2)
	}
}
