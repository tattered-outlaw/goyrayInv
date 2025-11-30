package engine

import (
	"math"
)

func OneBall(width, height int) Scene {
	pointLight := PointLight{
		Position:  Point(-10, 10, -10),
		Intensity: Color{1, 1, 1},
	}

	shapes := make([]*Shape, 0)

	shape := NSphere()
	shape.setMaterial(DefaultMaterial().
		withColor(Color{R: 1, G: 0.2, B: 1}))
	shape.rotateY(math.Pi / 2)

	shapes = append(shapes, shape)

	camera := NCamera(width, height, math.Pi/4, Point(0, 0, -12), Point(0, 0, 0), Vector(0, 1, 0))

	return NScene([]PointLight{pointLight}, camera, shapes)
}

func ThreeBalls(width, height int) Scene {
	pointLight := PointLight{
		Position:  Point(-10, 10, -10),
		Intensity: Color{1, 1, 1},
	}

	shapes := make([]*Shape, 0)

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

	camera := NCamera(width, height, math.Pi/4, Point(0, -1, -12), Point(0, 0, 0), Vector(0, 1, 0))

	return NScene([]PointLight{pointLight}, camera, shapes)
}

func Chapter7(width, height int) Scene {
	shapes := make([]*Shape, 0)

	floor := NSphere()
	floor.scale(10, 0.01, 10)
	floor.translateY(-0.5)
	floor.setMaterial(DefaultMaterial().withColor(Color{R: 1, G: 0.9, B: 0.9}).withSpecular(0))
	shapes = append(shapes, floor)

	leftWall := NSphere()
	leftWall.setMaterial(floor.material)
	leftWall.scale(10, 0.01, 10)
	leftWall.rotateX(math.Pi / 2)
	leftWall.rotateY(-math.Pi / 4)
	leftWall.translate(0, 0, 5)
	shapes = append(shapes, leftWall)

	rightWall := NSphere()
	rightWall.setMaterial(floor.material)
	rightWall.scale(10, 0.01, 10)
	rightWall.rotateX(math.Pi / 2)
	rightWall.rotateY(math.Pi / 4)
	rightWall.translate(0, 0, 5)
	shapes = append(shapes, rightWall)

	middle := NSphere()
	middle.setMaterial(DefaultMaterial().withColor(Color{R: 0.1, G: 1, B: 0.5}).withDiffuse(0.7).withSpecular(0.3))
	middle.translate(-0.5, 1, 0.5)
	shapes = append(shapes, middle)

	right := NSphere()
	right.setMaterial(DefaultMaterial().withColor(Color{R: 0.5, G: 1, B: 0.1}).withDiffuse(0.7).withSpecular(0.3))
	right.scale(0.5, 0.5, 0.5)
	right.translate(1.5, 0.5, -0.5)
	shapes = append(shapes, right)

	left := NSphere()
	left.setMaterial(DefaultMaterial().withColor(Color{R: 1, G: 0.8, B: 0.1}).withDiffuse(0.7).withSpecular(0.3))
	left.scale(0.33, 0.33, 0.33)
	left.translate(-1.5, 0.33, -0.75)
	shapes = append(shapes, left)

	light := PointLight{Point(-10, 10, -10), Color{1, 1, 1}}

	camera := NCamera(width, height, math.Pi/3, Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	return NScene([]PointLight{light}, camera, shapes)

}
