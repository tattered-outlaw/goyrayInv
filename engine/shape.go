package engine

type Shape interface {
	intersect(Ray) []Intersect
	calculateInverseTransformation()
	getInverseTransformation() Matrix4x4
	translateX(float64)
}
type BaseShape struct {
	transformation        Matrix4x4
	inverseTransformation Matrix4x4
}

func (o *BaseShape) calculateInverseTransformation() {
	inv, _ := o.transformation.Inverse()
	o.inverseTransformation = inv
}

func (o *BaseShape) getInverseTransformation() Matrix4x4 {
	return o.inverseTransformation
}

func (o *BaseShape) translateX(x float64) {
	o.transformation = o.transformation.TranslateX(x)
}
