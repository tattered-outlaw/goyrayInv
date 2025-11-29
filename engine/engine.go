package engine

import (
	"image/color"
)

func Prepare(canvasPixels int) *Scene {
	shape := NSphere()
	//shape.translateX(1)
	//shape.scaleY(1.5)
	shape.setMaterial(DefaultMaterial())
	x := NScene(canvasPixels, 10.0, 15.0, shape, Point(0, 0, -5))
	return x
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
	shape.calculateInverseTransformation()
	return &scene
}

func (s *Scene) GetPixel(x, y int) color.Color {

	worldX := -s.halfWallSize + s.pixelSize*float64(x)
	worldY := s.halfWallSize - s.pixelSize*float64(y)

	position := Point(worldX, worldY, s.wallZ)
	ray, _ := NRay(s.rayOrigin, position.Sub(s.rayOrigin))
	ray, _ = ray.TransformToShape(s.shape)
	intersects := s.shape.intersect(ray)
	if len(intersects) == 0 {
		return LightCyan
	} else {
		return Salmon
	}

}
