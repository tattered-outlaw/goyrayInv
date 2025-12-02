package scenes

import (
	"math"

	. "goray/math"
	. "goray/rt"
	. "goray/shapes"
)

func BallScene(width, height int) *Scene {
	camera := NCamera(width, height, math.Pi/3, Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))
	lightCount := 2.0
	light1 := PointLight{
		Position:  Point(-10, 10, -10),
		Intensity: Color{R: 1, G: 1, B: 1}.Scale(1.0 / lightCount),
	}
	light2 := PointLight{
		Position:  Point(-5, 5, -20),
		Intensity: Color{R: 0.75, G: 0.75, B: 0.75}.Scale(1.0 / lightCount)}
	scene := NScene([]PointLight{light1, light2}, camera)

	material := DefaultMaterial().WithColor(Color{R: 1, G: 0.9, B: 0.9}).WithSpecular(0)

	// floor
	scene.AddShape(NShape(Sphere{}).WithMaterial(material).
		Scale(10, 0.01, 10))

	// left wall
	scene.AddShape(NShape(Sphere{}).WithMaterial(material).
		Scale(10, 0.01, 10).RotateX(math.Pi/2).RotateY(-math.Pi/4).Translate(0, 0, 5))

	// right wall
	scene.AddShape(NShape(Sphere{}).WithMaterial(material).
		Scale(10, 0.01, 10).RotateX(math.Pi/2).RotateY(math.Pi/4).Translate(0, 0, 5))

	//middle
	scene.AddShape(NShape(Sphere{}).WithMaterial(DefaultMaterial().WithColor(Color{R: 0.1, G: 1, B: 0.5}).WithDiffuse(0.7).WithSpecular(0.3)).
		Translate(-0.5, 1, 0.5))

	//right
	scene.AddShape(NShape(Sphere{}).WithMaterial(DefaultMaterial().WithColor(Color{R: 0.5, G: 1, B: 0.1}).WithDiffuse(0.7).WithSpecular(0.3)).
		Scale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5))

	//left
	scene.AddShape(NShape(Sphere{}).WithMaterial(DefaultMaterial().WithColor(Color{R: 1, G: 0.8, B: 0.1}).WithDiffuse(0.7).WithSpecular(0.3)).
		Scale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75))

	return scene
}
