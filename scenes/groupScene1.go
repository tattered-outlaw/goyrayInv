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
	camera := NCamera(width, height, math.Pi/5, Point(0, 2, -20), Point(0, 0, 0), Vector(0, 1, 0))
	scene := NScene([]PointLight{pointLight1, pointLight2}, camera)

	zSpace := 3.0

	front := half()
	front.TranslateZ(-zSpace)
	scene.AddShape(front)

	back := half()
	back.TranslateZ(zSpace)
	scene.AddShape(back)

	return scene
}

func half() *Shape {
	xSpace := 3.0

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

	ySpace := 1.25

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

	sphereCentre := NShape(&Sphere{})
	sphereCentre.Material(DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}))
	group.Add(groupShape, sphereCentre)

	sphereWest := NShape(&Sphere{})
	sphereWest.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereWest.Scale(0.5, 0.5, 0.5)
	sphereWest.TranslateX(-1.75)
	group.Add(groupShape, sphereWest)

	sphereEast := NShape(&Sphere{})
	sphereEast.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereEast.Scale(0.5, 0.5, 0.5)
	sphereEast.TranslateX(1.75)
	group.Add(groupShape, sphereEast)

	sphereNorth := NShape(&Sphere{})
	sphereNorth.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereNorth.Scale(0.5, 0.5, 0.5)
	sphereNorth.TranslateZ(1.75)
	group.Add(groupShape, sphereNorth)

	sphereSouth := NShape(&Sphere{})
	sphereSouth.Material(DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	sphereSouth.Scale(0.5, 0.5, 0.5)
	sphereSouth.TranslateZ(-1.75)
	group.Add(groupShape, sphereSouth)

	return groupShape
}
