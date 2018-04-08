package dim3

import "github.com/lysrt/bricks/matrix"

type Point struct {
	X, Y, Z float64
}

func (p Point) ToVector() matrix.Matrix {
	m := matrix.Zeros(3, 1)
	m.Set(0, 0, p.X)
	m.Set(1, 0, p.Y)
	m.Set(2, 0, p.Z)
	return m
}
