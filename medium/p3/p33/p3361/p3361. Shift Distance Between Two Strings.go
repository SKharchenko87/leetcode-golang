package p3361

/*

 a b  c  d  e f g h i j k l m  n  o p q r s  t u v  w x y z
 1,9,10, 7,10,8,1,9,1,3,3,6,4,10,10,5,6,6,8, 5,7,3,10,5,6,3
 7,1, 2,10, 2,1,5,8,6,8,6,7,4, 6, 5,1,1,5,8,10,1,8, 4,2,3,6

s->m


z->s


f->s

s->f

f->m


m->z


   f   s   z   f   s   f   m   f   s
   f   m   s   s   f   m   z   m   f
   0  26  34  57  70  31  77  31  70
   0  26  60 117 187 218 285 316 386
*/

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	//Extra nextCost
	nextTmp := nextCost[:26]
	nextCost = append(nextCost, nextTmp...)
	prev := nextCost[0]
	nextCost[0] = 0
	for i := 1; i < len(nextCost); i++ {
		prev, nextCost[i] = nextCost[i], prev
		nextCost[i] += nextCost[i-1]
	}

	//Extra previousCost
	prevTmp := previousCost[:26]
	previousCost = append(previousCost, prevTmp...)
	prev = previousCost[25]
	previousCost[25] = 0
	for i := 24; i >= -26; i-- {
		prev, previousCost[(52+i)%52] = previousCost[(52+i)%52], prev
		previousCost[(52+i)%52] += previousCost[(52+i+1)%52]
	}

	//Perform
	var res int64
	for i := 0; i < len(s); i++ {
		sb, tb := s[i]-'a', t[i]-'a'
		if sb == tb {
			continue
		} else if sb < tb {
			res += int64(min(nextCost[tb]-nextCost[sb], previousCost[(tb+26)%52]-previousCost[sb]))
		} else { //sb > tb
			res += int64(min(nextCost[(tb+26)%52]-nextCost[sb], previousCost[tb]-previousCost[sb]))
		}
	}

	return res
}

/*
1  3  4  5  7
1  4  8 13 20
21 24 28 33

 0  3  7 12 19
17  0  4  9 16
13 16  0  5 12
 8 11 15  0  7
 1  4  8 13  0

*/

// abcdefghijklmnopqrstuvwxyz
func shiftDistance0(s string, t string, nextCost []int, previousCost []int) int64 {
	var res int64

	for i := 0; i < len(s); i++ {
		var cntNext, cntPrev int
		sb, tb := int(s[i]-'a'), int(t[i]-'a')
		if sb == tb {

		} else if sb <= tb {
			cntNext = tb - sb
			cntPrev = sb + 26 - tb

		} else if sb > tb {
			cntNext = 26 - sb + tb
			cntPrev = sb - tb
		}

		tmpNext := 0
		for j := 0; j < cntNext; j++ {
			tmpNext += nextCost[(sb+j)%26]
		}
		tmpPrev := 0
		for j := 0; j < cntPrev; j++ {
			tmpPrev += previousCost[(26+sb-j)%26]
		}

		res += int64(min(tmpNext, tmpPrev))
	}
	return res
}
