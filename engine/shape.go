package engine

type Shape interface {
	intersect(ray Ray) []Intersect
	normalAt(worldPoint Tuple) Tuple
	calculateInverseTransformation()
	getInverseTransformation() Matrix4x4
	getMaterial() Material
	setMaterial(Material)
	translateX(float64)
	translateY(float64)
	translateZ(float64)
	scaleY(float64)
}

type BaseShape struct {
	transformation        Matrix4x4
	inverseTransformation Matrix4x4
	material              Material
}

func (o *BaseShape) calculateInverseTransformation() {
	inv, _ := o.transformation.Inverse()
	o.inverseTransformation = inv
}

func (o *BaseShape) getInverseTransformation() Matrix4x4 {
	return o.inverseTransformation
}

func (o *BaseShape) getMaterial() Material {
	return o.material
}

func (o *BaseShape) setMaterial(m Material) {
	o.material = m
}

func (o *BaseShape) translateX(x float64) {
	o.transformation = o.transformation.TranslateX(x)
}

func (o *BaseShape) translateY(y float64) {
	o.transformation = o.transformation.TranslateY(y)
}

func (o *BaseShape) translateZ(z float64) {
	o.transformation = o.transformation.TranslateZ(z)
}

func (o *BaseShape) scaleY(y float64) {
	o.transformation = o.transformation.ScaleY(y)
}
