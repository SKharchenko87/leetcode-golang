package p0374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
var cur int

func guess(num int) int {
	//fmt.Println(cur, num)
	if num < cur {
		return 1
	} else if num > cur {
		return -1
	}
	return 0
}

func guessNumber(n int) int {
	l, r := 0, n
	for {
		//fmt.Println(l, n, r)
		switch guess(n) {
		case -1:
			r, n = n, (r+l)/2
		case 1:
			l, n = n, (n+(r+l))/2
		default:
			return n
		}
	}
}

func run(x, n int) int {
	cur = n
	return guessNumber(x)
}
