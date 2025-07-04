package p3307

import "math/bits"

func kthCharacter(k int64, operations []int) byte {
	k--
	importantOperationN := bits.Len64(uint64(k))
	positions := make([]int64, importantOperationN)
	for i := importantOperationN - 1; i >= 0; i-- {
		positions[i] = k
		if k >= 1<<i {
			k = k - 1<<i
		}
	}
	result := byte('a')
	for i := 0; i < importantOperationN; i++ {
		if positions[i] >= 1<<i {
			if operations[i] == 1 {
				if result == 'z' {
					result = 'a'
				} else {
					result++
				}
			}
		}

	}
	return result
}

/*

Input: k = 10, operations = [0,1,0,1]

Output: "b"

Explanation:

Initially, word == "a". Alice performs the four operations as follows:

Appends "a" to "a", word becomes "aa".
Appends "bb" to "aa", word becomes "aabb".
Appends "aabb" to "aabb", word becomes "aabbaabb".
Appends "bbccbbcc" to "aabbaabb", word becomes "aabbaabbbbccbbcc".

10=0b1010

         _
aabbaabbbbccbbcc
 _
aabbaabb
 _
aabb
 _
aa


*/
