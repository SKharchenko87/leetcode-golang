package p1717

type stack []byte

func (s *stack) Push(v byte) {
	*s = append(*s, v)
}
func (s *stack) Pop() (val byte, exists bool) {
	l := len(*s)
	if l == 0 {
		return 0, false
	}
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v, true
}

func (s *stack) Top() (val byte, exists bool) {
	l := len(*s)
	if l == 0 {
		return 0, false
	}
	v := (*s)[l-1]
	return v, true
}

func calculateGain(s string, x int, chA, chB byte) (newS string, res int) {
	l := len(s)
	var st stack
	st = make([]byte, 0, l)
	for i := 0; i < l; i++ {
		prev, ok := st.Top()
		v := s[i]
		switch v {
		case chA:
		case chB:
			if ok && prev == chA {
				st.Pop()
				res += x
				continue
			}
		default:
			if ok && prev == '+' {
				continue
			}
			v = '+'
		}

		st.Push(v)
	}
	return string(st), res
}

func maximumGain(s string, x int, y int) int {
	res := 0
	chA, chB := byte('a'), byte('b')
	if x < y {
		x, y = y, x
		chA, chB = chB, chA
	}

	newS, newRes := calculateGain(s, x, chA, chB)
	res += newRes
	_, newRes = calculateGain(newS, y, chB, chA)
	res += newRes
	return res
}

// aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha
// aabbkbbbfvybssbtaoaaataaadabbbmakgaboapbbbboabvqhbbzbbkapaavbbeghacamdpaaqbqabbjbmkajasaabbwabrbfuyaazbha 95040
// kbbbfvybssbtaoaaataaadbbmakgoapbbbbovqhbbzbbkapaavbbeghacamdpaaqbqbjbmkajaswrbfuyaazbha 17334

// aabbrtababbabmaaaeaabeawmvaataabnaabbaaaybbbaabbabbbjpjaabbtabbxaaavsmmnblbbabaeuasvababjbbabbabbasxbbtgbrbbajeabbbfbarbagha
// rtbmaaaeaeawmvaatanaaaybbbbbjpjtbxaaavsmmnblbbaeuasvjbbbbasxbbtgbrbbajebbfbarbagha 178164
// rtbmaaaeaeawmvaatanaaaybbbbbjpjtbxaaavsmmnblbeuasvjbbbsxbbtgbrbjebbfrgha 20480+178164=198644

// aabbaaxybbaabb 20
// aaxybb 20

type node struct {
	val  interface{}
	pre  *node
	next *node
}

type Queue struct {
	first *node
	last  *node
}

func newQueue() *Queue {
	return &Queue{nil, nil}
}

func (q *Queue) Push(v interface{}) {
	newNode := &node{v, q.last, nil}
	if q.first == nil {
		q.first = newNode
	}
	if q.last != nil {
		q.last.next = newNode
	}
	q.last = newNode
}

func (q *Queue) Pop() *node {
	if q.last == nil {
		return nil
	}
	v := q.last
	q.last = q.last.pre
	if q.last != nil {
		q.last.next = nil
	}

	if q.last == nil {
		q.first = nil
	}
	return v
}

func (q *Queue) Peek() *node {
	if q.last == nil {
		return nil
	}
	v := q.last
	return v
}

func maximumGain1(s string, x int, y int) int {
	res := 0
	chA, chB := byte('a'), byte('b')
	if x < y {
		x, y = y, x
		chA, chB = chB, chA
	}

	l := len(s)
	q := newQueue()
	for i := 0; i < l; i++ {
		prev := q.Peek()
		var prevVal byte
		if prev != nil {
			prevVal = prev.val.(byte)
		}
		switch v := s[i]; v {
		case chA:
			q.Push(v)
		case chB:
			if prevVal == chA {
				q.Pop()
				res += x
				continue
			}
			q.Push(v)
		default:
			if prevVal == '+' {
				continue
			}
			q.Push(byte('+'))
		}
	}

	root := q.first

	for root != nil {
		next := root.next
		if root.val.(byte) == chA {
			if root.pre != nil && root.pre.val.(byte) == chB {
				if root.pre.pre != nil {
					root.pre.pre.next = next
					if next != nil {
						next.pre = root.pre.pre
					}
				}
				res += y
			}
		}
		root = next
	}

	return res
}

//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd
//afsd

func maximumGain0(s string, x int, y int) int {
	res := 0
	b := []byte(s)
	if x >= y {
		res += ab(&b) * x
		res += ba(&b) * y
	} else {
		res += ba(&b) * y
		res += ab(&b) * x
	}
	return res
}

func ab(b *[]byte) int {
	cnt := 0
	for i := 1; i < len(*b); i++ {
		if (*b)[i-1] == 'a' && (*b)[i] == 'b' {
			cnt++
			*b = append((*b)[:i-1], (*b)[i+1:]...)
			i -= 2
			if i <= 0 {
				i = 0
			}
		}
	}
	return cnt
}

func ba(b *[]byte) int {
	cnt := 0
	for i := 1; i < len(*b); i++ {
		if (*b)[i-1] == 'b' && (*b)[i] == 'a' {
			cnt++
			*b = append((*b)[:i-1], (*b)[i+1:]...)
			i -= 2
			if i <= 0 {
				i = 0
			}
		}
	}
	return cnt
}
