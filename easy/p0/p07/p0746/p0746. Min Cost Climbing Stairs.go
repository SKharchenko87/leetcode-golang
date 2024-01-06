package p0746

func minCostClimbingStairs(cost []int) int {
	switch len(cost) {
	case 0:
		return 0
	case 1:
		return cost[0]
	case 2:
		return min(cost[0], cost[1])
	default:
		sum := 0
		l := len(cost)
		for i := 0; i < l; {
			if i == len(cost)-1 {
				cost = append(cost, 0, 0)
			} else if i == len(cost)-2 {
				cost = append(cost, 0)
			}
			if cost[i]+cost[i+2] <= cost[i+1] {
				sum += cost[i] + cost[i+2]
				i += 3
			} else {
				sum += cost[i+1]
				i += 2
			}
		}
		return sum
	}
}
