package p0399

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]map[string]float64{}
	var addToGraph = func(divisible, divider string, value float64) {
		if m, ok := graph[divisible]; !ok {
			graph[divisible] = map[string]float64{divider: value}
		} else {
			m[divider] = value
		}
	}
	for i, arr := range equations {
		addToGraph(arr[0], arr[1], values[i])
		addToGraph(arr[1], arr[0], 1/values[i])
	}

	var dfs func(divisible string, divider string, path map[string]bool) float64
	dfs = func(divisible string, divider string, path map[string]bool) float64 {
		path[divisible] = true
		if v, ok := graph[divisible][divider]; ok {
			return v
		} else {
			for k, val := range graph[divisible] {
				if _, ok := path[k]; ok {
					continue
				}
				tmp := dfs(k, divider, path)
				if tmp != -1.0 {
					return val * tmp
				}
			}
		}
		return -1.0
	}

	lq := len(queries)
	res := make([]float64, lq)
	for i := 0; i < lq; i++ {
		divisible, divider := queries[i][0], queries[i][1]
		_, okDivisible := graph[divisible]
		_, okDivider := graph[divider]
		if !okDivisible || !okDivider {
			res[i] = -1.0
		} else {
			if divisible == divider {
				res[i] = 1.0
			} else {
				res[i] = dfs(divisible, divider, map[string]bool{})
			}
		}
	}
	return res
}
