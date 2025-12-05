package internal

type PointLight struct {
	Position  Tuple
	Intensity Color
}

type Scene struct {
	pointLights []*PointLight
	camera      *Camera
	shapes      []*Shape
}

func NScene(pointLights []PointLight, camera Camera) *Scene {
	scene := &Scene{
		pointLights: make([]*PointLight, len(pointLights)),
		camera:      &camera,
		shapes:      make([]*Shape, 0),
	}
	for i, pointLight := range pointLights {
		scene.pointLights[i] = &pointLight
	}
	return scene
}

func (scene *Scene) AddShape(shape *Shape) {
	shape.CalculateInverseTransformations()
	shape.CalculateBounds()
	scene.shapes = append(scene.shapes, shape)
}
