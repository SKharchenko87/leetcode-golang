package p2069

type Robot struct {
	width, height int
	perimeter     int
	current       int
	hasMoved      bool
}

func Constructor(width int, height int) Robot {
	p := 2 * (width + height - 2)
	return Robot{width, height, p, 0, false}
}

func (this *Robot) Step(num int) {
	this.hasMoved = true
	this.current = (this.current + num) % this.perimeter
}

func (this *Robot) GetPos() []int {
	curr := this.current
	w, h := this.width, this.height

	// 1. Нижняя сторона (East)
	if curr < w {
		return []int{curr, 0}
	}
	// 2. Правая сторона (North)
	if curr < w+h-1 {
		return []int{w - 1, curr - (w - 1)}
	}
	// 3. Верхняя сторона (West)
	if curr < 2*w+h-2 {
		return []int{w - 1 - (curr - (w + h - 2)), h - 1}
	}
	// 4. Левая сторона (South)
	return []int{0, h - 1 - (curr - (2*w + h - 3))}
}

func (this *Robot) GetDir() string {
	if !this.hasMoved {
		return "East"
	}

	curr := this.current
	w, h := this.width, this.height

	if curr >= 1 && curr <= w-1 {
		return "East"
	} else if curr >= w && curr <= w+h-2 {
		return "North"
	} else if curr >= w+h-1 && curr <= 2*w+h-3 {
		return "West"
	}
	return "South"
}

func run(commands []string, values [][]int) []interface{} {
	res := make([]interface{}, len(values))
	var robot Robot
	for i, cmd := range commands {
		switch cmd {
		case "Robot":
			robot = Constructor(values[i][0], values[i][1])
			res[i] = nil
		case "step":
			robot.Step(values[i][0])
			res[i] = nil
		case "getDir":
			res[i] = robot.GetDir()
		case "getPos":
			res[i] = robot.GetPos()
		default:
			panic("invalid command")
		}
	}
	return res
}
