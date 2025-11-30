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

func (o *Shape) intersect(worldRay Ray) []Intersect {
	return o.strategy.localIntersect(*o, worldRay.TransformToShape(*o))
}

func (o *Shape) normalAt(worldPoint Tuple) Tuple {
	localNormal := o.strategy.localNormalAt(*o, o.getInverseTransformation().MulT(worldPoint))
	worldNormal := o.getTransposeInverse().MulT(localNormal)
	worldNormal[3] = 0
	return worldNormal.Normalize()
}

func (o *Shape) calculateInverseTransformations() {
	inv, _ := o.transformation.Inverse()
	o.inverseTransformation = inv
	o.transposeInverse = inv.Transpose()
}

func (o *Shape) getInverseTransformation() Matrix4x4 {
	return o.inverseTransformation
}

func (o *Shape) getTransposeInverse() Matrix4x4 {
	return o.transposeInverse
}

func (o *Shape) getMaterial() Material {
	return o.material
}

func (o *Shape) setMaterial(m Material) {
	o.material = m
}

func (o *Shape) translate(x, y, z float64) {
	o.transformation = o.transformation.Translate(x, y, z)
}

func (o *Shape) translateX(x float64) {
	o.transformation = o.transformation.TranslateX(x)
}

func (o *Shape) translateY(y float64) {
	o.transformation = o.transformation.TranslateY(y)
}

func (o *Shape) translateZ(z float64) {
	o.transformation = o.transformation.TranslateZ(z)
}

func (o *Shape) scale(x, y, z float64) {
	o.transformation = o.transformation.Scale(x, y, z)
}

func (o *Shape) scaleX(x float64) {
	o.transformation = o.transformation.ScaleX(x)
}
func (o *Shape) scaleY(y float64) {
	o.transformation = o.transformation.ScaleY(y)
}
func (o *Shape) scaleZ(z float64) {
	o.transformation = o.transformation.ScaleZ(z)
}

func (o *Shape) rotateX(x float64) {
	o.transformation = o.transformation.RotateX(x)
}
func (o *Shape) rotateY(y float64) {
	o.transformation = o.transformation.RotateY(y)
}
func (o *Shape) rotateZ(z float64) {
	o.transformation = o.transformation.RotateZ(z)
}
