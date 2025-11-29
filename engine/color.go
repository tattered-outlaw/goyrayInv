package engine

var Black = Color{}
var LightCyan = Color{R: 0.9, G: 1.0, B: 1.0}
var Salmon = Color{R: 1.0, G: 0.5, B: 0.5}

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
