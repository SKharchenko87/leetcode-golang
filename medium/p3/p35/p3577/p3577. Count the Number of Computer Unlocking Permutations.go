package p3577

const MOD = 1e9 + 7

func countPermutations(complexity []int) int {
	root := complexity[0]
	res := 1
	for i := 1; i < len(complexity); i++ {
		if complexity[i] <= root {
			return 0
		}
		res = (res * i) % MOD
	}
	return res
}
