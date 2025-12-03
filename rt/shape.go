package rt

import . "goray/math"

type ShapeStrategy interface {
	LocalIntersect(engine *Engine, shape *Shape, localRay *Ray, intersections *Intersections)
	LocalNormalAt(shape *Shape, worldPoint *Tuple) Tuple
	BoundsOf(*Shape) BoundingBox
}

type Shape struct {
	Parent                *Shape
	Children              []*Shape
	Bounds                *BoundingBox
	ParentSpaceBounds     *BoundingBox
	Transformation        *Matrix4x4
	InverseTransformation *Matrix4x4
	TransposeInverse      *Matrix4x4
	material              *Material
	strategy              ShapeStrategy
}

func NShape(strategy ShapeStrategy) *Shape {
	material := DefaultMaterial()
	return &Shape{
		Parent:                nil,
		Transformation:        &Identity4,
		InverseTransformation: &Identity4,
		TransposeInverse:      &Identity4,
		material:              &material,
		strategy:              strategy,
	}
}

func (shape *Shape) CalculateInverseTransformations() {
	inv, _ := shape.Transformation.Inverse()
	shape.InverseTransformation = &inv
	transInverse := inv.Transpose()
	shape.TransposeInverse = &transInverse
}

func (shape *Shape) CalculateBounds() {
	bounds := shape.strategy.BoundsOf(shape)
	shape.Bounds = &bounds
	parentSpaceBounds := TransformBoundingBox(*shape.Bounds, shape.Transformation)
	shape.ParentSpaceBounds = &parentSpaceBounds
}

func (shape *Shape) getMaterial() *Material {
	return shape.material
}

func (shape *Shape) Material(m Material) {
	shape.material = &m
}

func (shape *Shape) Translate(x, y, z float64) {
	shape.Transformation = shape.Transformation.Translate(x, y, z)
}

func (shape *Shape) TranslateX(x float64) {
	shape.Transformation = shape.Transformation.TranslateX(x)
}

func (shape *Shape) TranslateY(y float64) {
	shape.Transformation = shape.Transformation.TranslateY(y)
}

func (shape *Shape) TranslateZ(z float64) {
	shape.Transformation = shape.Transformation.TranslateZ(z)
}

func (shape *Shape) Scale(x, y, z float64) {
	shape.Transformation = shape.Transformation.Scale(x, y, z)
}

func (shape *Shape) ScaleX(x float64) {
	shape.Transformation = shape.Transformation.ScaleX(x)
}
func (shape *Shape) ScaleY(y float64) {
	shape.Transformation = shape.Transformation.ScaleY(y)
}
func (shape *Shape) ScaleZ(z float64) {
	shape.Transformation = shape.Transformation.ScaleZ(z)
}

func (shape *Shape) RotateX(x float64) {
	shape.Transformation = shape.Transformation.RotateX(x)
}

func (shape *Shape) RotateY(y float64) {
	shape.Transformation = shape.Transformation.RotateY(y)
}
func (shape *Shape) RotateZ(z float64) {
	shape.Transformation = shape.Transformation.RotateZ(z)
}
