package internal

//import (
//	"math"
//)
//
//func BBoxTest(width, height int) *Scene {
//	pointLight := PointLight{
//		Position:  Point(-10, 10, -10),
//		Intensity: Color{R: 1, G: 1, B: 1},
//	}
//	camera := NCamera(width, height, math.Pi/4, Point(0, 0, -8), Point(0, 0, 0), Vector(0, 1, 0))
//	scene := NScene([]PointLight{pointLight}, camera)
//
//	group := &Group{}
//	groupShape := NShape(group)
//
//	ball := NShape(&Sphere{})
//	ball.Material(DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}))
//	ball.TranslateX(-1)
//	group.add(groupShape, ball)
//
//	ball = NShape(&Sphere{})
//	ball.Material(DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}))
//	ball.TranslateX(1)
//	group.add(groupShape, ball)
//
//	scene.AddShape(groupShape)
//
//	return scene
//}
