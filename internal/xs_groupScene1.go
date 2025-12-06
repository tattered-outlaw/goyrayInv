package internal

import (
	"math"
)

func GroupScene1(width, height int) *Scene {
	lightScale := 0.4
	pointLight1 := PointLight{
		Position:  Point(-10, 10, -20),
		Intensity: Color{R: 1, G: 1, B: 1}.Scale(lightScale),
	}
	pointLight2 := PointLight{
		Position:  Point(10, 5, -20),
		Intensity: Color{R: 1, G: 1, B: 1}.Scale(lightScale),
	}
	camera := NCamera(width, height, math.Pi/5, Point(0, 2, -50), Point(0, -2, 0), Vector(0, 1, 0))
	scene := newScene([]PointLight{pointLight1, pointLight2}, camera)

	zSpace := 2.0

	group := scene.rootGroup

	floor := newPlane()
	translate(floor, 0, -4, 0)
	group.add(floor)
	setMaterial(floor, DefaultMaterial().WithColor(Color{R: 0.5, G: 0.5, B: 0.5}).WithReflectivity(0.02))

	middle := newSphere()
	material := DefaultMaterial().WithColor(Color{R: 0.1, G: 0.1, B: 0.1}).WithReflectivity(0.4)
	setMaterial(middle, material)
	scale(middle, 1.5, 1.5, 1.5)
	group.add(middle)

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
	setMaterial(sphereCentre, DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}).WithReflectivity(0.025))
	group.add(sphereCentre)

	material := DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}).WithReflectivity(0.025)

	sphereNE := newSphere()
	setMaterial(sphereNE, material)
	scale(sphereNE, 0.5, 0.5, 0.5)
	translate(sphereNE, outComp, 0, outComp)
	group.add(sphereNE)

	sphereSE := newSphere()
	setMaterial(sphereSE, material)
	scale(sphereSE, 0.5, 0.5, 0.5)
	translate(sphereSE, outComp, 0, -outComp)
	group.add(sphereSE)

	sphereSW := newSphere()
	setMaterial(sphereSW, material)
	scale(sphereSW, 0.5, 0.5, 0.5)
	translate(sphereSW, -outComp, 0, -outComp)
	group.add(sphereSW)

	sphereNW := newSphere()
	setMaterial(sphereNW, material)
	scale(sphereNW, 0.5, 0.5, 0.5)
	translate(sphereNW, -outComp, 0, outComp)

	group.add(sphereNW)

	return group
}
