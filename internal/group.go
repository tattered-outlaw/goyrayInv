package internal

type Group struct {
	objectConfig *SceneObjectCommonState
	children     []SceneObject
}

func newGroup() *Group {
	return &Group{newSceneObjectCommonState(), make([]SceneObject, 0)}
}

func (group *Group) getCommonState() *SceneObjectCommonState {
	return group.objectConfig
}

func (group *Group) localIntersect(engine *Engine, localRay *Ray, intersections *Intersections) {
	if !BBHitBy(group.getCommonState().bounds, localRay) {
		return
	}

	for _, child := range group.children {
		intersectObject(engine, child, localRay, intersections)
	}
}

func (*Group) localNormalAt(_ *Tuple) Tuple {
	panic("should never be called")
}

func (group *Group) boundsOf() *BoundingBox {
	bounds := emptyBoundingBox()
	for _, child := range group.children {
		addBoundingBoxTo(bounds, child.getCommonState().parentSpaceBounds)
	}
	return bounds
}

func (group *Group) add(child SceneObject) {
	calculateInverseTransformations(child)
	calculateBounds(child)
	group.children = append(group.children, child)
	setParent(child, group)
}
