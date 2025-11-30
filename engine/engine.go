package engine

import (
	"image/color"
	"sort"
)

func Prepare(canvasPixels int) Scene {
	pointLight := PointLight{
		Position:  Point(-10, 10, -10),
		Intensity: Color{1, 1, 1},
	}

	shapes := make([]Shape, 0)

	shape := NSphere()
	shape.setMaterial(DefaultMaterial().
		withColor(Color{R: 1, G: 0.2, B: 1}))
	shape.translateX(1.1)
	shape.scaleY(1)

	shapes = append(shapes, shape)

	shape = NSphere()
	shape.setMaterial(DefaultMaterial().
		withColor(Color{R: 0.2, G: 1, B: 0.2}))
	shape.translateX(-1.1)
	shape.scaleY(1)

	shapes = append(shapes, shape)

	shape = NSphere()
	shape.setMaterial(DefaultMaterial().
		withColor(Color{R: 0.2, G: 0.2, B: 1}))
	shape.translateZ(2)
	shape.translateY(1)
	shape.scaleY(1)

	shapes = append(shapes, shape)

	x := NScene(pointLight, canvasPixels, 10.0, 15.0, shapes, Point(0, 0, -5))
	return x
}

type Scene struct {
	pointLight   PointLight
	wallZ        float64
	wallSize     float64
	canvasPixels int
	shapes       []Shape
	rayOrigin    Tuple
	pixelSize    float64
	halfWallSize float64
}

func NScene(pointLight PointLight, canvasPixels int, wallZ float64, wallSize float64, shapes []Shape, rayOrigin Tuple) Scene {
	for _, s := range shapes {
		s.calculateInverseTransformations()
	}
	scene := Scene{
		pointLight:   pointLight,
		canvasPixels: canvasPixels,
		wallZ:        wallZ,
		wallSize:     wallSize,
		shapes:       shapes,
		rayOrigin:    rayOrigin,
		pixelSize:    wallSize / float64(canvasPixels),
		halfWallSize: wallSize / 2.0,
	}
	return scene
}

func (scene Scene) GetPixel(x, y int) color.Color {

	worldX := -scene.halfWallSize + scene.pixelSize*float64(x)
	worldY := scene.halfWallSize - scene.pixelSize*float64(y)

	position := Point(worldX, worldY, scene.wallZ)
	intersects := make([]Intersect, 0)
	ray := NRay(scene.rayOrigin, position.Sub(scene.rayOrigin).Normalize())
	for _, shape := range scene.shapes {
		intersects = append(intersects, shape.intersect(ray)...)
	}
	if len(intersects) == 0 {
		return Black
	} else {
		sort.Slice(intersects, func(i, j int) bool {
			return intersects[i].T < intersects[j].T
		})
		hit := intersects[0]
		point := ray.Position(hit.T)
		normal := hit.Shape.normalAt(point)
		eye := ray.Direction.Negate()
		return lighting(hit.Shape.getMaterial(), scene.pointLight, point, eye, normal)
	}

}
