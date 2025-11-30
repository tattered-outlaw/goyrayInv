package engine

import (
	"math"
)

type Sphere struct {
	*BaseShape
}

func NSphere() *Sphere {
	return &Sphere{DefaultBaseshape()}
}

func (s *Sphere) intersect(ray Ray) []Intersect {
	ray.TransformToShape(s)
	sphereToRay := ray.Origin.Sub(Point(0, 0, 0))
	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return make([]Intersect, 0)
	} else {
		t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
		return []Intersect{{T: t1, Shape: s}, {T: t2, Shape: s}}
	}
}

func (s *Sphere) normalAt(worldPoint Tuple) Tuple {
	objectPoint := s.getInverseTransformation().MulT(worldPoint)
	objectNormal := objectPoint.Sub(Point(0, 0, 0))
	worldNormal := s.getTransposeInverse().MulT(objectNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
}
