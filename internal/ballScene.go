package internal

//import (
//	"math"
//)
//
//func BallScene(width, height int) *Scene {
//	camera := NewCamera(width, height, math.Pi/3, Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))
//	lightCount := 2.0
//	light1 := PointLight{
//		Position:  Point(-10, 10, -10),
//		Intensity: Color{R: 1, G: 1, B: 1}.Scale(1.0 / lightCount),
//	}
//	light2 := PointLight{
//		Position:  Point(-5, 5, -20),
//		Intensity: Color{R: 0.75, G: 0.75, B: 0.75}.Scale(1.0 / lightCount)}
//	scene := NScene([]PointLight{light1, light2}, camera)
//
//	material := DefaultMaterial().WithColor(Color{R: 1, G: 0.9, B: 0.9}).WithSpecular(0)
//
//	// floor
//	floor := NShape(&Sphere{})
//	floor.Material(material)
//	floor.Scale(10, 0.01, 10)
//	scene.AddShape(floor)
//
//	// left wall
//	leftWall := NShape(&Sphere{})
//	leftWall.Material(material)
//	leftWall.Scale(10, 0.01, 10)
//	leftWall.RotateX(math.Pi / 2)
//	leftWall.RotateY(-math.Pi / 4)
//	leftWall.Translate(0, 0, 5)
//	scene.AddShape(leftWall)
//
//	// right wall
//	rightWall := NShape(&Sphere{})
//	rightWall.Material(material)
//	rightWall.Scale(10, 0.01, 10)
//	rightWall.RotateX(math.Pi / 2)
//	rightWall.RotateY(math.Pi / 4)
//	rightWall.Translate(0, 0, 5)
//	scene.AddShape(rightWall)
//
//	//middle
//	middle := NShape(&Sphere{})
//	middle.Material(DefaultMaterial().WithColor(Color{R: 0.1, G: 1, B: 0.5}).WithDiffuse(0.7).WithSpecular(0.3))
//	middle.Translate(-0.5, 1, 0.5)
//	scene.AddShape(middle)
//
//	//right
//	right := NShape(&Sphere{})
//	right.Material(DefaultMaterial().WithColor(Color{R: 0.5, G: 1, B: 0.1}).WithDiffuse(0.7).WithSpecular(0.3))
//	right.Scale(0.5, 0.5, 0.5)
//	right.Translate(1.5, 0.5, -0.5)
//	scene.AddShape(right)
//
//	//left
//	left := NShape(&Sphere{})
//	left.Material(DefaultMaterial().WithColor(Color{R: 1, G: 0.8, B: 0.1}).WithDiffuse(0.7).WithSpecular(0.3))
//	left.Scale(0.33, 0.33, 0.33)
//	left.Translate(-1.5, 0.33, -0.75)
//	scene.AddShape(left)
//
//	return scene
//}
