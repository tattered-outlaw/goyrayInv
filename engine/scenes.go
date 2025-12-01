package engine

import (
	"math"
)

func OneBall(width, height int) *Scene {
	pointLight := PointLight{
		Position:  Point(-10, 10, -10),
		Intensity: Color{1, 1, 1},
	}
	camera := NCamera(width, height, math.Pi/4, Point(0, 0, -12), Point(0, 0, 0), Vector(0, 1, 0))
	scene := NScene([]PointLight{pointLight}, camera)

	scene.AddShape(NShape(Sphere{}).
		withMaterial(DefaultMaterial().
			withColor(Color{R: 1, G: 0.2, B: 1})).translateX(-2))

	return scene
}

func ThreeBalls(width, height int) *Scene {
	pointLight := PointLight{
		Position:  Point(-10, 10, -10),
		Intensity: Color{1, 1, 1},
	}
	camera := NCamera(width, height, math.Pi/4, Point(0, -1, -12), Point(0, 0, 0), Vector(0, 1, 0))
	scene := NScene([]PointLight{pointLight}, camera)

	scene.AddShape(NShape(Sphere{}).
		withMaterial(DefaultMaterial().withColor(Color{R: 1, G: 0.2, B: 1})).
		translateX(1.1).scaleY(1))

	scene.AddShape(NShape(Sphere{}).
		withMaterial(DefaultMaterial().withColor(Color{R: 0.2, G: 1, B: 0.2})).
		translateX(-1.1).scaleY(1))

	scene.AddShape(NShape(Sphere{}).
		withMaterial(DefaultMaterial().withColor(Color{R: 0.2, G: 0.2, B: 1})).translateZ(2).
		translateY(1).scaleY(1))

	return scene
}

func BallScene(width, height int) *Scene {
	camera := NCamera(width, height, math.Pi/3, Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))
	lightCount := 2.0
	light1 := PointLight{Point(-10, 10, -10), Color{1, 1, 1}.Scale(1.0 / lightCount)}
	light2 := PointLight{Point(-5, 5, -20), Color{0.75, 0.75, 0.75}.Scale(1.0 / lightCount)}
	scene := NScene([]PointLight{light1, light2}, camera)

	material := DefaultMaterial().withColor(Color{R: 1, G: 0.9, B: 0.9}).withSpecular(0)

	// floor
	scene.AddShape(NShape(Sphere{}).withMaterial(material).
		scale(10, 0.01, 10))

	// left wall
	scene.AddShape(NShape(Sphere{}).withMaterial(material).
		scale(10, 0.01, 10).rotateX(math.Pi/2).rotateY(-math.Pi/4).translate(0, 0, 5))

	// right wall
	scene.AddShape(NShape(Sphere{}).withMaterial(material).
		scale(10, 0.01, 10).rotateX(math.Pi/2).rotateY(math.Pi/4).translate(0, 0, 5))

	//middle
	scene.AddShape(NShape(Sphere{}).withMaterial(DefaultMaterial().withColor(Color{R: 0.1, G: 1, B: 0.5}).withDiffuse(0.7).withSpecular(0.3)).
		translate(-0.5, 1, 0.5))

	//right
	scene.AddShape(NShape(Sphere{}).withMaterial(DefaultMaterial().withColor(Color{R: 0.5, G: 1, B: 0.1}).withDiffuse(0.7).withSpecular(0.3)).
		scale(0.5, 0.5, 0.5).translate(1.5, 0.5, -0.5))

	//left
	scene.AddShape(NShape(Sphere{}).withMaterial(DefaultMaterial().withColor(Color{R: 1, G: 0.8, B: 0.1}).withDiffuse(0.7).withSpecular(0.3)).
		scale(0.33, 0.33, 0.33).translate(-1.5, 0.33, -0.75))

	return scene
}
