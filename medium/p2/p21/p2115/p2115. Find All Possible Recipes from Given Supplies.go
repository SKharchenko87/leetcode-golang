package p2115

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	suppliesMap := make(map[string]bool)
	for _, supply := range supplies {
		suppliesMap[supply] = true
	}
	resultMap := make(map[string]bool, len(recipes))
	cnt := len(suppliesMap)
	for cnt > 0 {
		cnt = 0
	LOOP_RECIPE:
		for i, recipe := range recipes {
			if suppliesMap[recipe] {
				resultMap[recipe] = true
				continue LOOP_RECIPE
			}
			for _, ingredient := range ingredients[i] {
				if !suppliesMap[ingredient] {
					continue LOOP_RECIPE
				}
			}
			suppliesMap[recipe] = true
			resultMap[recipe] = true
			cnt++
		}
	}
	result := make([]string, 0, len(resultMap))
	for s := range resultMap {
		result = append(result, s)
	}
	return result
}
