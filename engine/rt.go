package engine

import (
	"fmt"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func NRay(origin, direction Tuple) (Ray, error) {
	result := Ray{}
	if !origin.IsPoint() {
		return result, fmt.Errorf("origin must be a point")
	}
	if !direction.IsVector() {
		return result, fmt.Errorf("direction must be a vector")
	}
	result.Origin = origin
	result.Direction = direction
	return result, nil
}

func (r Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Scale(t))
}

func (r Ray) TransformToShape(s Shape) (Ray, error) {
	t := s.getInverseTransformation()
	return NRay(t.MulT(r.Origin), t.MulT(r.Direction))
}

type Intersect struct {
	T     float64
	Shape Shape
}
