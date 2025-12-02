package shapes

import (
	. "goray/math"
	. "goray/rt"
)

type Group struct{}

func (group *Group) Add(groupShape *Shape, shape *Shape) {
	shape.Parent = groupShape
	shape.CalculateInverseTransformations()
	shape.CalculateBounds()
	groupShape.Children = append(groupShape.Children, shape)
}

func (group *Group) LocalIntersect(engine *Engine, groupShape *Shape, localRay *Ray, intersections *Intersections) {
	if !BBHitBy(groupShape.Bounds, localRay) {
		return
	}

	for _, child := range groupShape.Children {
		IntersectShape(engine, child, localRay, intersections)
	}
}

func (*Group) LocalNormalAt(_ *Shape, _ *Tuple) Tuple {
	panic("should never be called")
}

func (*Group) BoundsOf(groupShape *Shape) BoundingBox {
	bounds := EmptyBoundingBox()
	for _, child := range groupShape.Children {
		bounds = AddBoundingBoxes(bounds, *child.ParentSpaceBounds)
	}
	return bounds
}
