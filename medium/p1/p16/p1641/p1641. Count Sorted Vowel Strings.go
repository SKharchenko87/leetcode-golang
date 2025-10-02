package p1641

var answer = make([]int, 51)

func init() {
	prevA, a := 0, 1
	prevE, e := 0, 1
	prevI, i := 0, 1
	prevO, o := 0, 1
	u := 1
	for j := 1; j < 51; j++ {
		answer[j] = a + e + i + o + u
		prevA = a
		prevE = e
		prevI = i
		//prevU = u
		prevO = o
		a = answer[j]
		e = a - prevA
		i = e - prevE
		o = i - prevI
		u = o - prevO
	}
}

func countVowelStrings(n int) int {
	return answer[n]
}

// Через комбинаторику
// C(n,k) — это число способов выбрать k элементов из n без учёта порядка (т. е. количество сочетаний).
// n! / (k! * (n-k)!)
// Т.к. у нас 5 букв, то k=4 (0,1,2,3,4) =>
// => n! / (4!*(n-4)!) =>
// => ((n + 1) * (n + 2) * (n + 3) * (n + 4)) / (1 * 2 * 3 * 4)
func countVowelStrings0(n int) int {
	return ((n + 1) * (n + 2) * (n + 3) * (n + 4)) / (1 * 2 * 3 * 4)
}
