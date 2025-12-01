package rt

import (
	"math"
)

type Sphere struct{}

func (Sphere) localIntersect(shape *Shape, localRay *Ray, intersections *Intersections) {
	sphereToRay := localRay.origin.Sub(Point(0, 0, 0))
	a := localRay.direction.Dot(*localRay.direction)
	b := 2 * localRay.direction.Dot(sphereToRay)
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
