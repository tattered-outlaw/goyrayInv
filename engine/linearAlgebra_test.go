package engine

import (
	"testing"
)

func TestMatrix4x4_equality(t *testing.T) {
	m := Matrix4x4{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	}
	if !m.Eq(m) {
		t.Fail()
	}
}
func TestMatrix4x4_Determinant_non_zero(t *testing.T) {
	m := Matrix4x4{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	}
	if !FloatEquals(m.Determinant(), -2120) {
		t.Fail()
	}
}
func TestMatrix4x4_Determinant_zero(t *testing.T) {
	m := Matrix4x4{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	}
	if !FloatEquals(m.Determinant(), 0) {
		t.Fail()
	}
}

func TestMatrix4x4_Inverse(t *testing.T) {
	m := Matrix4x4{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	}
	inv, err := m.Inverse()
	if err != nil {
		t.Fail()
	}
	e := Matrix4x4{
		{-0.15384615384615385, -0.15384615384615385, -0.28205128205128205, -0.5384615384615384},
		{-0.07692307692307693, 0.12307692307692308, 0.02564102564102564, 0.03076923076923077},
		{0.358974358974359, 0.358974358974359, 0.4358974358974359, 0.9230769230769231},
		{-0.6923076923076923, -0.6923076923076923, -0.7692307692307693, -1.9230769230769231},
	}
	if !inv.Eq(e) {
		t.Fail()
		t.Logf("%v", inv.Transpose())
	}
}
