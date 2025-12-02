package rt

import . "goray/math"

type Material struct {
	Color                                 Color
	Ambient, Diffuse, Specular, Shininess float64
}

func (m Material) WithColor(c Color) Material {
	return Material{c, m.Ambient, m.Diffuse, m.Specular, m.Shininess}
}

func (m Material) WithAmbient(a float64) Material {
	return Material{m.Color, a, m.Diffuse, m.Specular, m.Shininess}
}

func (m Material) WithDiffuse(d float64) Material {
	return Material{m.Color, m.Ambient, d, m.Specular, m.Shininess}
}

func (m Material) WithSpecular(s float64) Material {
	return Material{m.Color, m.Ambient, m.Diffuse, s, m.Shininess}
}

func (m Material) WithShininess(s float64) Material {
	return Material{m.Color, m.Ambient, m.Diffuse, m.Specular, s}
}

func DefaultMaterial() Material {
	return Material{Color: Color{1, 1, 1}, Ambient: 0.15, Diffuse: 0.9, Specular: 0.9, Shininess: 200.0}
}
