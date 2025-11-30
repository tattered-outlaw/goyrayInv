package engine

import (
	"math"
)

func NSphere() *Shape {
	result := NShape(Sphere{})
	return &result
}

type Sphere struct{}

func (Sphere) localIntersect(shape Shape, localRay Ray) []Intersect {
	sphereToRay := localRay.Origin.Sub(Point(0, 0, 0))
	a := localRay.Direction.Dot(localRay.Direction)
	b := 2 * localRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil
	} else {
		t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
		return []Intersect{{T: t1, Shape: &shape}, {T: t2, Shape: &shape}}
	}
}

func (Sphere) localNormalAt(_ Shape, localPoint Tuple) Tuple {
	return localPoint.Sub(Point(0, 0, 0))
}
