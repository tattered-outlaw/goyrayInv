package internal

import "math"

type Plane struct {
	commonState *SceneObjectCommonState
}

func newPlane() *Plane {
	return &Plane{newSceneObjectCommonState()}
}

func (plane Plane) getCommonState() *SceneObjectCommonState {
	return plane.commonState
}

func (plane Plane) localIntersect(engine *Engine, localRay *Ray, intersections *Intersections) {
	if math.Abs(localRay.Direction[1]) < EPSILON {
		return
	}
	t := -localRay.Origin[1] / localRay.Direction[1]
	intersections.add(t, plane)
}

func (plane Plane) localNormalAt(localPoint *Tuple) Tuple {
	return Vector(0, 1, 0)
}

func (plane Plane) boundsOf() *BoundingBox {
	return newBoundingBox(Point(math.Inf(-1), 0, math.Inf(-1)), Point(math.Inf(+1), 0, math.Inf(+1)))
}
