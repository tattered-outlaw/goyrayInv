package scenes

import (
	"math"

	. "goray/math"
	. "goray/rt"
	. "goray/shapes"
)

func GroupScene1(width, height int) *Scene {
	pointLight1 := PointLight{
		Position:  Point(-10, 10, -20),
		Intensity: Color{R: 1, G: 1, B: 1}.Scale(0.5),
	}
	pointLight2 := PointLight{
		Position:  Point(10, 5, -20),
		Intensity: Color{R: 1, G: 1, B: 1}.Scale(0.5),
	}
	camera := NCamera(width, height, math.Pi/5, Point(0, 3, -20), Point(0, 0, 0), Vector(0, 1, 0))
	scene := NScene([]PointLight{pointLight1, pointLight2}, camera)

	zSpace := 2.0

	group := &Group{}
	groupShape := NShape(group)

	front := half()
	front.TranslateZ(-zSpace)
	group.Add(groupShape, front)

	back := half()
	back.TranslateZ(zSpace)
	group.Add(groupShape, back)

	scene.AddShape(groupShape)

	return scene
}

func half() *Shape {
	xSpace := 2.0

	group := &Group{}
	groupShape := NShape(group)

	left := quarter()
	left.TranslateX(-xSpace)
	group.Add(groupShape, left)

	right := quarter()
	right.TranslateX(xSpace)
	group.Add(groupShape, right)

	return groupShape
}

func quarter() *Shape {
	group := &Group{}
	groupShape := NShape(group)

	ySpace := 2.0

	member := eighth()
	member.TranslateY(ySpace)
	group.Add(groupShape, member)

	member = eighth()
	member.TranslateY(-ySpace)
	group.Add(groupShape, member)

	return groupShape
}

func eighth() *Shape {
	group := &Group{}
	groupShape := NShape(group)

	outGap := 1.75
	outComp := math.Sqrt(outGap * outGap / 2)

	sphereCentre := NShape(&Sphere{})
	sphereCentre.Material(DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}))
	group.Add(groupShape, sphereCentre)

	sphereNE := NShape(&Sphere{})
	sphereNE.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereNE.Scale(0.5, 0.5, 0.5)
	sphereNE.Translate(outComp, 0, outComp)
	group.Add(groupShape, sphereNE)

	sphereSE := NShape(&Sphere{})
	sphereSE.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereSE.Scale(0.5, 0.5, 0.5)
	sphereSE.Translate(outComp, 0, -outComp)
	group.Add(groupShape, sphereSE)

	sphereSW := NShape(&Sphere{})
	sphereSW.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereSW.Scale(0.5, 0.5, 0.5)
	sphereSW.Translate(-outComp, 0, -outComp)
	group.Add(groupShape, sphereSW)

	sphereNW := NShape(&Sphere{})
	sphereNW.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereNW.Scale(0.5, 0.5, 0.5)
	sphereNW.Translate(-outComp, 0, outComp)

	group.Add(groupShape, sphereNW)

	return groupShape
}
