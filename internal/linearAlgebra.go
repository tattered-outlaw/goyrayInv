package internal

import (
	"fmt"
	"math"
)

const EPSILON = 0.00001

func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

type Tuple [4]float64

func (t Tuple) IsPoint() bool {
	return t[3] == 1
}

func (t Tuple) IsVector() bool {
	return t[3] == 0
}

func (t Tuple) Add(u Tuple) Tuple {
	return Tuple{t[0] + u[0], t[1] + u[1], t[2] + u[2], t[3] + u[3]}
}

func (t Tuple) Sub(u Tuple) Tuple {
	return Tuple{t[0] - u[0], t[1] - u[1], t[2] - u[2], t[3] - u[3]}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t[0], -t[1], -t[2], -t[3]}
}

func (t Tuple) Scale(f float64) Tuple {
	return Tuple{t[0] * f, t[1] * f, t[2] * f, t[3] * f}
}

func (t Tuple) Divide(f float64) Tuple {
	return Tuple{t[0] / f, t[1] / f, t[2] / f, t[3] / f}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t[0]*t[0] + t[1]*t[1] + t[2]*t[2])
}

func (t Tuple) Normalize() Tuple {
	return t.Divide(t.Magnitude())
}

func (t Tuple) Dot(u Tuple) float64 {
	return t[0]*u[0] + t[1]*u[1] + t[2]*u[2] + t[3]*u[3]
}

func (t Tuple) Cross(u Tuple) Tuple {
	return Vector(t[1]*u[2]-t[2]*u[1], t[2]*u[0]-t[0]*u[2], t[0]*u[1]-t[1]*u[0])
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

type Matrix2x2 [2][2]float64

func (m Matrix2x2) Determinant() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

type Matrix3x3 [3][3]float64

func (m Matrix3x3) Submatrix(row, column int) Matrix2x2 {
	result := Matrix2x2{}
	skipRow := 0
	for r := 0; r < 3; r++ {
		if r == row {
			skipRow = 1
			continue
		}
		skipColumn := 0
		for c := 0; c < 3; c++ {
			if c == column {
				skipColumn = 1
				continue
			}
			result[r-skipRow][c-skipColumn] = m[r][c]
		}
	}
	return result
}

func (m Matrix3x3) Minor(row, column int) float64 {
	return m.Submatrix(row, column).Determinant()
}

func (m Matrix3x3) CoFactor(row, column int) float64 {
	minor := m.Minor(row, column)
	if (row+column)%2 == 0 {
		return minor
	} else {
		return -minor
	}
}

func (m Matrix3x3) Determinant() float64 {
	result := 0.0
	for c := 0; c < 3; c++ {
		result += m[0][c] * m.CoFactor(0, c)
	}
	return result
}

type Matrix4x4 [4][4]float64

func (m Matrix4x4) Submatrix(row, column int) Matrix3x3 {
	result := Matrix3x3{}
	skipRow := 0
	for r := 0; r < 4; r++ {
		if r == row {
			skipRow = 1
			continue
		}
		skipColumn := 0
		for c := 0; c < 4; c++ {
			if c == column {
				skipColumn = 1
				continue
			}
			result[r-skipRow][c-skipColumn] = m[r][c]
		}
	}
	return result
}

func (m Matrix4x4) Minor(row, column int) float64 {
	return m.Submatrix(row, column).Determinant()
}

func (m Matrix4x4) CoFactor(row, column int) float64 {
	minor := m.Minor(row, column)
	if (row+column)%2 == 0 {
		return minor
	} else {
		return -minor
	}
}

func (m Matrix4x4) Determinant() float64 {
	result := 0.0
	for c := 0; c < 4; c++ {
		result += m[0][c] * m.CoFactor(0, c)
	}
	return result

}

var Identity4 = Matrix4x4{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func newIdentity4() *Matrix4x4 {
	i := Identity4
	// return address of copy for mutation
	return &i
}

func (m Matrix4x4) MulT(t Tuple) Tuple {
	result := Tuple{}
	for r := 0; r < 4; r++ {
		result[r] = m[r][0]*t[0] + m[r][1]*t[1] + m[r][2]*t[2] + m[r][3]*t[3]
	}
	return result
}

func (m Matrix4x4) Transpose() Matrix4x4 {
	result := Matrix4x4{}
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			result[r][c] = m[c][r]
		}
	}
	return result
}

func transpose4(m *Matrix4x4) *Matrix4x4 {
	mT := m.Transpose()
	return &mT
}

func (m *Matrix4x4) mul4x4(n *Matrix4x4) *Matrix4x4 {
	result := Matrix4x4{}
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			result[r][c] = m[r][0]*n[0][c] + m[r][1]*n[1][c] + m[r][2]*n[2][c] + m[r][3]*n[3][c]
		}
	}
	return &result
}

func inverse4(m *Matrix4x4) *Matrix4x4 {
	result, err := m.Inverse()
	if err != nil {
		panic(err)
	}
	return &result
}

func (m Matrix4x4) Inverse() (Matrix4x4, error) {
	result := Matrix4x4{}
	det := m.Determinant()
	if FloatEquals(det, 0) {
		return result, fmt.Errorf("determinant is zero")
	}
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			cf := m.CoFactor(r, c)
			result[c][r] = cf / det
		}
	}
	return result, nil
}

func (m Matrix4x4) Eq(n Matrix4x4) bool {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEquals(m[c][r], n[c][r]) {
				return false
			}
		}
	}
	return true
}

func Translation(x, y, z float64) *Matrix4x4 {
	return &Matrix4x4{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

func (m *Matrix4x4) Translate(x, y, z float64) *Matrix4x4 {
	return Translation(x, y, z).mul4x4(m)
}

func (m *Matrix4x4) TranslateX(x float64) *Matrix4x4 {
	return m.Translate(x, 0, 0)
}

func (m *Matrix4x4) TranslateY(y float64) *Matrix4x4 {
	return m.Translate(0, y, 0)
}

func (m *Matrix4x4) TranslateZ(z float64) *Matrix4x4 {
	return m.Translate(0, 0, z)
}

func (m *Matrix4x4) Scale(x, y, z float64) *Matrix4x4 {
	return (&Matrix4x4{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}).mul4x4(m)
}

func Scaling(x, y, z float64) *Matrix4x4 {
	return &Matrix4x4{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

func (m *Matrix4x4) ScaleX(x float64) *Matrix4x4 {
	return m.Scale(x, 1, 1)
}

func (m *Matrix4x4) ScaleY(y float64) *Matrix4x4 { return m.Scale(1, y, 1) }

func (m *Matrix4x4) ScaleZ(z float64) *Matrix4x4 {
	return m.Scale(1, 1, z)
}

func RotationX(theta float64) *Matrix4x4 {
	c := math.Cos(theta)
	s := math.Sin(theta)
	return &Matrix4x4{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	}
}

func RotationY(theta float64) *Matrix4x4 {
	c := math.Cos(theta)
	s := math.Sin(theta)
	return &Matrix4x4{
		{c, 0, s, 0},
		{0, 1, 0, 0},
		{-s, 0, c, 0},
		{0, 0, 0, 1},
	}
}

func RotationZ(theta float64) *Matrix4x4 {
	c := math.Cos(theta)
	s := math.Sin(theta)
	return &Matrix4x4{
		{c, -s, 0, 0},
		{s, c, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}
