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
	result, err := NRay(t.MulT(r.Origin), t.MulT(r.Direction).Normalize())
	if err != nil {
		panic(err)
	}
	return result, err
}

type Intersect struct {
	T     float64
	Shape Shape
}

type PointLight struct {
	Position  Tuple
	Intensity Color
}

type Material struct {
	Color                                 Color
	Ambient, Diffuse, Specular, Shininess float64
}

func (m Material) withColor(c Color) Material {
	return Material{c, m.Ambient, m.Diffuse, m.Specular, m.Shininess}
}

func (m Material) withAmbient(a float64) Material {
	return Material{m.Color, a, m.Diffuse, m.Specular, m.Shininess}
}

func (m Material) withDiffuse(d float64) Material {
	return Material{m.Color, m.Ambient, d, m.Specular, m.Shininess}
}

func (m Material) withSpecular(s float64) Material {
	return Material{m.Color, m.Ambient, m.Diffuse, s, m.Shininess}
}

func (m Material) withShininess(s float64) Material {
	return Material{m.Color, m.Ambient, m.Diffuse, m.Specular, s}
}

func DefaultMaterial() Material {
	return Material{Color: Color{1, 1, 1}, Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200.0}
}

func lighting(material Material, light PointLight, point Tuple, eyeV Tuple, normalV Tuple) {

}
