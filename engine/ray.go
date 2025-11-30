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

func (r *Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Scale(t))
}

func (r *Ray) TransformToShape(s Shape) {
	t := s.getInverseTransformation()
	newRay := NRay(t.MulT(r.Origin), t.MulT(r.Direction))
	r.Origin = newRay.Origin
	r.Direction = newRay.Direction
}
