package engine

import (
	"testing"
)

func TestSphere(t *testing.T) {
	ray, err := NRay(Point(0, 0, 5), Vector(0, 0, 1))
	if err != nil {
		t.Fatalf("%v", err)
	}
	sphere := NSphere()
	t.Logf("%v", sphere.intersect(ray))
}
