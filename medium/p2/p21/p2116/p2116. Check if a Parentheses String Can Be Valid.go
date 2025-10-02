package p2116

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}
	var open, star, cls int
	for i := 0; i < n; i++ {
		if !exec(&locked, &s, &star, &open, i, '(') {
			return false
		}
	}
	star = 0
	for i := n - 1; i >= 0; i-- {
		if !exec(&locked, &s, &star, &cls, i, ')') {
			return false
		}
	}
	return true
}

func exec(locked, s *string, star, breaks *int, i int, ch byte) bool {
	if (*locked)[i] == '0' {
		(*star)++
	} else if (*s)[i] == ch { // locked[i] == 1
		(*breaks)++
	} else if (*breaks) == 0 { // locked[i] == 1 and s[i] != ch
		(*star)--
	} else { // locked[i] == 1 and s[i] != ch and breaks > 0
		(*breaks)--
	}
	if (*star) < 0 {
		return false
	}
	return true
}

func canBeValid1(s string, locked string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}
	var open, star, cls int
	for i := 0; i < n; i++ {
		if locked[i] == '0' {
			star++
		} else if s[i] == '(' { // locked[i] == 1
			open++
		} else if open == 0 { // locked[i] == 1 and s[i] == ')'
			star--
		} else { // locked[i] == 1 and s[i] == ')' and open > 0
			open--
		}
		if star < 0 {
			return false
		}
	}
	star = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' {
			star++
		} else if s[i] == ')' { // locked[i] == 1
			cls++
		} else if cls == 0 { // locked[i] == 1 and s[i] == '('
			star--
		} else { // locked[i] == 1 and s[i] == '(' and cls > 0
			cls--
		}
		if star < 0 {
			return false
		}
	}

	return true
}

func canBeValid0(s string, locked string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}
	opens := make([]int, 0, n)
	stars := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				opens = append(opens, i)
			} else {
				if len(opens) == 0 {
					if len(stars) == 0 {
						return false
					} else {
						stars = stars[:len(stars)-1]
					}
				} else {
					opens = opens[:len(opens)-1]
				}
			}
		} else {
			stars = append(stars, i)
		}
	}

	io := len(opens) - 1
	if io == -1 {
		return true
	}
	for i := len(stars) - 1; i >= 0; i-- {
		if stars[i] > opens[io] {
			opens = opens[:io]
			io--
			if io == -1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

// *(((**

// (*)*
