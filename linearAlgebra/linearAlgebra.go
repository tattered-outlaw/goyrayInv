package linearAlgebra

type Tuple [4]float64

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

type Matrix2x2 [2][2]float64
type Matrix3x3 [3][3]float64
type Matrix4x4 [4][4]float64

type Color struct {
	R, G, B float64
}

// RGBA Range clip the R, G, B components and convert to uint32
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
