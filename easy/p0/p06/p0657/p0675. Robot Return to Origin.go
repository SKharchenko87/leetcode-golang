package p0657

const (
	Up    = 'U'
	Down  = 'D'
	Left  = 'L'
	Right = 'R'
)

func judgeCircle(moves string) bool {
	vertical, horizon := 0, 0
	for _, ch := range moves {
		switch ch {
		case Up:
			vertical++
		case Down:
			vertical--
		case Left:
			horizon++
		case Right:
			horizon--
		}
	}
	return vertical == 0 && horizon == 0
}
