package engine

import "sync"

type ShapeStrategy interface {
	localIntersect(shape *Shape, localRay *Ray, intersections *Intersections)
	localNormalAt(shape *Shape, worldPoint *Tuple) Tuple
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

func intersectShape(shape *Shape, worldRay *Ray, intersections *Intersections, rayPool *sync.Pool) {
	localRay := rayPool.Get().(*Ray)
	defer rayPool.Put(localRay)
	worldRay.TransformToShape(shape, localRay)
	shape.strategy.localIntersect(shape, localRay, intersections)
}

func (shape Shape) normalAt(worldPoint *Tuple, tuplePool *sync.Pool) Tuple {
	localPoint := tuplePool.Get().(*Tuple)
	defer tuplePool.Put(localPoint)
	MulTInPlace(shape.InverseTransformation, worldPoint, localPoint)
	localNormal := shape.strategy.localNormalAt(&shape, localPoint)
	worldNormal := shape.TransposeInverse.MulT(localNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
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

func (shape Shape) withMaterial(m Material) Shape {
	shape.material = m
	return shape
}

func (shape Shape) translate(x, y, z float64) Shape {
	shape.transformation = shape.transformation.Translate(x, y, z)
	return shape
}

func (shape Shape) translateX(x float64) Shape {
	shape.transformation = shape.transformation.TranslateX(x)
	return shape
}

func (shape Shape) translateY(y float64) Shape {
	shape.transformation = shape.transformation.TranslateY(y)
	return shape
}

func (shape Shape) translateZ(z float64) Shape {
	shape.transformation = shape.transformation.TranslateZ(z)
	return shape
}

func (shape Shape) scale(x, y, z float64) Shape {
	shape.transformation = shape.transformation.Scale(x, y, z)
	return shape
}

func (shape Shape) scaleX(x float64) Shape {
	shape.transformation = shape.transformation.ScaleX(x)
	return shape
}
func (shape Shape) scaleY(y float64) Shape {
	shape.transformation = shape.transformation.ScaleY(y)
	return shape
}
func (shape Shape) scaleZ(z float64) Shape {
	shape.transformation = shape.transformation.ScaleZ(z)
	return shape
}

func (shape Shape) rotateX(x float64) Shape {
	shape.transformation = shape.transformation.RotateX(x)
	return shape
}

func (shape Shape) rotateY(y float64) Shape {
	shape.transformation = shape.transformation.RotateY(y)
	return shape
}
func (shape Shape) rotateZ(z float64) Shape {
	shape.transformation = shape.transformation.RotateZ(z)
	return shape
}
