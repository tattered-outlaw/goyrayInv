package internal

type SceneObject interface {
	getCommonState() *SceneObjectCommonState
	localIntersect(engine *Engine, localRay *Ray, intersections *Intersections)
	localNormalAt(localPoint *Tuple) Tuple
	boundsOf() *BoundingBox
}

type SceneObjectCommonState struct {
	parent                SceneObject
	transformation        *Matrix4x4
	material              *Material
	bounds                *BoundingBox
	parentSpaceBounds     *BoundingBox
	isIdentity            bool
	inverseTransformation *Matrix4x4
	transposeInverse      *Matrix4x4
}

func newSceneObjectCommonState() *SceneObjectCommonState {
	material := DefaultMaterial()
	return &SceneObjectCommonState{
		transformation:        newIdentity4(),
		isIdentity:            true,
		inverseTransformation: newIdentity4(),
		transposeInverse:      newIdentity4(),
		material:              &material,
	}
}

func calculateInverseTransformations(object SceneObject) {
	config := object.getCommonState()
	transformation := config.transformation
	if *transformation == Identity4 {
		config.isIdentity = true
		config.inverseTransformation = transformation
		config.transposeInverse = transformation
	} else {
		config.isIdentity = false
		inv := inverse4(config.transformation)
		config.inverseTransformation = inv
		config.transposeInverse = transpose4(inv)
	}
}

func setParent(object SceneObject, parent SceneObject) {
	object.getCommonState().parent = parent
}

func calculateBounds(object SceneObject) {
	config := object.getCommonState()
	bounds := object.boundsOf()
	config.bounds = bounds
	parentSpaceBounds := transformBoundingBox(config.bounds, config.transformation)
	config.parentSpaceBounds = parentSpaceBounds
}

func transform(object SceneObject, transformation *Matrix4x4) {
	t := object.getCommonState().transformation
	object.getCommonState().transformation = transformation.mul4x4(t)
}
