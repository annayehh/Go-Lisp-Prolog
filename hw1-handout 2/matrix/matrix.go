package matrix

// If needed, you may define helper functions here.

// AreAdjacent returns true iff a and b are adjacent in lst.
func AreAdjacent(lst []int, a, b int) bool {
	if lst == nil || len(lst) == 0 {
		return false
	}

	for i := 0; i < len(lst)-1; i++ {
		if (lst[i] == a && lst[i+1] == b) || (lst[i] == b && lst[i+1] == a) {
			return true
		}
	}

	return false
}

// Transpose returns the transpose of the 2D matrix mat.
func Transpose(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	rows, cols := len(mat), len(mat[0])
	result := make([][]int, cols)

	for i := range result {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = mat[i][j]
		}
	}

	return result
}

// AreNeighbors returns true iff a and b are neighbors in the 2D matrix mat.
func AreNeighbors(matrix [][]int, a, b int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	n, m := len(matrix), len(matrix[0])

	var aRow, aCol, bRow, bCol int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == a {
				aRow, aCol = i, j
			}
			if matrix[i][j] == b {
				bRow, bCol = i, j
			}
		}
	}

	rowDiff := aRow - bRow
	colDiff := aCol - bCol

	return (rowDiff == 0 && (colDiff == 1 || colDiff == -1)) || (colDiff == 0 && (rowDiff == 1 || rowDiff == -1))
}