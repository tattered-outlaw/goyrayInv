package internal

import (
	"math"
)

type Sphere struct {
	commonState *SceneObjectCommonState
}

func newSphere() *Sphere {
	return &Sphere{newSceneObjectCommonState()}
}

func (sphere *Sphere) getCommonState() *SceneObjectCommonState {
	return sphere.commonState
}

func (sphere *Sphere) localIntersect(_ *Engine, localRay *Ray, intersections *Intersections) {
	sphereToRay := localRay.Origin.Sub(Point(0, 0, 0))
	a := localRay.Direction.Dot(*localRay.Direction)
	b := 2 * localRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant >= 0 {
		t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
		intersections.add(t1, sphere)
		intersections.add(t2, sphere)
	}
}

func (*Sphere) localNormalAt(localPoint *Tuple) Tuple {
	return localPoint.Sub(Point(0, 0, 0))
}

func (*Sphere) boundsOf() *BoundingBox {
	return newBoundingBox(Point(-1, -1, -1), Point(1, 1, 1))
}
