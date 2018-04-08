package dim3

import (
	"math"

	"github.com/lysrt/bricks/matrix"
)

func Rx(theta float64) matrix.Matrix {
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)

	rx := matrix.Zeros(3, 3)
	rx.Set(0, 0, 1)
	rx.Set(0, 1, 0)
	rx.Set(0, 2, 0)

	rx.Set(1, 0, 0)
	rx.Set(1, 1, cosT)
	rx.Set(1, 2, -sinT)

	rx.Set(2, 0, 0)
	rx.Set(2, 1, sinT)
	rx.Set(2, 2, cosT)

	return rx
}

func Ry(theta float64) matrix.Matrix {
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)

	ry := matrix.Zeros(3, 3)
	ry.Set(0, 0, cosT)
	ry.Set(0, 1, 0)
	ry.Set(0, 2, sinT)

	ry.Set(1, 0, 0)
	ry.Set(1, 1, 1)
	ry.Set(1, 2, 0)

	ry.Set(2, 0, -sinT)
	ry.Set(2, 1, 0)
	ry.Set(2, 2, cosT)

	return ry
}

func Rz(theta float64) matrix.Matrix {
	cosT := math.Cos(theta)
	sinT := math.Sin(theta)

	rz := matrix.Zeros(3, 3)
	rz.Set(0, 0, cosT)
	rz.Set(0, 1, -sinT)
	rz.Set(0, 2, 0)

	rz.Set(1, 0, sinT)
	rz.Set(1, 1, cosT)
	rz.Set(1, 2, 0)

	rz.Set(2, 0, 0)
	rz.Set(2, 1, 0)
	rz.Set(2, 2, 1)

	return rz
}
