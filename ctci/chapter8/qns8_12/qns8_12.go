package qns8_12

func NQueens(n int) [][]string {
	var res [][]string

	var iter func(row int, diagonals, antiDiagonals, cols map[int]struct{}, state [][]rune)
	iter = func(row int, diagonals, antiDiagonals, cols map[int]struct{}, state [][]rune) {
		if row == n {
			res = append(res, createGrid(state))
			return
		}

		for col := 0; col < n; col++ {
			currDiagonal := row - col
			currAntiDiagonal := row + col

			if _, ok := cols[col]; ok {
				continue
			}
			if _, ok := diagonals[currDiagonal]; ok {
				continue
			}
			if _, ok := antiDiagonals[currAntiDiagonal]; ok {
				continue
			}

			cols[col] = struct{}{}
			diagonals[currDiagonal] = struct{}{}
			antiDiagonals[currAntiDiagonal] = struct{}{}
			state[row][col] = 'Q'

			iter(row+1, diagonals, antiDiagonals, cols, state)

			delete(cols, col)
			delete(diagonals, currDiagonal)
			delete(antiDiagonals, currAntiDiagonal)
			state[row][col] = '.'
		}
	}

	emptyBoard := make([][]rune, n)
	for i := 0; i < n; i++ {
		emptyBoard[i] = make([]rune, n)
		for j, _ := range emptyBoard[i] {
			emptyBoard[i][j] = '.'
		}
	}
	iter(0, make(map[int]struct{}), make(map[int]struct{}), make(map[int]struct{}), emptyBoard)
	return res
}

func createGrid(state [][]rune) []string {
	var grid []string
	for _, row := range state {
		grid = append(grid, string(row))
	}
	return grid
}
