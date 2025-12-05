package internal

import (
	"math"
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
	scene := newScene([]PointLight{pointLight1, pointLight2}, camera)

	zSpace := 2.0

	group := scene.rootGroup

	front := half()
	translate(front, 0, 0, -zSpace)
	group.add(front)

	back := half()
	translate(back, 0, 0, zSpace)
	group.add(back)

	return scene
}

func half() SceneObject {
	xSpace := 2.0

	group := newGroup()

	left := quarter()
	translate(left, -xSpace, 0, 0)
	group.add(left)

	right := quarter()
	translate(right, xSpace, 0, 0)
	group.add(right)

	return group
}

func quarter() SceneObject {
	group := newGroup()

	ySpace := 2.0

	member := eighth()
	translate(member, 0, -ySpace, 0)
	group.add(member)

	member = eighth()
	translate(member, 0, ySpace, 0)
	group.add(member)

	return group
}

func eighth() SceneObject {
	group := newGroup()

	outGap := 1.75
	outComp := math.Sqrt(outGap * outGap / 2)

	sphereCentre := newSphere()
	setMaterial(sphereCentre, DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}))
	group.add(sphereCentre)

	sphereNE := newSphere()
	setMaterial(sphereNE, DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	scale(sphereNE, 0.5, 0.5, 0.5)
	translate(sphereNE, outComp, 0, outComp)
	group.add(sphereNE)

	sphereSE := newSphere()
	setMaterial(sphereSE, DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	scale(sphereSE, 0.5, 0.5, 0.5)
	translate(sphereSE, outComp, 0, -outComp)
	group.add(sphereSE)

	sphereSW := newSphere()
	setMaterial(sphereSW, DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	scale(sphereSW, 0.5, 0.5, 0.5)
	translate(sphereSW, -outComp, 0, -outComp)
	group.add(sphereSW)

	sphereNW := newSphere()
	setMaterial(sphereNW, DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}))
	scale(sphereNW, 0.5, 0.5, 0.5)
	translate(sphereNW, -outComp, 0, outComp)

	group.add(sphereNW)

	return group
}
