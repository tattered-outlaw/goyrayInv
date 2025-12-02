package rt

import . "goray/math"

type ShapeStrategy interface {
	LocalIntersect(shape *Shape, localRay *Ray, intersections *Intersections)
	LocalNormalAt(shape *Shape, worldPoint *Tuple) Tuple
}

type Shape struct {
	transformation        Matrix4x4
	InverseTransformation *Matrix4x4
	TransposeInverse      *Matrix4x4
	material              Material
	strategy              ShapeStrategy
}

func NShape(strategy ShapeStrategy) *Shape {
	return &Shape{
		transformation:        Identity4,
		InverseTransformation: &Identity4,
		TransposeInverse:      &Identity4,
		material:              DefaultMaterial(),
		strategy:              strategy,
	}
}

func (shape Shape) calculateInverseTransformations() Shape {
	inv, _ := shape.transformation.Inverse()
	shape.InverseTransformation = &inv
	transInverse := inv.Transpose()
	shape.TransposeInverse = &transInverse
	return shape
}

func (shape Shape) getMaterial() Material {
	return shape.material
}

func (shape Shape) WithMaterial(m Material) Shape {
	shape.material = m
	return shape
}

func (shape Shape) Translate(x, y, z float64) Shape {
	shape.transformation = shape.transformation.Translate(x, y, z)
	return shape
}

func (shape Shape) TranslateX(x float64) Shape {
	shape.transformation = shape.transformation.TranslateX(x)
	return shape
}

func (shape Shape) TranslateY(y float64) Shape {
	shape.transformation = shape.transformation.TranslateY(y)
	return shape
}

func (shape Shape) TranslateZ(z float64) Shape {
	shape.transformation = shape.transformation.TranslateZ(z)
	return shape
}

func (shape Shape) Scale(x, y, z float64) Shape {
	shape.transformation = shape.transformation.Scale(x, y, z)
	return shape
}

func (shape Shape) ScaleX(x float64) Shape {
	shape.transformation = shape.transformation.ScaleX(x)
	return shape
}
func (shape Shape) ScaleY(y float64) Shape {
	shape.transformation = shape.transformation.ScaleY(y)
	return shape
}
func (shape Shape) ScaleZ(z float64) Shape {
	shape.transformation = shape.transformation.ScaleZ(z)
	return shape
}

func (shape Shape) RotateX(x float64) Shape {
	shape.transformation = shape.transformation.RotateX(x)
	return shape
}

func (shape Shape) RotateY(y float64) Shape {
	shape.transformation = shape.transformation.RotateY(y)
	return shape
}
func (shape Shape) RotateZ(z float64) Shape {
	shape.transformation = shape.transformation.RotateZ(z)
	return shape
}
