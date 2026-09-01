package lc3568

func minMoves(classroom []string, energy int) int {
	m := len(classroom)
	n := len(classroom[0])

	type state struct {
		x               int
		y               int
		mask            int
		remainingEnergy int
	}

	start := state{
		x:               -1,
		y:               -1,
		mask:            0,
		remainingEnergy: energy,
	}

	litters := [][]int{}
	litterIndex := make(map[[2]int]int)
	for i, row := range classroom {
		for j, cell := range row {
			switch cell {
			case 'S':
				start.x = j
				start.y = i
			case 'L':
				litterIndex[[2]int{i, j}] = len(litters)
				litters = append(litters, []int{i, j})
			}
		}
	}

	k := len(litters)

	bestEnergy := make([][][]int, n)
	for x := range bestEnergy {
		bestEnergy[x] = make([][]int, m)
		for y := range bestEnergy[x] {
			bestEnergy[x][y] = make([]int, 1<<k)
			for mask := range 1 << k {
				bestEnergy[x][y][mask] = -1
			}
		}
	}

	shortestMove := func(start state) int {
		queue := []state{start}
		grid := make([][]byte, len(classroom))
		level := 0
		for i := range classroom {
			grid[i] = []byte(classroom[i])
		}
		dirs := [][]int{
			{0, 1},
			{1, 0},
			{-1, 0},
			{0, -1},
		}

		fullMask := (1 << k) - 1

		for len(queue) > 0 {
			qLen := len(queue)
			for range qLen {
				front := queue[0]
				queue = queue[1:]

				switch grid[front.y][front.x] {
				case 'L':
					idx := litterIndex[[2]int{front.y, front.x}]
					front.mask = front.mask | (1 << idx)
				case 'R':
					front.remainingEnergy = energy
				}

				if front.mask == fullMask {
					return level
				}

				if front.remainingEnergy == 0 && classroom[front.y][front.x] != 'R' {
					continue
				}

				for _, d := range dirs {
					neighborX := front.x - d[0]
					neighborY := front.y - d[1]

					if neighborX < 0 || neighborX >= n || neighborY < 0 || neighborY >= m {
						continue
					}

					if grid[neighborY][neighborX] == 'X' {
						continue
					}

					best := bestEnergy[neighborX][neighborY][front.mask]
					newEnergy := front.remainingEnergy - 1
					if newEnergy > best {
						bestEnergy[neighborX][neighborY][front.mask] = newEnergy
						queue = append(queue, state{
							x:               neighborX,
							y:               neighborY,
							mask:            front.mask,
							remainingEnergy: newEnergy,
						})
					}
				}
			}

			level++
		}

		return -1
	}

	res := shortestMove(start)
	return res
}
