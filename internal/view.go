package internal

import (
	"math"
)

func viewTransform(from, to, up Tuple) Matrix4x4 {
	forward := to.Sub(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)
	orientation := Matrix4x4{
		{left[0], left[1], left[2], 0},
		{trueUp[0], trueUp[1], trueUp[2], 0},
		{-forward[0], -forward[1], -forward[2], 0},
		{0, 0, 0, 1},
	}
	return *orientation.mul4x4(Translation(-from[0], -from[1], -from[2]))
}

type Camera struct {
	hSize, vSize          int
	halfWidth, halfHeight float64
	fov                   float64
	transform             Matrix4x4
	pixelSize             float64
}

func NCamera(width, height int, fov float64, from, to, up Tuple) Camera {
	camera := Camera{hSize: width, vSize: height}
	halfView := math.Tan(fov / 2)
	aspect := float64(width) / float64(height)
	if aspect >= 1 {
		camera.halfWidth = halfView
		camera.halfHeight = halfView / aspect
	} else {
		camera.halfWidth = halfView * aspect
		camera.halfHeight = halfView
	}
	camera.pixelSize = (camera.halfWidth * 2) / float64(width)
	transform, err := viewTransform(from, to, up).Inverse()
	if err != nil {
		panic(err)
	}
	camera.transform = transform
	return camera
}

func (camera Camera) rayForPixel(x, y int) Ray {
	xOffset := (float64(x) + 0.5) * camera.pixelSize
	yOffset := (float64(y) + 0.5) * camera.pixelSize
	worldX := camera.halfWidth - xOffset
	worldY := camera.halfHeight - yOffset
	pixel := camera.transform.MulT(Point(worldX, worldY, -1))
	origin := camera.transform.MulT(Point(0, 0, 0))
	direction := pixel.Sub(origin).Normalize()
	return NRay(&origin, &direction)
}
