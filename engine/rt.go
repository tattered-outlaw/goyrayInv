package engine

import (
	"math"
)

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

func lighting(material Material, light PointLight, point Tuple, eyeV Tuple, normalV Tuple) Color {
	effectiveColor := material.Color.Multiply(light.Intensity)
	lightV := light.Position.Sub(point).Normalize()
	ambient := effectiveColor.Scale(material.Ambient)
	lightDotNormal := lightV.Dot(normalV)
	diffuse := Black
	specular := Black
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.Scale(material.Diffuse * lightDotNormal)
		reflectV := reflect(lightV.Scale(-1), normalV)
		reflectDotEye := reflectV.Dot(eyeV)
		if reflectDotEye > 0 {
			specular = light.Intensity.Scale(material.Specular * math.Pow(reflectDotEye, material.Shininess))
		}
	}
	return ambient.Add(diffuse).Add(specular)
}

func reflect(in Tuple, normal Tuple) Tuple {
	return in.Sub(normal.Scale(2 * in.Dot(normal)))
}
