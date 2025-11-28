package engine

import (
	la "goray/linearAlgebra"
	"image/color"
)

func Prepare(width, height int) *CheckboxSource {
	lightCyan := la.Color{R: 0.9, G: 1.0, B: 1.0}
	salmon := la.Color{R: 1.0, G: 0.5, B: 0.5}
	return &CheckboxSource{width: width, height: height, squareSize: 100, even: lightCyan, odd: salmon}
}

type CheckboxSource struct {
	width, height, squareSize int
	even, odd                 color.Color
}

func (c *CheckboxSource) GetPixel(x, y int) color.Color {
	columnParity := -1
	if (x/c.squareSize)%2 == 0 {
		columnParity = 1
	}
	rowParity := -1
	if (y/c.squareSize)%2 == 0 {
		rowParity = 1
	}
	if columnParity*rowParity == 1 {
		return c.even
	} else {
		return c.odd
	}
}
