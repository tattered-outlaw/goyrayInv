package rt

import . "goray/math"

type ShapeStrategy interface {
	LocalIntersect(engine *Engine, shape *Shape, localRay *Ray, intersections *Intersections)
	LocalNormalAt(shape *Shape, worldPoint *Tuple) Tuple
}

type Shape struct {
	Parent                *Shape
	transformation        Matrix4x4
	InverseTransformation *Matrix4x4
	TransposeInverse      *Matrix4x4
	material              Material
	strategy              ShapeStrategy
}

func NShape(strategy ShapeStrategy) *Shape {
	return &Shape{
		Parent:                nil,
		transformation:        Identity4,
		InverseTransformation: &Identity4,
		TransposeInverse:      &Identity4,
		material:              DefaultMaterial(),
		strategy:              strategy,
	}
}

func (shape *Shape) CalculateInverseTransformations() {
	inv, _ := shape.transformation.Inverse()
	shape.InverseTransformation = &inv
	transInverse := inv.Transpose()
	shape.TransposeInverse = &transInverse
}

func (shape *Shape) getMaterial() Material {
	return shape.material
}

func (shape *Shape) Material(m Material) {
	shape.material = m
}

func (shape *Shape) Translate(x, y, z float64) {
	shape.transformation = shape.transformation.Translate(x, y, z)
}

func (shape *Shape) TranslateX(x float64) {
	shape.transformation = shape.transformation.TranslateX(x)
}

func (shape *Shape) TranslateY(y float64) {
	shape.transformation = shape.transformation.TranslateY(y)
}

func (shape *Shape) TranslateZ(z float64) {
	shape.transformation = shape.transformation.TranslateZ(z)
}

func (shape *Shape) Scale(x, y, z float64) {
	shape.transformation = shape.transformation.Scale(x, y, z)
}

func (shape *Shape) ScaleX(x float64) {
	shape.transformation = shape.transformation.ScaleX(x)
}
func (shape *Shape) ScaleY(y float64) {
	shape.transformation = shape.transformation.ScaleY(y)
}
func (shape *Shape) ScaleZ(z float64) {
	shape.transformation = shape.transformation.ScaleZ(z)
}

func (shape *Shape) RotateX(x float64) {
	shape.transformation = shape.transformation.RotateX(x)
}

func (shape *Shape) RotateY(y float64) {
	shape.transformation = shape.transformation.RotateY(y)
}
func (shape *Shape) RotateZ(z float64) {
	shape.transformation = shape.transformation.RotateZ(z)
}
