package math

// Det returns det of matrix in specified algebra
func Det(matrix [][]byte, alg ByteAlgebra) byte {
	size := len(matrix)
	if size == 1 {
		return matrix[0][0]
	}
	if size == 2 {
		return det2(matrix, alg)
	}
	var det byte = 0
	for i := 0; i < size; i++ {
		element := matrix[0][i]
		minor := Det(Cross(matrix, 0, i), alg)
		if i%2 == 0 {
			det = alg.Add(det, alg.Mul(element, minor))
		} else {
			det = alg.Sub(det, alg.Mul(element, minor))
		}
	}
	return det
}

// Reverse returns reverse matrix in specified algebra
func Reverse(matrix [][]byte, alg ByteAlgebra) [][]byte {
	return div(transpose(minors(matrix, alg)), Det(matrix, alg), alg)
}

// Mul returns multiplication result of given matrixes in specified algebra
func Mul(a, b [][]byte, alg ByteAlgebra) [][]byte {
	l := len(a)
	m := len(b)
	n := len(b[0])
	result := make([][]byte, l)
	for i := 0; i < l; i++ {
		result[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			var current byte = 0
			for r := 0; r < m; r++ {
				current = alg.Add(current, alg.Mul(a[i][r], b[r][j]))
			}
			result[i][j] = current
		}
	}
	return result
}

// Cross returns matrix with size n-1 crossing row and column specified
func Cross(matrix [][]byte, row, col int) [][]byte {
	size := len(matrix)
	result := make([][]byte, size-1)

	var i, j, a, b int
	for i, a = 0, 0; i < size; i++ {
		if i == row {
			continue
		}
		result[a] = make([]byte, size-1)
		for j, b = 0, 0; j < size; j++ {
			if j == col {
				continue
			}
			result[a][b] = matrix[i][j]
			b++
		}
		a++
	}
	return result
}

func div(matrix [][]byte, divider byte, alg ByteAlgebra) [][]byte {
	size := len(matrix)
	result := make([][]byte, size)

	for i := 0; i < size; i++ {
		result[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			result[i][j] = alg.Div(matrix[i][j], divider)
		}
	}
	return result
}

func transpose(matrix [][]byte) [][]byte {
	size := len(matrix)
	result := make([][]byte, size)

	for i := 0; i < size; i++ {
		result[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

func minors(matrix [][]byte, alg ByteAlgebra) [][]byte {
	size := len(matrix)
	result := make([][]byte, size)

	for i := 0; i < size; i++ {
		result[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			if i+j%2 != 0 {
				result[i][j] = alg.Sub(0, Det(Cross(matrix, i, j), alg))
			} else {
				result[i][j] = alg.Add(0, Det(Cross(matrix, i, j), alg))
			}
		}
	}
	return result
}

// returns det of 2x2 matrix
func det2(matrix [][]byte, alg ByteAlgebra) byte {
	return alg.Sub(alg.Mul(matrix[0][0], matrix[1][1]), alg.Mul(matrix[1][0], matrix[0][1]))
}
