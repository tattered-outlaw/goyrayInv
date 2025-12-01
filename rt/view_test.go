package rt

import (
	"testing"
)

func Test_viewTransform(t *testing.T) {
	t.Logf("%+v\n", viewTransform(Point(1, 3, 2), Point(4, -2, 8), Vector(1, 1, 0)))
}
