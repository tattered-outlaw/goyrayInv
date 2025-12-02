package rt

import . "goray/math"

type PointLight struct {
	Position  Tuple
	Intensity Color
}

type Scene struct {
	pointLights []PointLight
	camera      Camera
	shapes      []*Shape
}

func NScene(pointLights []PointLight, camera Camera) *Scene {
	scene := Scene{
		pointLights: pointLights,
		camera:      camera,
		shapes:      make([]*Shape, 0),
	}
	return &scene
}

func (scene *Scene) AddShape(shape Shape) {
	shape = shape.calculateInverseTransformations()
	scene.shapes = append(scene.shapes, &shape)
}
