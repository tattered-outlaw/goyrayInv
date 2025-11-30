package engine

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
