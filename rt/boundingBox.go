package rt

import (
	. "goray/math"
	"math"
)

type BoundingBox struct {
	min, max Tuple
}

func EmptyBoundingBox() BoundingBox {
	return BoundingBox{Point(math.Inf(+1), math.Inf(+1), math.Inf(+1)), Point(math.Inf(-1), math.Inf(-1), math.Inf(-1))}
}

func NBoundingBox(min, max Tuple) BoundingBox {
	return BoundingBox{min, max}
}

func AddPointToBoundingBox(boundingBox BoundingBox, point Tuple) BoundingBox {
	if point[0] < boundingBox.min[0] {
		boundingBox.min[0] = point[0]
	}
	if point[1] < boundingBox.min[1] {
		boundingBox.min[1] = point[1]
	}
	if point[2] < boundingBox.min[2] {
		boundingBox.min[2] = point[2]
	}
	if point[0] > boundingBox.max[0] {
		boundingBox.max[0] = point[0]
	}
	if point[1] > boundingBox.max[1] {
		boundingBox.max[1] = point[1]
	}
	if point[2] > boundingBox.max[2] {
		boundingBox.max[2] = point[2]
	}
	return boundingBox
}

func AddBoundingBoxes(boundingBox BoundingBox, otherBox BoundingBox) BoundingBox {
	boundingBox = AddPointToBoundingBox(boundingBox, otherBox.min)
	boundingBox = AddPointToBoundingBox(boundingBox, otherBox.max)
	return boundingBox
}

func TransformBoundingBox(boundingBox BoundingBox, transformation *Matrix4x4) BoundingBox {
	p1 := boundingBox.min
	p2 := Point(boundingBox.min[0], boundingBox.min[1], boundingBox.max[2])
	p3 := Point(boundingBox.min[0], boundingBox.max[1], boundingBox.min[2])
	p4 := Point(boundingBox.min[0], boundingBox.max[1], boundingBox.max[2])
	p5 := Point(boundingBox.max[0], boundingBox.min[1], boundingBox.min[2])
	p6 := Point(boundingBox.max[0], boundingBox.min[1], boundingBox.max[2])
	p7 := Point(boundingBox.max[0], boundingBox.max[1], boundingBox.min[2])
	p8 := boundingBox.max

	boundingBox = EmptyBoundingBox()

	for _, p := range [...]Tuple{p1, p2, p3, p4, p5, p6, p7, p8} {
		point := transformation.MulT(p)
		boundingBox = AddPointToBoundingBox(boundingBox, point)
	}

	return boundingBox
}

func BBHitBy(bbox *BoundingBox, localRay *Ray) bool {
	xTMin, xTMax := checkAxis(localRay.Origin[0], localRay.Direction[0], bbox.min[0], bbox.max[0])
	yTMin, yTMax := checkAxis(localRay.Origin[1], localRay.Direction[1], bbox.min[1], bbox.max[1])
	zTMin, zTMax := checkAxis(localRay.Origin[2], localRay.Direction[2], bbox.min[2], bbox.max[2])

	tMin := max(xTMin, max(yTMin, zTMin))
	tMax := min(xTMax, min(yTMax, zTMax))

	return tMin <= tMax
}

func checkAxis(origin, direction, min, max float64) (float64, float64) {
	tMinNumerator := min - origin
	tMaxNumerator := max - origin

	var tMin, tMax float64

	if math.Abs(direction) >= EPSILON {
		tMin = tMinNumerator / direction
		tMax = tMaxNumerator / direction
	} else {
		if tMinNumerator == 0.0 {
			tMin = 0
		} else {
			tMin = tMinNumerator * math.Inf(1)
		}
		if tMaxNumerator == 0.0 {
			tMax = 0
		} else {
			tMax = tMaxNumerator * math.Inf(1)
		}
	}

	if tMin <= tMax {
		return tMin, tMax
	} else {
		return tMax, tMin
	}
}

func (boundingBox BoundingBox) ContainsPoint(point Tuple) bool {
	return point[0] >= boundingBox.min[0] && point[0] <= boundingBox.max[0] &&
		point[1] >= boundingBox.min[1] && point[1] <= boundingBox.max[1] &&
		point[2] >= boundingBox.min[2] && point[2] <= boundingBox.max[2]
}

func (boundingBox BoundingBox) ContainsBox(otherBox BoundingBox) bool {
	return boundingBox.ContainsPoint(otherBox.min) && boundingBox.ContainsPoint(otherBox.max)
}
