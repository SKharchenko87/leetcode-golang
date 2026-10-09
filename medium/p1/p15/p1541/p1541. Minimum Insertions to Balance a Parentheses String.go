package p1541

func minInsertions(s string) int {
	cl, op := 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if cl == 1 {
				if op == 0 {
					res += 2
				}
				if op > 0 { //     ()(
					res++
					op--
				}
			}
			op++ //    (   |  (( |  ((( ...
			cl = 0
		} else {
			cl++
			if cl == 2 { //    ())
				if op == 0 {
					res++
				} else {
					op--
				}
				cl = 0
			}
		}
	}

	// --- Исправленный блок после цикла ---
	if cl == 1 {
		if op > 0 {
			res += 1 // <-- Исправлено: нужна только одна ')'
			op--
		} else {
			res += 2 // Нет открывающих, нужны одна '(' и одна ')'
		}
		cl = 0
	}

	if op > 0 {
		res += 2 * op // <-- Исправлено: добавлен знак умножения
	}

	return res
}
