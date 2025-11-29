package linearAlgebra

import (
	"fmt"
	"math"
)

const EPSILON = 0.00001

func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

type Tuple [4]float64

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

type Color struct {
	R, G, B float64
}

func (c Color) Add(u Color) Color {
	return Color{c.R + u.R, c.G + u.G, c.B + u.B}
}

func (c Color) Sub(u Color) Color {
	return Color{c.R - u.R, c.G - u.G, c.B - u.B}
}

func (c Color) Scale(f float64) Color {
	return Color{c.R * f, c.G * f, c.B * f}
}

func (c Color) Multiply(u Color) Color {
	return Color{c.R * u.R, c.G * u.G, c.B * u.B}
}

// RGBA Range clips the R, G, B components to be within to [0,1) and converts to uint32
// This method makes Color implement image.Color
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	r := uint32(0)
	if c.R >= 1 {
		r = 0xffff
	} else if c.R > 0 {
		r = uint32(c.R * 0xffff)
	}
	g := uint32(0)
	if c.G >= 1 {
		g = 0xffff
	} else if c.G > 0 {
		g = uint32(c.G * 0xffff)
	}
	b := uint32(0)
	if c.B >= 1 {
		b = 0xffff
	} else if c.B > 0 {
		b = uint32(c.B * 0xffff)
	}
	return r, g, b, 0xffff
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

func (m Matrix4x4) Multiply(t Tuple) Tuple {
	return Tuple{m[0][0]*t[0] + m[0][1]*t[1] + m[0][2]*t[2] + m[0][3]*t[3],
		m[1][0]*t[0] + m[1][1]*t[1] + m[1][2]*t[2] + m[1][3]*t[3],
		m[2][0]*t[0] + m[2][1]*t[1] + m[2][2]*t[2] + m[2][3]*t[3],
		m[3][0]*t[0] + m[3][1]*t[1] + m[3][2]*t[2] + m[3][3]*t[3]}
}

func (m Matrix4x4) Transpose() Matrix4x4 {
	return Matrix4x4{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

func (m Matrix4x4) Mul(n Matrix4x4) Matrix4x4 {
	result := Matrix4x4{}
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			result[r][c] = m[r][0]*n[0][c] + m[r][1]*n[1][c] + m[r][2]*n[2][c] + m[r][3]*n[3][c]
		}
	}
	return result
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
			if !FloatEquals(m[r][c], n[r][c]) {
				return false
			}
		}
	}
	return true
}

func Translation(x, y, z float64) Matrix4x4 {
	return Matrix4x4{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}.Transpose()
}

func (m Matrix4x4) Translate(x, y, z float64) Matrix4x4 {
	return m.Mul(Translation(x, y, z))
}

func (m Matrix4x4) TranslateX(x float64) Matrix4x4 {
	return m.Translate(x, 0, 0)
}

func (m Matrix4x4) TranslateY(y float64) Matrix4x4 {
	return m.Translate(0, y, 0)
}

func (m Matrix4x4) TranslateZ(z float64) Matrix4x4 {
	return m.Translate(0, 0, z)
}

func Scaling(x, y, z float64) Matrix4x4 {
	return Matrix4x4{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

func (m Matrix4x4) Scale(x, y, z float64) Matrix4x4 {
	return m.Mul(Scaling(x, y, z))
}

func (m Matrix4x4) ScaleX(x float64) Matrix4x4 {
	return m.Scale(x, 1, 1)
}

func (m Matrix4x4) ScaleY(y float64) Matrix4x4 {
	return m.Scale(1, y, 1)
}

func (m Matrix4x4) ScaleZ(z float64) Matrix4x4 {
	return m.Scale(1, 1, z)
}

func RotationX(theta float64) Matrix4x4 {
	c := math.Cos(theta)
	s := math.Sin(theta)
	return Matrix4x4{
		{1, 0, 0, 0},
		{0, c, -s, 0},
		{0, s, c, 0},
		{0, 0, 0, 1},
	}
}

func (m Matrix4x4) RotateX(theta float64) Matrix4x4 {
	return m.Mul(RotationX(theta))
}

func RotationY(theta float64) Matrix4x4 {
	c := math.Cos(theta)
	s := math.Sin(theta)
	return Matrix4x4{
		{c, 0, s, 0},
		{0, 1, 0, 0},
		{-s, 0, c, 0},
		{0, 0, 0, 1},
	}
}

func (m Matrix4x4) RotateY(theta float64) Matrix4x4 {
	return m.Mul(RotationY(theta))
}

func RotationZ(theta float64) Matrix4x4 {
	c := math.Cos(theta)
	s := math.Sin(theta)
	return Matrix4x4{
		{c, -s, 0, 0},
		{s, c, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func (m Matrix4x4) RotateZ(theta float64) Matrix4x4 {
	return m.Mul(RotationZ(theta))
}
