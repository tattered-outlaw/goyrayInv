package engine

type ShapeStrategy interface {
	localIntersect(shape Shape, localRay Ray) []Intersect
	localNormalAt(shape Shape, worldPoint Tuple) Tuple
}

type Shape struct {
	transformation        Matrix4x4
	inverseTransformation Matrix4x4
	transposeInverse      Matrix4x4
	material              Material
	strategy              ShapeStrategy
}

func NShape(strategy ShapeStrategy) Shape {
	return Shape{
		transformation:        Identity4,
		inverseTransformation: Identity4,
		transposeInverse:      Identity4,
		material:              DefaultMaterial(),
		strategy:              strategy,
	}
}

func (shape Shape) intersect(worldRay Ray) []Intersect {
	return shape.strategy.localIntersect(shape, worldRay.TransformToShape(shape))
}

func (shape Shape) normalAt(worldPoint Tuple) Tuple {
	localNormal := shape.strategy.localNormalAt(shape, shape.getInverseTransformation().MulT(worldPoint))
	worldNormal := shape.getTransposeInverse().MulT(localNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
}

func (shape Shape) calculateInverseTransformations() Shape {
	inv, _ := shape.transformation.Inverse()
	shape.inverseTransformation = inv
	shape.transposeInverse = inv.Transpose()
	return shape
}

func (shape Shape) getInverseTransformation() Matrix4x4 {
	return shape.inverseTransformation
}

func (shape Shape) getTransposeInverse() Matrix4x4 {
	return shape.transposeInverse
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
