package qns1_8

type Matrix [][]int

func ZeroMatrix(matrix Matrix) Matrix {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

func (matrix Matrix) equals(other Matrix) bool {
	for i, x := range matrix {
		for j, y := range x {
			if y != other[i][j] {
				return false
			}
		}
	}
	return true
}
