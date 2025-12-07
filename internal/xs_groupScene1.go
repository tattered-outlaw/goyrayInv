package internal

import (
	"math"
)

func GroupScene1(width, height int) *Scene {
	lightScale := 0.4
	zSpace := 2.0

	return buildScene(width, height).
		withCamera(math.Pi/5, Point(-10, 2, -50), Point(0, -2, 0), Vector(0, 1, 0)).
		addPointLight(Point(-10, 10, -20), Color{R: 1, G: 1, B: 1}.Scale(lightScale)).
		addPointLight(Point(50, 25, -100), Color{R: 1, G: 1, B: 1}.Scale(lightScale)).
		//wall
		add(object(newPlane()).rotateX(-math.Pi / 2).translateZ(12).
			material(DefaultMaterial().WithColor(Color{R: 0.2, G: 0.2, B: 0.2}).WithReflectivity(0.0))).
		//floor
		add(object(newPlane()).translateY(-4).
			material(DefaultMaterial().WithPattern(
				newCheckerPattern(Color{R: 0.8, G: 0.8, B: 0.8}, Color{R: 0.1, G: 0.1, B: 0.1})).WithReflectivity(0.07))).
		// middle shiny sphere
		add(object(newSphere()).
			material(DefaultMaterial().WithColor(Color{R: 0.1, G: 0.1, B: 0.1}).WithReflectivity(0.4)).
			scale(1.5)).
		add(half().translateZ(-zSpace)).
		add(half().translateZ(zSpace)).
		build()

}

func half() *SceneObjectBuilder {
	xSpace := 2.0
	return group().
		add(quarter().translateX(-xSpace)).
		add(quarter().translateX(xSpace))

}

func quarter() *SceneObjectBuilder {
	ySpace := 2.0
	return group().
		add(eighth().translateY(-ySpace)).
		add(eighth().translateY(ySpace))
}

func eighth() *SceneObjectBuilder {
	outGap := 1.75
	outComp := math.Sqrt(outGap * outGap / 2)
	innerMaterial := DefaultMaterial().WithColor(Color{R: 1, G: 0.2, B: 1}).WithReflectivity(0.025)
	outerMaterial := DefaultMaterial().WithColor(Color{R: 0.2, G: 1, B: 0.2}).WithReflectivity(0.025)
	return group().
		add(object(newSphere()).material(innerMaterial)).
		add(object(newSphere()).material(outerMaterial).scale(0.5).translate(outComp, 0, outComp)).
		add(object(newSphere()).material(outerMaterial).scale(0.5).translate(outComp, 0, -outComp)).
		add(object(newSphere()).material(outerMaterial).scale(0.5).translate(-outComp, 0, outComp)).
		add(object(newSphere()).material(outerMaterial).scale(0.5).translate(-outComp, 0, -outComp))
}
