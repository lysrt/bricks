package matrix

import "math/rand"
import "fmt"

// Matrix is a matrix
type Matrix struct {
	Rows  int
	Cols  int
	Cells []float64
}

// Zeros generates a zeroed Matrix
func Zeros(r, c int) Matrix {
	return Matrix{
		Rows:  r,
		Cols:  c,
		Cells: make([]float64, r*c),
	}
}

// RandomN generates a random r by c Matrix
func RandomN(r, c int) Matrix {
	m := Zeros(r, c)

	for i := range m.Cells {
		m.Cells[i] = rand.Float64()
	}

	return m
}

// has returns true if the matrix has a cell with row r and column c
func (m Matrix) has(r, c int) bool {
	return r >= 0 && c >= 0 && r < m.Rows && c < m.Cols
}

// At gets the value of the matrix at row r and column c
func (m Matrix) At(r, c int) float64 {
	if !m.has(r, c) {
		panic(fmt.Sprintf("Cannot get (%d, %d) in (%d, %d) matrix", r, c, m.Rows, m.Cols))
	}
	return m.at(r, c)
}

func (m Matrix) at(r, c int) float64 {
	return m.Cells[m.Cols*r+c]
}

// Set sets the value of the matrix at row r and column c to value
func (m Matrix) Set(r, c int, value float64) {
	if !m.has(r, c) {
		panic(fmt.Sprintf("Cannot set (%d, %d) in (%d, %d) matrix", r, c, m.Rows, m.Cols))
	}
	m.set(r, c, value)
}

func (m Matrix) set(r, c int, value float64) {
	m.Cells[m.Cols*r+c] = value
}

// Dot returns the product of 2 matrices
func (m Matrix) Dot(other Matrix) Matrix {
	if m.Cols != other.Rows {
		panic(fmt.Sprintf("Cannot dot matrixes (%d, %d) and (%d, %d)", m.Rows, m.Cols, other.Rows, other.Cols))
	}

	r := Zeros(m.Rows, other.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < other.Cols; j++ {
			sum := float64(0.0)
			for k := 0; k < m.Cols; k++ {
				a := m.at(i, k)
				b := other.at(k, j)
				sum += a * b
			}
			r.set(i, j, sum)
		}
	}
	return r
}

// Add returns the sum of 2 matrices
func (m Matrix) Add(other Matrix) Matrix {
	if m.Cols != other.Cols || m.Rows != other.Rows {
		panic(fmt.Sprintf("Cannot add matrixes (%d, %d) and (%d, %d) of different sizes", m.Rows, m.Cols, other.Rows, other.Cols))
	}

	r := Zeros(m.Rows, m.Cols)
	for i := range r.Cells {
		r.Cells[i] = m.Cells[i] + other.Cells[i]
	}

	return r
}

// Sub returns the difference of 2 matrices
func (m Matrix) Sub(other Matrix) Matrix {
	if m.Cols != other.Cols || m.Rows != other.Rows {
		panic(fmt.Sprintf("Cannot substract matrixes (%d, %d) and (%d, %d) of different sizes", m.Rows, m.Cols, other.Rows, other.Cols))
	}

	r := Zeros(m.Rows, m.Cols)
	for i := range r.Cells {
		r.Cells[i] = m.Cells[i] - other.Cells[i]
	}

	return r
}

// Multiply returns the elementwise multiplication of 2 matrices
func (m Matrix) Multiply(other Matrix) Matrix {
	if m.Cols != other.Cols || m.Rows != other.Rows {
		panic(fmt.Sprintf("Cannot multiply matrixes (%d, %d) and (%d, %d) of different sizes", m.Rows, m.Cols, other.Rows, other.Cols))
	}

	r := Zeros(m.Rows, m.Cols)
	for i := range r.Cells {
		r.Cells[i] = m.Cells[i] * other.Cells[i]
	}

	return r
}

// Transpose returns the transposed Matrix
func (m Matrix) Transpose() Matrix {
	r := Zeros(m.Cols, m.Rows) // Inverted from m
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			r.set(j, i, m.at(i, j))
		}
	}
	return r
}

// Scale returns the scaled Matrix
func (m Matrix) Scale(scalar float64) Matrix {
	r := Zeros(m.Rows, m.Cols)
	for i := range r.Cells {
		r.Cells[i] = m.Cells[i] * scalar
	}
	return r
}

// String prints a matrix
func (m Matrix) String() string {
	s := fmt.Sprintf("Matrix (%d, %d)\n[", m.Rows, m.Cols)
	for i, c := range m.Cells {
		if i%m.Cols == 0 && i != 0 {
			s += fmt.Sprint(",\n ")
		}
		s += fmt.Sprintf("%v ", c)
	}
	s += fmt.Sprintln("]")
	return s
}

func ArrayToMatrix(a []float64) Matrix {
	m := Zeros(len(a), 1)
	for i, v := range a {
		m.set(i, 0, v)
	}
	return m
}

func MatrixToArray(m Matrix) []float64 {
	if m.Cols != 1 {
		panic(fmt.Sprintf("cannot vectorize a matrix with %d columns", m.Cols))
	}
	var result []float64
	for _, v := range m.Cells {
		result = append(result, v)
	}
	return result
}
