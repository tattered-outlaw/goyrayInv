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

func unGroup(group *Group, isRoot bool) []SceneObject {
	result := make([]SceneObject, 0, len(group.children))
	for _, child := range group.children {
		switch child.(type) {
		case *Group:
			result = append(result, unGroup(child.(*Group), false)...)
		default:
			result = append(result, child)
		}
	}
	if !isRoot {
		m := group.getCommonState().transformation
		for _, object := range result {
			transform(object, m)
			calculateInverseTransformations(object)
			calculateBounds(object)
			object.getCommonState().parent = group.getCommonState().parent
		}
	}
	return result
}

func divideGroup(group *Group, threshold int) {
	if threshold <= len(group.children) {
		newChildren := make([]SceneObject, 0)
		leftChildren := make([]SceneObject, 0)
		rightChildren := make([]SceneObject, 0)
		overlapChildren := make([]SceneObject, 0)
		leftBox, rightBox := splitBoundingBox(group.boundsOf())
		for _, child := range group.children {
			bounds := child.getCommonState().parentSpaceBounds
			if leftBox.ContainsBox(bounds) {
				leftChildren = append(leftChildren, child)
			} else if rightBox.ContainsBox(bounds) {
				rightChildren = append(rightChildren, child)
			} else {
				overlapChildren = append(overlapChildren, child)
			}
		}
		if len(leftChildren) > 0 {
			leftGroup := newGroup()
			leftGroup.children = leftChildren
			leftGroup.getCommonState().parent = group.getCommonState().parent
			leftGroup.getCommonState().bounds = leftBox
			divideGroup(leftGroup, threshold)
			newChildren = append(newChildren, leftGroup)
		}
		if len(rightChildren) > 0 {
			rightGroup := newGroup()
			rightGroup.children = rightChildren
			rightGroup.getCommonState().parent = group.getCommonState().parent
			rightGroup.getCommonState().bounds = rightBox
			divideGroup(rightGroup, threshold)
			newChildren = append(newChildren, rightGroup)
		}
		newChildren = append(newChildren, overlapChildren...)
		group.children = newChildren
	}
}
