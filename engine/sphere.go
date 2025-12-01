package engine

import (
	"math"
)

type Sphere struct{}

func (Sphere) localIntersect(shape *Shape, localRay *Ray, intersections *Intersections) {
	sphereToRay := localRay.Origin.Sub(Point(0, 0, 0))
	a := localRay.Direction.Dot(*localRay.Direction)
	b := 2 * localRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant >= 0 {
		t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
		intersections.Add(t1, shape)
		intersections.Add(t2, shape)
	}
}

func (Sphere) localNormalAt(_ *Shape, localPoint *Tuple) Tuple {
	return localPoint.Sub(Point(0, 0, 0))
}
