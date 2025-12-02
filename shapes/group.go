package shapes

import (
	. "goray/math"
	. "goray/rt"
)

type Group struct {
	shapes []*Shape
}

func (group *Group) Add(groupShape *Shape, shape *Shape) {
	shape.Parent = groupShape
	shape.CalculateInverseTransformations()
	group.shapes = append(group.shapes, shape)
}

func (group *Group) LocalIntersect(engine *Engine, _ *Shape, localRay *Ray, intersections *Intersections) {
	for _, child := range group.shapes {
		IntersectShape(engine, child, localRay, intersections)
	}
}

func (*Group) LocalNormalAt(_ *Shape, _ *Tuple) Tuple {
	panic("should never be called")
}
