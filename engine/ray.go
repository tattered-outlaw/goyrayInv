package engine

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func NRay(origin, direction Tuple) Ray {
	result := Ray{}
	result.Origin = origin
	result.Direction = direction
	return result
}

func (ray Ray) Position(t float64) Tuple {
	return ray.Origin.Add(ray.Direction.Scale(t))
}

func (ray Ray) TransformToShape(s Shape) Ray {
	t := s.getInverseTransformation()
	ray.Origin = t.MulT(ray.Origin)
	ray.Direction = t.MulT(ray.Direction)
	return ray
}
