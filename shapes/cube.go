package shapes

import (
	"math"

	. "goray/math"
	. "goray/rt"
)

type Cube struct{}

func (*Cube) LocalIntersect(_ *Engine, shape *Shape, localRay *Ray, intersections *Intersections) {
	xTMin, xTMax := checkAxis(localRay.Origin[0], localRay.Direction[0])
	yTMin, yTMax := checkAxis(localRay.Origin[1], localRay.Direction[1])
	zTMin, zTMax := checkAxis(localRay.Origin[2], localRay.Direction[2])

	tMin := max(xTMin, max(yTMin, zTMin))
	tMax := min(xTMax, min(yTMax, zTMax))

	if tMin <= tMax {
		intersections.Add(tMin, shape)
		intersections.Add(tMax, shape)
	}
}

func (*Cube) LocalNormalAt(_ *Shape, localPoint *Tuple) Tuple {
	x := math.Abs(localPoint[0])
	y := math.Abs(localPoint[1])
	z := math.Abs(localPoint[2])

	maxC := max(x, max(y, z))

	if maxC == x {
		return Vector(x, 0, 0)
	}
	if maxC == y {
		return Vector(0, y, 0)
	}
	return Vector(0, 0, z)
}

func (*Cube) BoundsOf(_ *Shape) BoundingBox {
	min := Point(-1, -1, -1)
	max := Point(1, 1, 1)
	return NBoundingBox(min, max)
}

func checkAxis(origin, direction float64) (float64, float64) {
	tMinNumerator := -1.0 - origin
	tMaxNumerator := 1.0 - origin

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
