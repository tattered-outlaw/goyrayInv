package internal

//import (
//	"math"
//)
//
//func GroupScene0(width, height int) *Scene {
//	pointLight1 := PointLight{
//		Position:  Point(-10, 10, -20),
//		Intensity: Color{R: 1, G: 1, B: 1}.Scale(0.5),
//	}
//	pointLight2 := PointLight{
//		Position:  Point(10, 5, -20),
//		Intensity: Color{R: 1, G: 1, B: 1}.Scale(0.5),
//	}
//	camera := NewCamera(width, height, math.Pi/5, Point(0, 2, -20), Point(0, 0, 0), Vector(0, 1, 0))
//	scene := NScene([]PointLight{pointLight1, pointLight2}, camera)
//
//	zSpace := 1.0
//
//	group := &Group{}
//	groupShape := NShape(group)
//
//	front := half0()
//	front.TranslateZ(-zSpace)
//	group.add(groupShape, front)
//
//	back := half0()
//	back.TranslateZ(zSpace)
//	group.add(groupShape, back)
//
//	scene.AddShape(groupShape)
//
//	return scene
//}
//
//func half0() *Shape {
//	xSpace := 1.0
//
//	group := &Group{}
//	groupShape := NShape(group)
//
//	left := quarter0()
//	left.TranslateX(-xSpace)
//	group.add(groupShape, left)
//
//	right := quarter0()
//	right.TranslateX(xSpace)
//	group.add(groupShape, right)
//
//	return groupShape
//}
//
//func quarter0() *Shape {
//	group := &Group{}
//	groupShape := NShape(group)
//
//	ySpace := 1.0
//
//	member := eighth0()
//	member.TranslateY(ySpace)
//	group.add(groupShape, member)
//
//	member = eighth0()
//	member.TranslateY(-ySpace)
//	group.add(groupShape, member)
//
//	return groupShape
//}
//
//func eighth0() *Shape {
//	sphere := NShape(&Sphere{})
//	sphere.Material(DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}))
//	return sphere
//}
