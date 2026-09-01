package p3568

/*ToDo*/
const (
	Start    = 'S'
	Litter   = 'L'
	Recharge = 'R'
	Obstacle = 'X'
	Empty    = '.'

	LitterBit   = 0b01
	RechargeBit = 0b11
	ObstacleBit = 0b10
	EmptyBit    = 0b00
)

// State используется исключительно как ключ для map
type State struct {
	i, j int
	grid [20]int64
}

// Item хранит все данные узла для очереди BFS
type Item struct {
	state   State
	energy  int
	moves   int
	cleaned int
}

func minMoves(classroom []string, energy int) int {
	m, n := len(classroom), len(classroom[0])
	startI, startJ := -1, -1
	countLitter := 0
	base := make([]int64, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch classroom[i][j] {
			case Litter:
				countLitter++
				base[i] |= LitterBit << (2 * j)
			case Recharge:
				base[i] |= RechargeBit << (2 * j)
			case Obstacle:
				base[i] |= ObstacleBit << (2 * j)
			case Empty:
				base[i] |= EmptyBit << (2 * j)
			case Start:
				startI, startJ = i, j
				// Стартовая ячейка проходима как пустая
				base[i] |= EmptyBit << (2 * j)
			}
		}
	}

	// Если убирать нечего, мы уже закончили
	if countLitter == 0 {
		return 0
	}

	// Копируем слайс в массив фиксированной длины для ключа map
	var startGrid [20]int64
	copy(startGrid[:], base)

	startState := State{i: startI, j: startJ, grid: startGrid}

	// visited хранит максимальную энергию для состояния
	visited := make(map[State]int)
	visited[startState] = energy

	queue := []Item{{
		state:   startState,
		energy:  energy,
		moves:   0,
		cleaned: 0,
	}}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Если заряд иссяк, шагнуть из этой точки мы уже не можем
		if curr.energy == 0 {
			continue
		}

		for _, d := range dirs {
			ni, nj := curr.state.i+d[0], curr.state.j+d[1]

			if ni >= 0 && ni < m && nj >= 0 && nj < n {
				// Читаем 2 бита целевой ячейки
				cellType := (curr.state.grid[ni] >> (2 * nj)) & 0b11

				if cellType == ObstacleBit {
					continue
				}

				nextGrid := curr.state.grid
				nextEnergy := curr.energy - 1
				nextCleaned := curr.cleaned

				if cellType == LitterBit {
					// Сбрасываем 2 бита в 00, превращая ячейку в Empty
					nextGrid[ni] &^= int64(3) << (2 * nj)
					nextCleaned++

					// Ранний выход: если собрали всё, возвращаем ответ
					if nextCleaned == countLitter {
						return curr.moves + 1
					}
				} else if cellType == RechargeBit {
					// Полностью восстанавливаем заряд
					nextEnergy = energy
				}

				nextState := State{i: ni, j: nj, grid: nextGrid}

				// Прунинг: идем туда, только если раньше не были там с таким же или бОльшим запасом энергии
				if maxE, exists := visited[nextState]; exists && maxE >= nextEnergy {
					continue
				}

				visited[nextState] = nextEnergy

				queue = append(queue, Item{
					state:   nextState,
					energy:  nextEnergy,
					moves:   curr.moves + 1,
					cleaned: nextCleaned,
				})
			}
		}
	}

	return -1
}
