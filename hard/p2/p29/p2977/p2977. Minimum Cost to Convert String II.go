package p2977

import "sort"

var pow []uint64

//var globalHash map[string]uint64

func init() {
	//globalHash = make(map[string]uint64, 200)
	pow = make([]uint64, 1001)
	pow[0] = 1
	for i := 1; i < 1001; i++ {
		pow[i] = (pow[i-1] * base) % mod
	}
}

type Change struct {
	originalHash uint64
	changedHash  uint64
	cost         uint64
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	// Собираем в группы по длине уникальные строки из original и changed
	mOCCLen := make(map[int][]Change)
	lengths := []int{} // все возможные длины
	lo := len(original)
	for i := 0; i < lo; i++ {
		l := len(original[i])
		if _, ok := mOCCLen[l]; !ok {
			mOCCLen[l] = make([]Change, 0, 1)
			lengths = append(lengths, l)
		}
		mOCCLen[l] = append(mOCCLen[l], Change{hash(original[i]), hash(changed[i]), uint64(cost[i])})
	}

	if _, ok := mOCCLen[1]; !ok {
		lengths = append(lengths, 1)
	}

	sort.Ints(lengths)

	// Вычисляем миниимальные косты для всех возможных пар original -> changed
	minCost := make(map[int]map[[2]uint64]uint64, len(mOCCLen))
	for l, changes := range mOCCLen {
		minCost[l] = execFloydWarshall(changes)
	}

	sourcePrefixHash := getPrefixHash(source)
	targetPrefixHash := getPrefixHash(target)

	// подготавливем dp - в нем будет храниться минимальное значение возможное после изменения строки в символе i+1
	sourceLen := len(source)
	dp := make([]uint64, sourceLen+1)
	for i := 1; i <= sourceLen; i++ {
		dp[i] = inf
	}
	for i := 0; i < sourceLen; i++ {
		// Т.к. возможно изменение только одинаковых по длине фрагментов проходим по длинам
		for index := 0; index < len(lengths); index++ {
			j := i - lengths[index] + 1 // начало подстарки
			if j < 0 {
				break
			}

			rightSourceHash := substringHash(sourcePrefixHash, pow, j, i)
			rightTargetHash := substringHash(targetPrefixHash, pow, j, i)

			if rightSourceHash == rightTargetHash && dp[j] != inf {

				dp[i+1] = min(dp[i+1], dp[j])

				continue
			}

			key := [2]uint64{rightSourceHash, rightTargetHash}
			if v, ok := minCost[i-j+1][key]; ok && dp[j] != inf {

				dp[i+1] = min(dp[i+1], dp[j]+v)

			}
		}
	}
	if dp[sourceLen] == inf {
		return -1
	}
	return int64(dp[sourceLen])
}

const base = uint64(31)
const mod = uint64(1e9 + 7)
const offset = uint8('a' - 1)

func substringHash(prefixHash []uint64, pow []uint64, l, r int) uint64 {
	return (prefixHash[r+1] + mod - (prefixHash[l]*pow[r-l+1])%mod) % mod
}

func getPrefixHash(s string) []uint64 {
	l := len(s)
	h := make([]uint64, l+1)
	h[0] = 0
	for i := 0; i < l; i++ {
		h[i+1] = (base*h[i] + uint64(s[i]-offset)) % mod
	}
	return h
}

func hash(s string) uint64 {
	//if v, ok := globalHash[s]; ok {
	//	return v
	//}
	var h uint64
	for i := 0; i < len(s); i++ {
		h = (base*h + uint64(s[i]-offset)) % mod
	}
	//globalHash[s] = h
	return h
}

const inf = uint64(1<<63 - 1)

func execFloydWarshall(changes []Change) map[[2]uint64]uint64 {
	set := map[uint64]int{}
	keys := []uint64{}
	index := 0
	addSet := func(h uint64) {
		if _, ok := set[h]; !ok {
			keys = append(keys, h)
			set[h] = index
			index++
		}
	}
	changeLength := len(changes)
	for i := 0; i < changeLength; i++ {
		addSet(changes[i].originalHash)
		addSet(changes[i].changedHash)
	}
	l := index
	data := make([]uint64, l*l)
	matrix := make([][]uint64, l)
	for i := 0; i < l; i++ {
		matrix[i] = data[i*l : i*l+l]
	}
	for i := 0; i < l*l; i++ {
		data[i] = inf
	}
	for i := 0; i < l; i++ {
		matrix[i][i] = 0
	}

	for i := 0; i < changeLength; i++ {
		oi := set[changes[i].originalHash]
		ci := set[changes[i].changedHash]
		matrix[oi][ci] = min(matrix[oi][ci], changes[i].cost)
	}

	size := 0
	for k := 0; k < l; k++ {
		for i := 0; i < l; i++ {
			if matrix[i][k] == inf {
				continue
			}
			for j := 0; j < l; j++ {
				if matrix[k][j] == inf {
					continue
				}
				if matrix[i][j] > matrix[i][k]+matrix[k][j] {
					if matrix[i][j] == inf {
						size++
					}
					matrix[i][j] = matrix[i][k] + matrix[k][j]
				}
			}
		}
	}

	res := make(map[[2]uint64]uint64, size)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if v := matrix[i][j]; v != inf {
				k := [2]uint64{keys[i], keys[j]}
				res[k] = v
			}
		}
	}

	return res
}
