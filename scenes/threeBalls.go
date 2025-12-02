package scenes

import (
	"math"

	. "goray/math"
	. "goray/rt"
	. "goray/shapes"
)

func ThreeBalls(width, height int) *Scene {
	pointLight := PointLight{
		Position:  Point(-10, 10, -10),
		Intensity: Color{R: 1, G: 1, B: 1},
	}
	camera := NCamera(width, height, math.Pi/4, Point(0, -1, -12), Point(0, 0, 0), Vector(0, 1, 0))
	scene := NScene([]PointLight{pointLight}, camera)

	ball := NShape(&Sphere{})
	ball.Material(DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}))
	ball.TranslateX(1.1)
	ball.ScaleY(1)
	scene.AddShape(ball)

	ball = NShape(&Sphere{})
	ball.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	ball.TranslateX(-1.1)
	ball.ScaleY(1)
	scene.AddShape(ball)

	ball = NShape(&Sphere{})
	ball.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 0.2, B: 1}))
	ball.TranslateZ(2)
	ball.TranslateY(1)
	ball.ScaleY(1)
	scene.AddShape(ball)

	return scene
}
