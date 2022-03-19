package math

import (
	"testing"
)

var algebra ByteAlgebra = new(ByteAlgebraImpl)
var galoisAlgebra4 ByteAlgebra = New(4)
var galoisAlgebra8 ByteAlgebra = New(8)

func TestDetTrivialMatrix(t *testing.T) {
	trivialMatrix := [][]byte{
		{1},
	}
	detTrivial := Det(trivialMatrix, algebra)
	if detTrivial != 1 {
		t.Errorf("Det of trivial matrix must be 1 but it's %d", detTrivial)
	}
}

func TestDetZeroMatrix(t *testing.T) {
	zeroMatrix := [][]byte{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	det3ZeroMatrix := Det(zeroMatrix, algebra)
	if det3ZeroMatrix != 0 {
		t.Errorf("Det of zero matrix must be 0 but it's %d", det3ZeroMatrix)
	}
}

func TestDetOneMatrix(t *testing.T) {
	oneMatrix := [][]byte{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	det3OneMatrix := Det(oneMatrix, algebra)
	if det3OneMatrix != 1 {
		t.Errorf("Det of one matrix must be 1 but it's %d", det3OneMatrix)
	}
}

func TestDetMatrix1(t *testing.T) {
	m := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	det := Det(m, algebra)
	if det != 0 {
		t.Errorf("Det of matrix must be 0 but it's %d", det)
	}
}

func TestDetMatrix2(t *testing.T) {
	m := [][]byte{
		{5, 5, 2},
		{5, 3, 2},
		{5, 5, 1},
	}
	det := Det(m, algebra)
	if det != 10 {
		t.Errorf("Det of matrix must be 10 but it's %d", det)
	}
}

func TestDetMatrix3(t *testing.T) {
	m := [][]byte{
		{1, 5, 2, 4, 3},
		{2, 5, 7, 8, 5},
		{2, 6, 4, 8, 7},
		{3, 5, 4, 1, 2},
		{3, 5, 2, 4, 5},
	}
	det := Det(m, algebra)
	if det != 198 {
		t.Errorf("Det3 of matrix must be 198 but it's %d", det)
	}
}

func TestCross(t *testing.T) {
	matrix := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]byte{
		{1, 3},
		{7, 9},
	}
	cross := Cross(matrix, 1, 1)
	if !equal(expected, cross) {
		t.Error("Cross function returns wrong result")
	}
}

func TestMatrixDiv(t *testing.T) {
	matrix := [][]byte{
		{2, 4},
		{6, 8},
	}
	expected := [][]byte{
		{1, 2},
		{3, 4},
	}
	div := div(matrix, 2, algebra)
	if !equal(expected, div) {
		t.Error("Matrix division returns wrong result")
	}
}

func TestReverse(t *testing.T) {
	matrix := [][]byte{
		{0, 1, 0},
		{1, 1, 1},
		{1, 4, 5},
	}
	expected := [][]byte{
		{13, 12, 13},
		{1, 0, 0},
		{12, 13, 13},
	}
	reverse := Reverse(matrix, galoisAlgebra4)
	if !equal(expected, reverse) {
		t.Error("Matrix reverse operation returns wrong result")
	}
}

func TestReverseGF256(t *testing.T) {
	matrix := [][]byte{
		{0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5, 6},
		{1, 8, 15, 64, 85, 120},
	}
	expected := [][]byte{
		{142, 71, 0, 114, 35, 80},
		{1, 0, 0, 0, 0, 0},
		{143, 201, 244, 179, 195, 112},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 143, 245, 192, 224, 32},
	}
	reverse := Reverse(matrix, galoisAlgebra8)
	if !equal(expected, reverse) {
		t.Error("Matrix reverse operation returns wrong result")
	}
}

func TestMatrixMul(t *testing.T) {
	a := [][]byte{
		{1, 1, 1},
		{1, 2, 3},
		{1, 4, 5},
	}
	b := [][]byte{
		{7},
		{3},
		{2},
	}
	mul := Mul(a, b, galoisAlgebra4)
	expected := [][]byte{
		{6},
		{7},
		{1},
	}
	if !equal(expected, mul) {
		t.Error("Matrix multiplication operation returns wrong result")
	}
}

func TestMatrixMulGF256(t *testing.T) {
	a := [][]byte{
		{142, 71, 0, 114, 35, 80},
		{1, 0, 0, 0, 0, 0},
		{143, 201, 244, 179, 195, 112},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 143, 245, 192, 224, 32},
	}
	b := [][]byte{
		{92},
		{21},
		{12},
		{113},
		{164},
		{109},
	}
	expected := [][]byte{
		{78},
		{92},
		{94},
		{21},
		{12},
		{36},
	}
	mul := Mul(a, b, galoisAlgebra8)
	if !equal(expected, mul) {
		t.Error("Matrix multiplication operation returns wrong result")
	}
}

func equal(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[0] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
