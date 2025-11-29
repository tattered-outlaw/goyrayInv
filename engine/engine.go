package engine

import (
	"image/color"
)

func Prepare(canvasPixels int) *Scene {
	x := NScene(canvasPixels, 10.0, 7.0, NSphere(), Point(0, 0, -5))
	return x
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

type Scene struct {
	wallZ        float64
	wallSize     float64
	canvasPixels int
	shape        Shape
	rayOrigin    Tuple
	pixelSize    float64
	halfWallSize float64
}

func NScene(canvasPixels int, wallZ float64, wallSize float64, shape Shape, rayOrigin Tuple) *Scene {
	scene := Scene{
		canvasPixels: canvasPixels,
		wallZ:        wallZ,
		wallSize:     wallSize,
		shape:        shape,
		rayOrigin:    rayOrigin,
		pixelSize:    wallSize / float64(canvasPixels),
		halfWallSize: wallSize / 2.0,
	}
	return &scene
}

func (s *Scene) GetPixel(x, y int) color.Color {
	s.shape.calculateInverseTransformation()

	worldX := -s.halfWallSize + s.pixelSize*float64(x)
	worldY := s.halfWallSize - s.pixelSize*float64(y)

	position := Point(worldX, worldY, s.wallZ)
	ray, _ := NRay(s.rayOrigin, position.Sub(s.rayOrigin).Normalize())
	intersects := s.shape.intersect(ray)
	if len(intersects) == 0 {
		return LightCyan
	} else {
		return Salmon
	}

}
