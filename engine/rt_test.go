package engine

import (
	"testing"
)

func TestRay_create(t *testing.T) {
	Ray, err := NRay(Point(0, 0, -5), Vector(0, 0, 1))
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Log(Ray)
}

func TestRay_position(t *testing.T) {
	Ray, err := NRay(Point(2, 3, 4), Vector(1, 0, 0))
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Log(Ray.Position(0))
	t.Log(Ray.Position(1))
	t.Log(Ray.Position(-1))
	t.Log(Ray.Position(2.5))
}
